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
func TestCodes(t *testing.T) {
	c1 := []Code{Code{1, 1}, Code{200, -1}, Code{250, 1}, Code{300, -1}}
	c2 := []Code{Code{150, 1}, Code{270, -1}}
	for i := range mergeCodes(c1, c2, 1) {
		t.Log(i)
	}
	for i := range mergeCodes(c1, c2, 0) {
		t.Log(i)
	}
}
func TestReader(t *testing.T) {
	f1 := ReadBedFileToTuringMap("../cmd/peakDiv/ENCFF001UJP.bed.gz")
	f2 := ReadBedFileToTuringMap("../cmd/peakDiv/ENCFF001UJQ.bed.gz")
	for i := range mergeCodes(f1["chr1"], f2["chr1"], 1) {
		t.Log(i)
	}
}
