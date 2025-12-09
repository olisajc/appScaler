package kubeclient

import (
	"fmt"
	"path/filepath"
	"sync"

	"k8s.io/client-go/dynamic"

	kubernetes "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	ConfigPath               = filepath.Join(homedir.HomeDir(), ".kube", "config")
	ErrMultipleConfigSources = fmt.Errorf("multiple kubeconfig sources specified")
	k8ClientOnce             sync.Once
	k8Client                 *KubeClient
)

func InitKubeClient(isDevelopment bool) (*KubeClient, error) {
	var err error
	k8ClientOnce.Do(func() {
		if isDevelopment {
			k8Client, err = newDevelopmentKubeClient()
			return
		}
		if k8Client, err = newClusterKubeClient(); err != nil {
			return
		}
	})

	return k8Client, nil
}

func developmentConfig() (*rest.Config, error) {
	return clientcmd.BuildConfigFromFlags("", ConfigPath)
}

func clusterConfig() (*rest.Config, error) {
	return rest.InClusterConfig()
}

func developmentConfigOption() KubeClientOption {
	return func(o *KubeClientOptions) error {
		config, err := developmentConfig()
		if err != nil {
			return err
		}
		o.config = config
		return nil
	}
}

func clusterConfigOption() KubeClientOption {
	return func(o *KubeClientOptions) error {
		config, err := clusterConfig()
		if err != nil {
			return err
		}
		o.config = config
		return nil
	}
}

func clientSetOption() KubeClientOption {
	return func(o *KubeClientOptions) error {
		if o.config == nil {
			return fmt.Errorf("kube config is nil")
		}
		clientset, err := kubernetes.NewForConfig(o.config)
		if err != nil {
			return err
		}
		o.DefaultClient = clientset
		return nil
	}
}

func dynamicClientOption() KubeClientOption {
	return func(o *KubeClientOptions) error {
		if o.config == nil {
			return fmt.Errorf("kube config is nil")
		}
		dynamicClient, err := dynamic.NewForConfig(o.config)
		if err != nil {
			return err
		}
		o.DynamicClient = dynamicClient
		return nil
	}
}

func withKubeClientOptions(opts ...KubeClientOption) (KubeClientOptions, error) {
	var kubeOptions KubeClientOptions
	for _, opt := range opts {
		if err := opt(&kubeOptions); err != nil {
			return KubeClientOptions{}, err
		}
	}
	return kubeOptions, nil
}

func newDevelopmentKubeClient() (*KubeClient, error) {
	kubeOptions, err := withKubeClientOptions(
		developmentConfigOption(),
		clientSetOption(),
		dynamicClientOption(),
	)

	if err != nil {
		return nil, err
	}

	return &KubeClient{
		client: kubeOptions.DefaultClient,
		dynCli: kubeOptions.DynamicClient,
	}, nil

}

func newClusterKubeClient() (*KubeClient, error) {
	kubeOptions, err := withKubeClientOptions(
		clusterConfigOption(),
		clientSetOption(),
		dynamicClientOption(),
	)

	if err != nil {
		return nil, err
	}

	client := kubeOptions.DefaultClient
	dyn := kubeOptions.DynamicClient

	kubecli := &KubeClient{
		client: client,
		dynCli: dyn,
	}

	return kubecli, nil

}
