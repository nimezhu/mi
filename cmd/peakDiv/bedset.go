package main

import (
	"path"
	"strings"

	"github.com/nimezhu/netio"
)

type BedMap map[string][]RangeI

func ReadBedFile(fn string) BedMap {
	ext := path.Ext(fn)
	count := make(map[string]int)
	bMap := make(map[string][]RangeI)
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
		bMap[k] = make([]RangeI, v)
		count[k] = 0
	}
	for _, bed := range b {
		k := bed.Chr()
		v, _ := count[k]
		bMap[k][v] = RangeI{bed.Start(), bed.End()}
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
			union += unionLength(a[k])
		}
		if v == 2 {
			union += unionLength(b[k])
		}
		if v == 3 {
			c := make([]RangeI, len(a[k])+len(b[k]))
			for i, v := range a[k] {
				c[i] = v
			}
			l := len(a[k])
			for i, v := range b[k] {
				c[i+l] = v
			}
			overlap += overlapLength(c)
			union += unionLength(c)

		}
	}
	return float64(overlap) / float64(union)
}
