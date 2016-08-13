package mi

import (
	"log"
	"testing"
)

var (
	m1 = [][]float64{[]float64{0.2, 0.5, 0.2, 0.1},
		[]float64{0.3, 0.2, 0.3, 0.2},
		[]float64{0.2, 0.4, 0.2, 0.2},
	}

	m2 = [][]float64{[]float64{0.7, 0.1, 0.1, 0.1},
		[]float64{0.2, 0.5, 0.2, 0.1},
		[]float64{0.1, 0.1, 0.4, 0.4},
		[]float64{0.1, 0.1, 0.4, 0.4},
	}
	m3 = [][]float64{[]float64{1.0, 0, 0, 0}}
	m4 = [][]float64{[]float64{0.0, 0, 0, 1.0}}
	b1 = []float64{0.2, 0.3, 0.3, 0.2}
	b2 = []float64{0.25, 0.25, 0.25, 0.25}
)

func TestKLDiv(t *testing.T) {
	a, ok := KLDiv([]float64{0.2, 0.5, 0.3}, []float64{0.5, 0.2, 0.3})
	log.Println(a, ok)
}

func TestJSDiv(t *testing.T) {
	js := jsDivArrays(m2, m1)
	log.Println(js)
}

func TestAddBg(t *testing.T) {
	e := addBg(m2, b1)
	log.Println(e)
}
func TestJSDivs(t *testing.T) {
	e1 := addBg(m3, b2)
	e2 := addBg(m4, b2)
	mat := jsDivArrays(e1, e2)
	log.Println(mat)
	trace := tracer{mat, 3, 3, make([]float64, 0, 15)}
	trace.tracePath(0, 0, 0, 0.0)
	log.Println(trace.Scores())
}

func TestMotifJsDiv(t *testing.T) {
	dis := MotifJsDiv(m1, m2, b1)
	log.Println("Distance", dis)
	log.Println("Distance", MotifJsDiv(m1, m1, b1))
	log.Println("Distance A to T", MotifJsDiv(m3, m4, b2))
}
