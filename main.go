package main

import (
	"fmt"

	"github.com/0226zy/alg_train/lc"
	"github.com/0226zy/alg_train/sort"
)

func testLc() {
	A := make([]int, 20)
	A[0] = 2
	A[1] = 5
	A[2] = 8
	A[3] = 9
	A[4] = 9
	A[5] = 10
	B := []int{1, 3, 4, 8, 10}
	lc.Merge(A, 6, B, 5)
	fmt.Println(A)
}
func main() {
	testLc()
	return
	nums := []int{1, 5, 4, 3, 2}
	fmt.Printf("input: %v\n", nums)
	sort.HeapSort(nums)
	fmt.Printf("sort: %v\n", nums)
}
