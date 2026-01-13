package algo

import (
	"reflect"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	expected := make([]int, len(data))
	copy(expected, data)
	sort.Ints(expected)

	QuickSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, got %v", expected, data)
	}
}

func benchmarkQuickSort(b *testing.B) {
	baseArr := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		baseArr[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr := make([]int, 1000)
		copy(arr, baseArr)
		QuickSort(arr)
	}
}