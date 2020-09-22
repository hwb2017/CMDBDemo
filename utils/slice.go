package utils

import "reflect"

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

// InterfaceSlice converting sliece of type to slice of interface
func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i:=0; i<s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}