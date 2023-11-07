package datatypes

func ToValue[T any](t *T) T {
	if t == nil {
		return *new(T)
	}
	return *t
}

func ToPtr[T comparable](t T) *T {
	if t == *new(T) {
		return nil
	}
	return &t
}
