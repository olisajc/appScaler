package kubeclient

import (
	"k8s.io/client-go/dynamic"
	kubernetes "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type KubeClientOptions struct {
	config        *rest.Config
	DefaultClient *kubernetes.Clientset
	DynamicClient dynamic.Interface
}

type KubeClient struct {
	client *kubernetes.Clientset
	dynCli dynamic.Interface
}

type KubeClientOption func(*KubeClientOptions) error
