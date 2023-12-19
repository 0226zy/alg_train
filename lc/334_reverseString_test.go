package lc_test

import (
	"testing"

	"github.com/0226zy/alg_train/lc"

	"github.com/glycerine/goconvey/convey"
)

func TestReverseString(t *testing.T) {
	cases := []struct {
		input  []byte
		output []byte
	}{
		{
			input:  []byte{'h', 'e', 'l', 'l', 'o'},
			output: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			input:  []byte{'h'},
			output: []byte{'h'},
		},
	}
	for _, tcase := range cases {
		convey.Convey("ReverseString", t, func() {
			lc.ReverseString(tcase.input)
			convey.So(tcase.input, convey.ShouldResemble, tcase.output)
		})
	}

}
