package main

import (
	"path"
	"strings"

	"github.com/nimezhu/mi/turing"
	"github.com/nimezhu/netio"
)

type BedMap map[string][]turing.RangeI

func ReadBedFile(fn string) BedMap {
	ext := path.Ext(fn)
	count := make(map[string]int)
	bMap := make(map[string][]turing.RangeI)
	var b []BedLine
	if ext == ".gz" || ext == ".bed" {
		//log.Println("Processing", m)
		a, _ := netio.ReadAll(fn)
		arr := strings.Split(string(a), "\n")
		b = make([]BedLine, len(arr)-1) //TODO
		for i, v := range arr {
			if v == "" {
				continue
			}
			b[i] = parseBed(v)
			if _, ok := count[b[i].Chr()]; ok {
				count[b[i].Chr()]++
			} else {
				count[b[i].Chr()] = 1
			}
		}
	}
	for k, v := range count {
		bMap[k] = make([]turing.RangeI, v)
		count[k] = 0
	}
	for _, bed := range b {
		k := bed.Chr()
		v, _ := count[k]
		bMap[k][v] = turing.NewRangeI(bed.Start(), bed.End())
		count[k]++
	}
	return bMap
}

func JaccardIndex(a BedMap, b BedMap) float64 {
	chrs := make(map[string]int)
	for k, _ := range a {
		chrs[k] = 1
	}
	for k, _ := range b {
		if _, ok := chrs[k]; ok {
			chrs[k] += 2
		} else {
			chrs[k] = 2
		}
	}
	union := 0
	overlap := 0
	for k, v := range chrs {
		if v == 1 {
			union += turing.UnionLength(a[k])
		}
		if v == 2 {
			union += turing.UnionLength(b[k])
		}
		if v == 3 {
			c := make([]turing.RangeI, len(a[k])+len(b[k]))
			for i, v := range a[k] {
				c[i] = v
			}
			l := len(a[k])
			for i, v := range b[k] {
				c[i+l] = v
			}
			overlap += turing.OverlapLength(c)
			union += turing.UnionLength(c)

		}
	}
	return float64(overlap) / float64(union)
}
