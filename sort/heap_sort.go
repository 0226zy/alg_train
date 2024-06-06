package sort

func HeapSort(nums []int) {

	// 构建一个最大堆
	for i := len(nums)/2 - 1; i >= 0; i-- {
		maxHeap(nums, i, len(nums))
	}
	// 升序，将最大的值交互到数组最后，然后重新调整最大堆
	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		maxHeap(nums, 0, i)
	}
}

func maxHeap(nums []int, curr, heapSize int) {
	child := curr*2 + 1

	for child < heapSize {
		// 检查右节点是否更大
		if child+1 < heapSize && nums[child+1] > nums[child] {
			child++
		}

		// 已经是最大堆
		// 因为我们构建最大堆时是从最后一个非子节点开始，从右往左，从下往上处理所有的非叶子节点
		// curr 有 3 种情况，任何一种都说明 以 curr 为根节点的树已经是最大堆
		if nums[child] < nums[curr] {
			break
		}
		nums[curr], nums[child] = nums[child], nums[curr]
		curr = child
		child = child*2 + 1
	}

}
