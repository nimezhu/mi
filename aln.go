package mi

import (
	"bytes"
	"fmt"
)

type aln struct {
	m       [][]float64
	row     int
	col     int
	scores  []float64
	offsets []int
}

func (t *aln) Scores() []float64 {
	return t.scores
}
func (t *aln) AlnDiv() (float64, int) {
	offset := t.offsets[0]
	score := t.scores[0]
	for i := 1; i < len(t.scores); i++ {
		if score > t.scores[i] {
			score = t.scores[i]
			offset = t.offsets[i]
		}
	}
	return score, offset

}
func (t *aln) tracePath(i int, j int, s float64, l int, offset int) {
	s += t.m[i][j]
	l += 1
	if i == t.row-1 && j == t.col-1 {
		t.scores = append(t.scores, s/float64(l-2))
		t.offsets = append(t.offsets, offset)
		return
	}
	if i == 0 && j < t.col-3 { //remove bg vs bg
		t.tracePath(i, j+1, s, l, offset)
	}
	if j == 0 && i < t.row-3 { // remove bg vs bg
		t.tracePath(i+1, j, s, l, offset)
	}
	if i == t.row-1 && j < t.col-1 {
		t.tracePath(i, j+1, s, l, offset)
	}
	if j == t.col-1 && i < t.row-1 {
		t.tracePath(i+1, j, s, l, offset)
	}
	if i < t.row-1 && j < t.col-1 {
		t.tracePath(i+1, j+1, s, l, i-j)
	}
	return
}

func MotifJsAln(mat1 [][]float64, mat2 [][]float64, bg []float64) (float64, int) {
	m1 := addBg(mat1, bg)
	m2 := addBg(mat2, bg)
	l1 := len(m1)
	l2 := len(m2)
	mat := jsDivArrays(m1, m2)
	trace := aln{mat, l1, l2, make([]float64, 0, l1+l2), make([]int, 0, l1+l2)}
	trace.tracePath(0, 0, 0, 0.0, -1000)
	return trace.AlnDiv()
}

func pos(i int, l int, mat [][]float64) string {
	//log.Printf(i)
	if i < 0 || i >= l {
		return "NA"
	} else {
		return fmt.Sprintf("%f", mat[i])
	}
}
func MotifJsAlnReport(mat1 [][]float64, mat2 [][]float64, bg []float64) string {
	m1 := addBg(mat1, bg)
	m2 := addBg(mat2, bg)
	l1 := len(m1)
	l2 := len(m2)
	mat := jsDivArrays(m1, m2)
	trace := aln{mat, l1, l2, make([]float64, 0, l1+l2), make([]int, 0, l1+l2)}
	trace.tracePath(0, 0, 0, 0.0, -1000)
	var buffer bytes.Buffer
	score, offset := trace.AlnDiv()
	buffer.WriteString(fmt.Sprintf("Mininum JS Divengence: %f\n", score))
	buffer.WriteString(fmt.Sprintf("Based on background: %f\n", bg))
	i := 0
	if offset > 0 {
		i = -offset
	}
	for ; i+offset < l1-2 || i < l2-2; i++ {
		buffer.WriteString(fmt.Sprintf("%d %s\t%d %s\n", i+offset, pos(i+offset, l1-2, mat1), i, pos(i, l2-2, mat2)))
	}
	return buffer.String()

}
