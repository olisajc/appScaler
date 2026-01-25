package controller

import (
	"fmt"

	cache "k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

func NewPolicyScalerController(informer cache.SharedIndexInformer, indexer cache.Indexer, workqueue workqueue.TypedRateLimitingInterface[string]) Controller {
	return &PolicyScalerController{
		informer:  informer,
		indexer:   indexer,
		workqueue: workqueue,
	}
}

func ValidateControllerOptions(controller Controller) error {

	if controller == nil {
		return fmt.Errorf("controller is nil")
	}

	if controller.Informer() == nil {
		return fmt.Errorf("informer is nil")
	}

	if controller.Indexer() == nil {
		return fmt.Errorf("indexer is nil")
	}

	if controller.Workqueue() == nil {
		return fmt.Errorf("workqueue is nil")
	}
	return nil
}
