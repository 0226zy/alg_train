package main

import (
	"fmt"

	"github.com/0226zy/alg_train/sort"
)

func main() {
	nums := []int{1, 5, 4, 3, 2}
	fmt.Printf("input: %v\n", nums)
	sort.HeapSort(nums)
	fmt.Printf("sort: %v\n", nums)
}
