package internal

type Hashset[T comparable] struct {
	m map[T]struct{}
}

func (hs *Hashset[T]) Set(k T) {
	hs.m[k] = struct{}{}
}

func (hs *Hashset[T]) Delete(k T) {
	delete(hs.m, k)
}

func (hs *Hashset[T]) Exists(k T) bool {
	_, ok := hs.m[k]

	return ok
}

func (hs *Hashset[T]) List() []T {
	ans := make([]T, len(hs.m))

	i := 0
	for k := range hs.m {
		ans[i] = k
		i += 1
	}

	return ans
}

func (hs *Hashset[T]) Len() int {
	return len(hs.m)
}

func NewHashset[T comparable](l []T) *Hashset[T] {
	hs := Hashset[T]{
		m: map[T]struct{}{},
	}

	for _, elem := range l {
		hs.Set(elem)
	}

	return &hs
}

func NewEmptyHashset[T comparable]() *Hashset[T] {
	return NewHashset([]T{})
}
