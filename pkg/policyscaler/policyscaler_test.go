package policyscaler

import (
	"context"
	"testing"

	"github.com/olisajc/appScaler/pkg/kubeclient"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	dynamicfake "k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/kubernetes/scheme"
)

var fakeScheme = runtime.NewScheme()

var fakeObjects = []runtime.Object{
	&unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "extensions.example.com/v1",
			"kind":       "policyScaler",
			"metadata": map[string]interface{}{
				"name":      "policyName",
				"namespace": "default",
			},
			"spec": map[string]interface{}{
				"policies": map[string]interface{}{},
			},
		},
	},

	&unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "extensions.example.com/v1",
			"kind":       "policyScaler",
			"metadata": map[string]interface{}{
				"name":      "anotherPolicy",
				"namespace": "default",
			},
			"spec": map[string]interface{}{
				"policies": map[string]interface{}{},
			},
		},
	},

	&unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "extensions.example.com/v1",
			"kind":       "policyScaler",
			"metadata": map[string]interface{}{
				"name":      "anotherPolicy",
				"namespace": "custom-namespace",
			},
			"spec": map[string]interface{}{
				"policies": map[string]interface{}{},
			},
		},
	},
}

var fakeDynamicClient = dynamicfake.NewSimpleDynamicClient(fakeScheme, fakeObjects...)

func TestGetPolicy(t *testing.T) {
	_ = scheme.AddToScheme(fakeScheme)
	t.Run("kube client is nil", func(t *testing.T) {
		service := &PolicyScalerService{}
		_, err := service.Get(context.Background(), nil, "policyName", "default")
		assert.Error(t, err)
	})

	t.Run("policyScaler name is empty", func(t *testing.T) {
		service := &PolicyScalerService{}
		client := &kubeclient.KubeClient{}
		_, err := service.Get(context.Background(), client, "", "default")
		assert.Error(t, err)
	})

	t.Run("namespace is empty, should default to 'default'", func(t *testing.T) {
		service := NewPolicyScalerService()
		client := kubeclient.FakeKubeClient(fakeDynamicClient)
		policyscaler, err := service.Get(context.Background(), client, "policyName", "")
		assert.NoError(t, err)
		assert.NotNil(t, policyscaler)
	})

	t.Run("successful get policyScaler", func(t *testing.T) {
		service := NewPolicyScalerService()
		client := kubeclient.FakeKubeClient(fakeDynamicClient)
		policyscaler, err := service.Get(context.Background(), client, "policyName", "default")
		assert.NoError(t, err)
		assert.NotNil(t, policyscaler)
	})

	t.Run("policyScaler not found", func(t *testing.T) {
		service := NewPolicyScalerService()
		client := kubeclient.FakeKubeClient(fakeDynamicClient)
		_, err := service.Get(context.Background(), client, "nonExistentPolicy", "default")
		assert.Error(t, err)
	})

	t.Run("different namespace", func(t *testing.T) {
		service := NewPolicyScalerService()
		client := kubeclient.FakeKubeClient(fakeDynamicClient)
		policyscaler, err := service.Get(context.Background(), client, "anotherPolicy", "custom-namespace")
		assert.NoError(t, err)
		assert.NotNil(t, policyscaler)
	})

}

func TestList(t *testing.T) {
	_ = scheme.AddToScheme(fakeScheme)
	t.Run("kube client is nil", func(t *testing.T) {
		service := &PolicyScalerService{}
		_, err := service.List(context.Background(), nil, "default")
		assert.Error(t, err)
	})

	t.Run("namespace is empty, should default to 'default'", func(t *testing.T) {
		service := NewPolicyScalerService()
		client := kubeclient.FakeKubeClient(fakeDynamicClient)
		policyscalers, err := service.List(context.Background(), client, "")
		assert.NoError(t, err)
		assert.NotNil(t, policyscalers)
	})
	t.Run("list policiescaler", func(t *testing.T) {
		service := NewPolicyScalerService()
		client := kubeclient.FakeKubeClient(fakeDynamicClient)
		policyscalers, err := service.List(context.Background(), client, "default")
		assert.NoError(t, err)
		assert.NotNil(t, policyscalers)
	})

}
