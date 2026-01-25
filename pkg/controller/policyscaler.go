package controller

import (
	"context"

	cache "k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

func (c *PolicyScalerController) Informer() cache.SharedIndexInformer {
	return c.informer
}

func (c *PolicyScalerController) Indexer() cache.Indexer {
	return c.indexer
}

func (c *PolicyScalerController) Run(ctx context.Context, workers int) error {
	return nil
}

func (c *PolicyScalerController) processNextItem() error {
	return nil
}
func (c *PolicyScalerController) Workqueue() workqueue.TypedRateLimitingInterface[string] {
	return c.workqueue
}
