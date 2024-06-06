package sort_test

import (
	"fmt"
	"testing"

	"github.com/0226zy/alg_train/sort"
	"github.com/glycerine/goconvey/convey"
)

func TestHeapSort(t *testing.T) {
	sort_cases := []struct {
		input  []int
		expect []int
	}{

		{
			input:  []int{5, 60, 38, 41, 88, 44, 85, 93, 41, 75},
			expect: []int{5, 38, 41, 41, 44, 60, 75, 85, 88, 93},
		},
		{
			input:  []int{65, 5, 39, 61, 51, 85, 31, 87, 28, 98},
			expect: []int{5, 28, 31, 39, 51, 61, 65, 85, 87, 98},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			expect: []int{1, 2, 3, 4, 5},
		},
		{
			input:  []int{1, 5, 4, 3, 2},
			expect: []int{1, 2, 3, 4, 5},
		},
		{
			input:  []int{5, 4, 3, 2, 1},
			expect: []int{1, 2, 3, 4, 5},
		},
	}

	for _, sort_case := range sort_cases {
		convey.Convey("QuickSort:", t, func() {
			fmt.Println(sort_case.input)
			sort.HeapSort(sort_case.input)
			convey.Convey("input:"+toString(sort_case.input), func() {
				convey.So(sort_case.input, convey.ShouldResemble, sort_case.expect)
			})
		})
	}
}
