package services

import (
	"testing"
	"testing-golang/src/api/utils/sort"
)

func TestSort(t *testing.T) {
	values := []int{91, 34, 7, 87, 45, 25, 67, 24}
	Sort(values)
	if values[0] != 7 {
		t.Error("Something went wrong")
	}

	if values[len(values)-1] != 91 {
		t.Error("Something went wrong")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	values := sort.GetElements(10000)
	for i := 0; i < b.N; i++ {
		sort.BubbleSort(values)
	}

}

func BenchmarkSort(b *testing.B) {
	values := sort.GetElements(10000)
	for i := 0; i < b.N; i++ {
		sort.BubbleSort(values)
	}

}
