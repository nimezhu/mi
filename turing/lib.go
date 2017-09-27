package turing

import "sort"

type RangeI struct {
	start int
	end   int
}

func (b *RangeI) Start() int {
	return b.start
}
func (b *RangeI) End() int {
	return b.end
}

type code struct {
	pos  int
	code int8
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

func UnionLength(beds []RangeI) int {
	l := 0
	for v := range union(beds) {
		l += v.End() - v.Start()
	}
	return l
}
func OverlapLength(beds []RangeI) int {
	l := 0
	for v := range overlap(beds) {
		l += v.End() - v.Start()
	}
	return l
}
func union(beds []RangeI) <-chan RangeI {
	return mergeRange(beds, 0)
}
func overlap(beds []RangeI) <-chan RangeI {
	return mergeRange(beds, 1)
}

/* mergeBed with same chr */
func mergeRange(beds []RangeI, cutoff int) <-chan RangeI {
	ch := make(chan RangeI)
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
			state = state + int(v.code)
			if toggle && state == cutoff {
				if lastPos != v.pos {
					ch <- RangeI{lastPos, v.pos}
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

func NewRangeI(s, e int) RangeI {
	return RangeI{s, e}
}
