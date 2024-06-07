package sort

func MergeSort(nums []int) {
	result := make([]int, len(nums))
	mergeSort(nums, 0, len(nums)-1, result)
}

func mergeSort(nums []int, start, end int, result []int) {
	if start == end {
		return
	}
	mid := (start + end) / 2
	mergeSort(nums, start, mid, result)
	mergeSort(nums, mid+1, end, result)
	merge(nums, start, end, result)
}

func merge(nums []int, start, end int, result []int) {
	mid := (start + end) / 2
	idx1, end1 := start, mid
	idx2, end2 := mid+1, end
	ridx := idx1

	for idx1 <= end1 && idx2 <= end2 {
		if nums[idx1] <= nums[idx2] {
			result[ridx] = nums[idx1]
			ridx++
			idx1++
		} else {
			result[ridx] = nums[idx2]
			ridx++
			idx2++
		}
	}
	for ; idx1 <= end1; idx1++ {
		result[ridx] = nums[idx1]
		ridx++
	}

	for ; idx2 <= end2; idx2++ {
		result[ridx] = nums[idx2]
		ridx++
	}
	for i := start; i <= end; i++ {
		nums[i] = result[i]
	}
}
