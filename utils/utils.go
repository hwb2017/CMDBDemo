package utils

// StrSliceDiff return diff set of A - B
func StrSliceDiff(a , b []string) (diff []string) {
	m := make(map[string]bool)

	for _, item := range b {
        m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return diff
}

// StrSliceIntersection return intersection set of A and B
func StrSliceIntersection(a , b []string) (diff []string) {
	m := make(map[string]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; ok {
			diff = append(diff, item)
		}
	}
	return diff
}