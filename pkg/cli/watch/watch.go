package watch

import (
	"context"
	"fmt"

	controller "github.com/olisajc/appScaler/pkg/controller"
	"github.com/olisajc/appScaler/pkg/kubeclient"
	"github.com/olisajc/appScaler/pkg/policyscaler"
	ps "github.com/olisajc/appScaler/pkg/policyscaler"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

func WatchCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "watch",
		Short: "Watch resources",
		RunE: func(cmd *cobra.Command, args []string) error {

			fmt.Println("Starting to watch PolicyScaler resources...")

			service := policyscaler.NewPolicyScalerService()
			k8client, err := kubeclient.InitKubeClient(true)

			if err != nil {
				return fmt.Errorf("failed to initialize kubeclient: %v", err)
			}
			err = RunWatch(service, k8client)
			if err != nil {
				return fmt.Errorf("failed to run watch command: %v", err)
			}
			return nil
		},
	}

	return cmd
}

func RunWatch(service *ps.PolicyScalerService, k8client *kubeclient.KubeClient) error {

	grv := ps.GetPolicyScalerSchema()
	listWatch := createListWatch(k8client, grv)
	informer := createInformer(listWatch)
	indexer := informer.GetIndexer()

	addEventHandlers(informer, service)

	ctr := controller.NewPolicyScalerController(informer, indexer, createWorkqueue())

	err := controller.ValidateControllerOptions(ctr)
	if err != nil {
		return fmt.Errorf("invalid controller options: %v", err)
	}

	return nil
}

func createWorkqueue() workqueue.TypedRateLimitingInterface[string] {
	queue := workqueue.NewTypedRateLimitingQueue(workqueue.DefaultTypedControllerRateLimiter[string]())
	return queue
}

func addEventHandlers(informer cache.SharedIndexInformer, service *ps.PolicyScalerService) {
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			unstructuredObj := obj.(*unstructured.Unstructured)
			policyScaler, err := service.Converter.FromUnstructured(unstructuredObj.Object)
			if err != nil {
				fmt.Printf("Error converting added object: %v\n", err)
				return
			}
			fmt.Printf("PolicyScaler Added: %s\n", policyScaler.Name)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			unstructuredObj := newObj.(*unstructured.Unstructured)
			policyScaler, err := service.Converter.FromUnstructured(unstructuredObj.Object)
			if err != nil {
				fmt.Printf("Error converting updated object: %v\n", err)
				return
			}
			fmt.Printf("PolicyScaler Updated: %s\n", policyScaler.Name)
		},
		DeleteFunc: func(obj interface{}) {
			unstructuredObj := obj.(*unstructured.Unstructured)
			policyScaler, err := service.Converter.FromUnstructured(unstructuredObj.Object)
			if err != nil {
				fmt.Printf("Error converting deleted object: %v\n", err)
				return
			}
			fmt.Printf("PolicyScaler Deleted: %s\n", policyScaler.Name)
		},
	})
}

func createInformer(listWatch *cache.ListWatch) cache.SharedIndexInformer {
	informer := cache.NewSharedIndexInformer(
		listWatch,
		&unstructured.Unstructured{},
		0,
		cache.Indexers{},
	)

	return informer
}
func createListWatch(k8client *kubeclient.KubeClient, grv schema.GroupVersionResource) *cache.ListWatch {

	resourceInterface := k8client.DynamicClient().Resource(grv)

	return &cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			return resourceInterface.List(context.TODO(), options)
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			return resourceInterface.Watch(context.TODO(), options)
		},
	}
}
