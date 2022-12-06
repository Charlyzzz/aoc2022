package utils

func Every[T any](n int, elems []T) [][]T {
	var groups [][]T
	var group []T
	for _, elem := range elems {
		if len(group) < n {
			group = append(group, elem)
		} else {
			groups = append(groups, group)
			group = []T{elem}
		}
	}
	groups = append(groups, group)
	return groups
}
