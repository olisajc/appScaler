package policyscaler

import (
	"context"
	"fmt"

	"github.com/olisajc/appScaler/pkg/converter"
	"github.com/olisajc/appScaler/pkg/kubeclient"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type PolicyScalerService struct {
	Schema    schema.GroupVersionResource
	Converter converter.Converter[*PolicyScaler]
}

func (app *PolicyScalerService) List(context context.Context, client *kubeclient.KubeClient, namespace string) (PolicyScalerList, error) {
	if client == nil {
		return nil, kubeclient.ErrNilClient
	}

	if namespace == "" {
		namespace = "default"
	}

	dynamicClient := client.DynamicClient()
	policyResourceList, err := dynamicClient.Resource(app.Schema).Namespace(namespace).
		List(context, metav1.ListOptions{})

	if err != nil {
		return nil, fmt.Errorf("failed to list policiescalers in namespace %s: %v", namespace, err)
	}

	var policyscalers PolicyScalerList
	for _, obj := range policyResourceList.Items {
		object, err := app.Converter.FromUnstructured(obj.Object)
		if err != nil {
			return nil, fmt.Errorf("failed to convert unstructured to PolicyScaler: %v", err)
		}
		policyscalers = append(policyscalers, object)
	}

	return policyscalers, nil
}

func (app *PolicyScalerService) Get(ctx context.Context, client *kubeclient.KubeClient, name, namespace string) (*PolicyScaler, error) {
	if client == nil {
		return nil, kubeclient.ErrNilClient
	}

	if name == "" {
		return nil, fmt.Errorf("policy name is required")
	}

	if namespace == "" {
		namespace = "default"
	}

	dynamicClient := client.DynamicClient()
	policyResource, err := dynamicClient.Resource(app.Schema).Namespace(namespace).
		Get(ctx, name, metav1.GetOptions{})

	if err != nil {
		return nil, fmt.Errorf("failed to get policy %s in namespace %s: %v", name, namespace, err)
	}

	object, err := app.Converter.FromUnstructured(policyResource.Object)
	if err != nil {
		return nil, fmt.Errorf("failed to convert unstructured to PolicyScaler: %v", err)
	}

	return object, nil
}

func NewPolicyScalerService() *PolicyScalerService {
	return &PolicyScalerService{
		Schema:    GetPolicyScalerSchema(),
		Converter: converter.NewTypeConverter[*PolicyScaler](),
	}
}
