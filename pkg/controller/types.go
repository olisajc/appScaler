package controller

import (
	"context"

	cache "k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type Controller interface {
	Run(ctx context.Context, workers int) error
	Informer() cache.SharedIndexInformer
	Indexer() cache.Indexer
	Workqueue() workqueue.TypedRateLimitingInterface[string]
	processNextItem() error
}

type PolicyScalerController struct {
	informer  cache.SharedIndexInformer
	indexer   cache.Indexer
	workqueue workqueue.TypedRateLimitingInterface[string]
}
