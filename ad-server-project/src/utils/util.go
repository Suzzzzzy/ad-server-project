package utils

import "sort"

// SortMapByValue map 구조에서 value 값으로 sort
func SortMapByValue(m map[int]float64) []struct {
	Key   int
	Value float64
} {
	var sorted []struct {
		Key   int
		Value float64
	}

	for k, v := range m {
		sorted = append(sorted, struct {
			Key   int
			Value float64
		}{k, v})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	return sorted
}