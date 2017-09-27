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

type Code struct {
	Pos  int
	Code int
}
type Codes []Code

func (c Codes) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c Codes) Len() int {
	return len(c)
}
func (c Codes) Less(i, j int) bool {
	return less(c[i], c[j])
}
func less(a Code, b Code) bool {
	if a.Pos < b.Pos {
		return true
	}
	if a.Pos > b.Pos {
		return false
	}
	if a.Code > b.Code {
		return true
	}
	if a.Code < b.Code {
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
	l := make([]Code, len(beds)*2)
	for i, v := range beds {
		l[2*i] = Code{v.Start(), 1}
		l[2*i+1] = Code{v.End(), -1}
	}
	sort.Sort(Codes(l))
	//start := l[0].Pos
	lastPos := l[0].Pos
	state := 0
	toggle := false
	if state > cutoff {
		toggle = true
	}
	go func() {
		defer close(ch)
		for _, v := range l {
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

func NewRangeI(s, e int) RangeI {
	return RangeI{s, e}
}
