package sort

import (
	"math/rand"
	"time"
)

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	rand.Seed(time.Now().UnixNano())
	mid := partition(nums, start, end)
	quickSort(nums, start, mid-1)
	quickSort(nums, mid+1, end)
}

func partition(nums []int, start, end int) int {
	idx := rand.Intn(end-start+1) + start
	base := nums[idx]
	nums[idx], nums[start] = nums[start], nums[idx]
	left, right := start+1, end
	for left < right {
		for left < right && nums[left] <= base {
			left++
		}
		for left < right && nums[right] >= base {
			right--
		}
		if left != right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	if left == right && nums[right] > base {
		right--
	}

	nums[start], nums[right] = nums[right], nums[start]
	return right
}
