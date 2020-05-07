package sort

import (
	"sort"
)

func BubbleSort(values []int) {
	sort := true
	for sort {
		sort = false
		for i := 0; i < len(values)-1; i++ {
			if values[i] > values[i+1] {
				sort = true
				values[i], values[i+1] = values[i+1], values[i]
			}

		}
	}
}

func Sort(values []int) {
	sort.Ints(values)
}

func GetElements(n int) []int {
	result := make([]int, n)
	j := 0
	for i := n - 1; i > 0; i-- {
		result[j] = i
		j++
	}
	return result
}
