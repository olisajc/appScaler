package converter

type Converter[T any] interface {
	FromUnstructured(map[string]interface{}) (T, error)
}

type TypeConverter[T any] struct {
	FromUnstructuredFunc func(map[string]interface{}) (T, error)
}

func NewTypeConverter[T any]() *TypeConverter[T] {
	return &TypeConverter[T]{
		FromUnstructuredFunc: FromUnstructured[T],
	}
}
