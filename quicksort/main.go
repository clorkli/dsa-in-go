package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	arr := []int{3, 7, 2, 9, 1, 5, 8, 4, 6, 0}
	fmt.Println("原始数组：", arr)

	QuickSort(arr)

	fmt.Println("排序结果：", arr)

	//验证有序性
	if isSorted(arr) {
		fmt.Println("√ 验证通过：数组有序！")
	} else {
		fmt.Println("× 验证失败：数组无序！")
	}
}

func QuickSort(arr []int) {
	quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, left, right int) {
	if left >= right {
		return
	}

	pivotIndex := partition(arr, left, right)
	quickSortHelper(arr, left, pivotIndex-1)
	quickSortHelper(arr, pivotIndex+1, right)
}

func partition(arr []int, left, right int) int {
	randomIndex := left + rand.Intn(right-left+1)
	swap(arr, left, randomIndex)

	pivot := arr[left]
	idx := left + 1

	fmt.Printf("\n--- 分区开始: 范围[%d, %d] Pivot值:%d ---\n", left, right, pivot)

	for i := idx; i <= right; i++ {
		if arr[i] < pivot {
			swap(arr, i, idx)
			idx++
		}
	}

	swap(arr, left, idx-1)

	fmt.Printf("分区结束: 结果 %v, Pivot归位到下标: %d\n", arr, idx-1)

	return idx - 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}