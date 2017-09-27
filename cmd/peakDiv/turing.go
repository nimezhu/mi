package main

import "sort"

type BED3 struct {
	chr   string
	start int
	end   int
}

func (b *BED3) Chr() string {
	return b.chr
}
func (b *BED3) Start() int {
	return b.start
}
func (b *BED3) End() int {
	return b.end
}

type code struct {
	pos  int
	code int
}
type codes []code

func (c codes) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c codes) Len() int {
	return len(c)
}
func (c codes) Less(i, j int) bool {
	if c[i].pos < c[j].pos {
		return true
	}
	if c[i].pos > c[j].pos {
		return false
	}
	if c[i].code > c[j].code {
		return true
	}
	if c[i].code < c[j].code {
		return false
	}
	return false
}

func unionLength(beds []BedI) int {
	l := 0
	for v := range union(beds) {
		l += v.End() - v.Start()
	}
	return l
}
func overlapLength(beds []BedI) int {
	l := 0
	for v := range overlap(beds) {
		l += v.End() - v.Start()
	}
	return l
}
func union(beds []BedI) <-chan *BED3 {
	return mergeBed(beds, 0)
}
func overlap(beds []BedI) <-chan *BED3 {
	return mergeBed(beds, 1)
}

/* mergeBed with same chr */
func mergeBed(beds []BedI, cutoff int) <-chan *BED3 {
	ch := make(chan *BED3)
	chr := beds[0].Chr()
	l := make([]code, len(beds)*2)
	for i, v := range beds {
		l[2*i] = code{v.Start(), 1}
		l[2*i+1] = code{v.End(), -1}
	}
	sort.Sort(codes(l))
	//start := l[0].pos
	lastPos := l[0].pos
	state := 0
	toggle := true
	go func() {
		defer close(ch)
		for _, v := range l {
			state = state + v.code
			if toggle && state == cutoff {
				if lastPos != v.pos {
					ch <- &BED3{chr, lastPos, v.pos}
				}
				toggle = false
			} else if !toggle && state > cutoff {
				lastPos = v.pos
				toggle = true
			}
		}
	}()
	return ch

}
