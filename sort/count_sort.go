package sort

// CountSort 适合 nums[x] 最大值和最小值范围不大的数组
func CountSort(nums []int) {
	min, max := nums[0], nums[0]
	// 计算最大值、最小值，确定计数的范围
	for _, num := range nums {
		if min > num {
			min = num
		}
		if max < num {
			max = num
		}
	}

	counts := make([]int, max-min+1)

	// 统计 nums 每个数出现的频率
	for _, num := range nums {
		counts[num-min]++
	}

	// 计算 nums 中每个数在排序后的数组的最后的位置
	// counts[0] 表示有 n 个 0，因此 0 在排序后数组最后的一个位置应该是 n-1
	// count[i+1] 一定排在 count[i] 后面，因此 i+1 的位置等于 count[i+1]+count[i]
	counts[0]--
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	result := make([]int, len(nums))
	for _, num := range nums {

		result[counts[num-min]] = num
		// 如果 counts[x]>0 表示有多个 x，当前已经放好一个，需要-1
		counts[num-min]--
	}
	// 复制回 num
	for i := 0; i < len(result); i++ {
		nums[i] = result[i]
	}
}
