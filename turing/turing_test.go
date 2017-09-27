package turing

import (
	"fmt"
	"testing"
)

func TestMergeBed(t *testing.T) {
	a := make([]RangeI, 6)
	a[0] = RangeI{100, 200}
	a[1] = RangeI{150, 250}
	a[2] = RangeI{300, 400}
	a[3] = RangeI{100, 200}
	a[4] = RangeI{200, 320}
	a[5] = RangeI{380, 500}

	fmt.Println("union")
	for v := range mergeRange(a, 0) {
		fmt.Println(v)
	}
	fmt.Println("overlap")
	for v := range mergeRange(a, 1) {
		fmt.Println(v)
	}
	fmt.Println("overlap2")
	for v := range mergeRange(a, 2) {
		fmt.Println(v)
	}

}
