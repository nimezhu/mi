package turing

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/nimezhu/bed2x"
)

type Beds map[string][]BedI

type BedLine struct {
	chr   string
	start int
	end   int
	line  string
}
type BedSort []BedLine

func (a BedSort) Len() int { return len(a) }
func (a BedSort) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a BedSort) Less(i, j int) bool {
	if a[i].Chr() < a[j].Chr() {
		return true
	}
	if a[i].Chr() > a[j].Chr() {
		return false
	}
	if a[i].Start() < a[j].Start() {
		return true
	}
	if a[i].Start() > a[j].Start() {
		return false
	}
	if a[i].End() < a[j].End() {
		return true
	}
	if a[i].End() > a[j].End() {
		return false
	}
	return false
}
func (a BedLine) Start() int {
	return a.start
}
func (a BedLine) End() int {
	return a.end
}
func (a BedLine) Chr() string {
	return a.chr
}

type BedI interface {
	Chr() string
	Start() int
	End() int
}

func parseBed(line string) BedLine {
	a := strings.Split(line, "\t")
	start, _ := strconv.Atoi(a[1])
	end, _ := strconv.Atoi(a[2])
	return BedLine{a[0], start, end, line}
}

type TuringMap map[string][]Code //Code Need to be Sorted

func ReadBedFileToTuringMap(fn string) TuringMap {
	iter, err := bed2x.IterBed12(fn)
	//	ext := path.Ext(fn)
	count := make(map[string]int)
	bMap := make(map[string][]Code)
	beds := make([]*bed2x.Bed12, 0, 0)

	if err != nil {
		log.Println(err)
		return nil
	}
	for b := range iter {
		if _, ok := count[b.Chr()]; ok {
			count[b.Chr()]++
		} else {
			count[b.Chr()] = 1
		}
		beds = append(beds, b)
	}
	for k, v := range count {
		bMap[k] = make([]Code, v*2)
		count[k] = 0
	}
	for _, bed := range beds {
		k := bed.Chr()
		v, _ := count[k]
		bMap[k][v] = Code{bed.Start(), 1}
		count[k]++
		bMap[k][v+1] = Code{bed.End(), -1}
		count[k]++
	}
	for k := range count {
		sort.Sort(Codes(bMap[k]))
	}
	return bMap
}

func heap(a []Code, b []Code) <-chan Code {
	ch := make(chan Code)
	go func() {
		i := 0
		j := 0
		la := len(a)
		lb := len(b)
		for i < la || j < lb {
			if i == la {
				ch <- b[j]
				j++
			} else if j == lb {
				ch <- a[i]
				i++
			} else if less(a[i], b[j]) {
				ch <- a[i]
				i++
			} else {
				ch <- b[j]
				j++
			}
		}
		close(ch)
	}()
	return ch
}
func JI(a map[string][]Code, b map[string][]Code) float64 {
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
			union += unionLen(a[k], a[k])
		}
		if v == 2 {
			union += unionLen(b[k], b[k])
		}
		if v == 3 {
			overlap += overlapLen(a[k], b[k])
			union += unionLen(a[k], b[k])
		}
	}

	return float64(overlap) / float64(union)
}

func unionLen(a []Code, b []Code) int {
	l := 0
	for v := range unionCodes(a, b) {
		l += v.End() - v.Start()
	}
	return l
}
func overlapLen(a []Code, b []Code) int {
	l := 0
	for v := range overlapCodes(a, b) {
		l += v.End() - v.Start()
	}
	return l
}
func unionCodes(a []Code, b []Code) <-chan RangeI {
	return mergeCodes(a, b, 0)
}

func overlapCodes(a []Code, b []Code) <-chan RangeI {
	return mergeCodes(a, b, 1)
}

/* a and b are sorted Code */
func mergeCodes(a []Code, b []Code, cutoff int) <-chan RangeI {
	ch := make(chan RangeI)
	list := heap(a, b)
	v0 := <-list
	lastPos := v0.Pos
	state := int(v0.Code)
	toggle := false
	if state > cutoff {
		toggle = true
	}
	go func() {
		defer close(ch)
		for v := range list {
			state = state + int(v.Code)
			if toggle && state == cutoff {
				if lastPos != v.Pos {
					ch <- RangeI{lastPos, v.Pos}
				}
				toggle = false
			} else if !toggle && state > cutoff {
				lastPos = v.Pos
				toggle = true
			}
		}
	}()
	return ch

}
