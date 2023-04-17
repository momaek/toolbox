package sliceutils

// RemoveDuplicateStringSlice remove duplicate item in string slice
func RemoveDuplicateStringSlice(slice []string) []string {
	m := make(map[string]struct{})
	retSlice := make([]string, 0, len(slice))
	for _, v := range slice {
		m[v] = struct{}{}
		if len(m) > len(retSlice) {
			retSlice = append(retSlice, v)
		}
	}

	return retSlice
}

// RemoveDuplicate ..
func RemoveDuplicate[T comparable](slice []T) []T {
	m := make(map[T]struct{})
	retSlice := make([]T, 0, len(slice))
	for _, v := range slice {
		m[v] = struct{}{}
		if len(m) > len(retSlice) {
			retSlice = append(retSlice, v)
		}
	}

	return retSlice
}
