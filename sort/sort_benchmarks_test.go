package sort

import (
	"math/rand"
	"testing"
	"time"
)

func genRandomSlice(size int) []int {
	return []int{3, 4, 5}
	slice := make([]int, size)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(size * 10) // [0,size*10]
	}
	return slice
}

func BenchmarkQuickSort(b *testing.B) {

	slice := genRandomSlice(10)

	// reset time
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 去除复制时间
		b.StopTimer()
		tmp := make([]int, len(slice))
		copy(tmp, slice)
		b.StartTimer()

		QuickSort(slice)
	}

}
