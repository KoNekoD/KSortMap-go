package KSortMap

import "testing"

// TestStringMapSortingHasEffect - tests if the map of strings is sorted correctly(alphabetically)
func TestStringMapSortingHasEffect(t *testing.T) {
	// Arrange
	// Sorted map should be {A: a, B: b, C: c}
	mapForSort := map[string]string{
		"C": "c",
		"A": "a",
		"B": "b",
	}

	// Act
	sortedMap := Sort(mapForSort)

	// Assert
	// When sorting worked correctly keys should not hit: (  C  ->  A[NO_HIT]  ,  A  ->  B[NO_HIT]  ,  B  ->  C[NO_HIT]  )
	mapForSortPseudoIndex := 0
	for sortedMapKey, _ := range sortedMap {

		sortedMapKeyPseudoIndex := 0
		for mapForSortKey, _ := range mapForSort {

			if mapForSortPseudoIndex != sortedMapKeyPseudoIndex {
				sortedMapKeyPseudoIndex++
				continue
			}

			// In theory when maps always sorted alphabetically it will be failed
			if sortedMapKey == mapForSortKey {
				t.Errorf("mapForSort: %v, sortedMap: %v", mapForSort, sortedMap)
			}

			break
		}

		mapForSortPseudoIndex++
	}
}
