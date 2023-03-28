package internal

func MapSlice[T any, M any](slice []T, f func(T) M) []M {
	n := make([]M, len(slice))
	for i, e := range slice {
		n[i] = f(e)
	}
	return n
}
