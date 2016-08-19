package mi

import (
	"math"
)

func addBg(mat [][]float64, bg []float64) [][]float64 {
	l1 := len(mat)
	m1 := make([][]float64, l1+2)
	m1[0] = bg
	m1[l1+1] = bg
	for i := 0; i < l1; i++ {
		m1[i+1] = mat[i]
	}
	return m1

}
func MotifJsDiv(mat1 [][]float64, mat2 [][]float64, bg []float64) float64 {
	m1 := addBg(mat1, bg)
	m2 := addBg(mat2, bg)
	l1 := len(m1)
	l2 := len(m2)
	mat := jsDivArrays(m1, m2)
	trace := tracer{mat, l1, l2, make([]float64, 0, l1+l2)}
	trace.tracePath(0, 0, 0, 0.0)
	return min(trace.Scores())
}

type tracer struct {
	m      [][]float64
	row    int
	col    int
	scores []float64
}

func (t *tracer) Scores() []float64 {
	return t.scores
}
func (t *tracer) tracePath(i int, j int, s float64, l int) {
	s += t.m[i][j]
	l += 1
	if i == t.row-1 && j == t.col-1 {
		t.scores = append(t.scores, s/float64(l-2))
		return
	}
	if i == 0 && j < t.col-3 { //remove bg vs bg
		t.tracePath(i, j+1, s, l)
	}
	if j == 0 && i < t.row-3 { // remove bg vs bg
		t.tracePath(i+1, j, s, l)
	}
	if i == t.row-1 && j < t.col-1 {
		t.tracePath(i, j+1, s, l)
	}
	if j == t.col-1 && i < t.row-1 {
		t.tracePath(i+1, j, s, l)
	}
	if i < t.row-1 && j < t.col-1 {
		t.tracePath(i+1, j+1, s, l)
	}
	return
}
func jsDivArrays(mat1 [][]float64, mat2 [][]float64) [][]float64 {
	row := len(mat1)
	col := len(mat2)
	M := make([][]float64, row)
	e := make([]float64, row*col)
	for i := range M {
		M[i] = e[i*col : (i+1)*col]
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			M[i][j] = JSDiv(mat1[i], mat2[j])
		}
	}
	return M
}

func jsDivArray(arr [][]float64, bg []float64) []float64 {
	l := len(arr)
	divs := make([]float64, l)
	for i := 0; i < l; i++ {
		divs[i] = JSDiv(arr[i], bg)
	}
	return divs
}

/*
	KL Divergence
	bits
*/
func KLDiv(p []float64, q []float64) (float64, bool) {
	d := 0.0
	ok := true
	for i := 0; i < len(p); i++ {
		if p[i] > 0.0 {
			if q[i] == 0.0 {
				ok = false
			} else {
				d += p[i] * math.Log2(p[i]/q[i])
			}
		}
	}
	return d, ok
}
func JSDiv(p []float64, q []float64) float64 {
	m := meanArray(p, q)
	d1, _ := KLDiv(p, m)
	d2, _ := KLDiv(q, m)
	d := 1.0/2*d1 + 1.0/2*d2
	return d
}
func meanArray(p []float64, q []float64) []float64 {
	l1 := len(p)
	l2 := len(q)
	if l1 != l2 {
		panic("not the same length")
	}
	m := make([]float64, l1)
	for i := 0; i < l1; i++ {
		m[i] = (p[i] + q[i]) / 2
	}
	return m

}
func normToDis(a []float64) []float64 {
	l := len(a)
	p := make([]float64, l)
	if l == 0 {
		return p
	}
	s := sum(a)
	if s == 0 {
		for i := 0; i < l; i++ {
			p[i] = 1.0 / float64(l)
		}
	} else {
		for i := 0; i < l; i++ {
			p[i] = a[i] / s
		}
	}
	return p
}

func sum(a []float64) float64 {
	s := 0.0
	for _, x := range a {
		s += x
	}
	return s
}
func min(a []float64) float64 {
	m := a[0]
	for i := 1; i < len(a); i++ {
		if m > a[i] {
			m = a[i]
		}
	}
	return m
}
