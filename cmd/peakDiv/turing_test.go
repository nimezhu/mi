package main

import (
	"fmt"
	"testing"
)

func TestMergeBed(t *testing.T) {
	a := make([]BED3, 6)
	a[0] = BED3{"chr1", 100, 200}
	a[1] = BED3{"chr1", 150, 250}
	a[2] = BED3{"chr1", 300, 400}
	a[3] = BED3{"chr1", 100, 200}
	a[4] = BED3{"chr1", 200, 320}
	a[5] = BED3{"chr1", 380, 500}
	fmt.Println("union")
	for v := range mergeBed(a, 0) {
		fmt.Println(v)
	}
	fmt.Println("overlap")
	for v := range mergeBed(a, 1) {
		fmt.Println(v)
	}
	fmt.Println("overlap2")
	for v := range mergeBed(a, 2) {
		fmt.Println(v)
	}

}
