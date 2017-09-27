package main

import (
	"fmt"
	"testing"
)

func TestMergeBed(t *testing.T) {
	a := make([]*BED3, 6)
	a[0] = &BED3{"chr1", 100, 200}
	a[1] = &BED3{"chr1", 150, 250}
	a[2] = &BED3{"chr1", 300, 400}
	a[3] = &BED3{"chr1", 100, 200}
	a[4] = &BED3{"chr1", 200, 320}
	a[5] = &BED3{"chr1", 380, 500}
	b := make([]BedI, 6)
	for i, v := range a {
		b[i] = v
	}
	//b := []BedI{a[0], a[1], a[2], a[3], a[4], a[5]}
	fmt.Println("union")
	for v := range mergeBed(b, 0) {
		fmt.Println(v)
	}
	fmt.Println("overlap")
	for v := range mergeBed(b, 1) {
		fmt.Println(v)
	}
	fmt.Println("overlap2")
	for v := range mergeBed(b, 2) {
		fmt.Println(v)
	}

}
