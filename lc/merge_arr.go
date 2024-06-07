package lc

import "fmt"

func Merge(A []int, m int, B []int, n int) {
	merge(A, m, B, n)
	return
	idx1, idx2 := 0, 0
	for idx1 < m+idx2 && idx2 < n {
		if A[idx1] >= B[idx2] {
			idx := idx2 + m
			fmt.Println(1, idx1, idx2, idx)
			for idx > idx1 {
				A[idx] = A[idx-1]
				idx--
			}
			A[idx] = B[idx2]
			idx2++
		}
		idx1++
	}
	for ; idx2 < n; idx2++ {
		A[idx1] = B[idx2]
		idx1++
	}
}

func merge(A []int, m int, B []int, n int) {
	size := m + n
	idx1, idx2 := m-1, n-1
	for idx2 >= 0 {
		if idx1 >= 0 && A[idx1] >= B[idx2] {
			size--
			A[size] = A[idx1]
			idx1--
		} else {
			size--
			A[size] = B[idx2]
			idx2--
		}
	}

}
