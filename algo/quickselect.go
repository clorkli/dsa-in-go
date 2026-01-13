package algo

import (
	"fmt"
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	targetIndex := len(nums) - k

	rand.Seed(time.Now().UnixNano())

	return quickSelect(nums, 0, len(nums)-1, targetIndex)
}

func quickSelect(nums []int, left, right, targetIndex int) int {
	randomIndex := left + rand.Intn(right-left+1)
	nums[randomIndex], nums[right] = nums[right], nums[randomIndex]

	pivot := nums[right]
	p := left

	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[p], nums[j] = nums[j], nums[p]
			p++
		}
	}

	nums[p], nums[right] = nums[right], nums[p]

	if p == targetIndex {
		return nums[p] 
	} else if p < targetIndex {
		return quickSelect(nums, p+1, right, targetIndex)
	} else {
		return quickSelect(nums, left, p-1, targetIndex)
	}
}

func main() {
	// 预期5
	nums1 := []int{3, 2, 1, 5, 6, 4}
	k1 := 2
	res1 := findKthLargest(nums1, k1)
	fmt.Printf("Arr: %v, K=%d, Result: %d (Expected: 5)\n", []int{3, 2, 1, 5, 6, 4}, k1, res1)

	// 预期4
	nums2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k2 := 4

	tempNums2 := make([]int, len(nums2))
	copy(tempNums2, nums2)
	
	res2 := findKthLargest(tempNums2, k2)
	fmt.Printf("Arr: %v, K=%d, Result: %d (Expected: 4)\n", nums2, k2, res2)
}