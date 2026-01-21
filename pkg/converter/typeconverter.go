package converter

import "k8s.io/apimachinery/pkg/runtime"

func (tc *TypeConverter[T]) FromUnstructured(data map[string]interface{}) (T, error) {
	return tc.FromUnstructuredFunc(data)
}

func FromUnstructured[T any](data map[string]interface{}) (T, error) {
	var zero T
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(data, &zero)
	return zero, err
}

func ToUnstructured[T any](obj T) (map[string]interface{}, error) {
	return runtime.DefaultUnstructuredConverter.ToUnstructured(&obj)
}
