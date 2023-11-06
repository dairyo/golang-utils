package datatypes

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](es ...T) Set[T] {
	s := make(Set[T])
	for _, e := range es {
		s.Add(e)
	}
	return s
}

func (s Set[T]) Add(e T) {
	s[e] = struct{}{}
}

func (s Set[T]) Append(es ...T) {
	for _, e := range es {
		s.Add(e)
	}
}

func (s Set[T]) Array() []T {
	keys := make([]T, 0, len(s))
	for e := range s {
		keys = append(keys, e)
	}
	return keys
}
