package cmd

import "sort"

// KSortMap - sorts a map by it's keys
func KSortMap[K ~int | ~int64 | ~string, V comparable](input map[K]V) map[K]V {
	keys := make([]K, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	result := make(map[K]V)
	for _, k := range keys {
		result[k] = input[k]
	}

	return result
}
