package mi

import (
	"os"
	"testing"
)

var (
	m1 = [][]float64{
		[]float64{0.3, 0.2, 0.3, 0.2},
		[]float64{0.2, 0.4, 0.2, 0.2},
		[]float64{0.1, 0.1, 0.4, 0.4},
		[]float64{0.2, 0.4, 0.2, 0.2},
		[]float64{0.2, 0.4, 0.2, 0.2},
		[]float64{0.2, 0.4, 0.2, 0.2},
	}

	m2 = [][]float64{
		[]float64{0.7, 0.1, 0.1, 0.1},
		[]float64{0.2, 0.5, 0.2, 0.1},
		[]float64{0.3, 0.2, 0.3, 0.2},
		[]float64{0.2, 0.4, 0.2, 0.2},
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
	t.Log(a, ok)
}

func TestJSDiv(t *testing.T) {
	js := jsDivArrays(m2, m1)
	t.Log(js)
}

func TestAddBg(t *testing.T) {
	e := addBg(m2, b1)
	t.Log(e)
}
func TestJSDivs(t *testing.T) {
	e1 := addBg(m3, b2)
	e2 := addBg(m4, b2)
	mat := jsDivArrays(e1, e2)
	t.Log(mat)
	trace := tracer{mat, 3, 3, make([]float64, 0, 15)}
	trace.tracePath(0, 0, 0, 0.0)
	t.Log(trace.Scores())
}

func TestMotifJsDiv(t *testing.T) {
	dis := MotifJsDiv(m1, m2, b1)
	t.Log("Distance", dis)
	t.Log("Distance", MotifJsDiv(m1, m1, b1))
	t.Log("Distance A to T", MotifJsDiv(m3, m4, b2))
}
func TestReadMotif(t *testing.T) {
	ml1 := LoadMotifFile("../data/M3624_1.02.txt")
	ml2 := LoadMotifFile("../data/M3625_1.02.txt")
	t.Log("Distance M1 to M2", MotifJsDiv(ml1.Pwm, ml2.Pwm, b2))

	t.Log(ml1)
	m2, _ := ml2.TxtEncode()
	t.Log(m2)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func TestMotifDb(t *testing.T) {
	ml1 := LoadMotifFile("../data/M3624_1.02.txt")
	ml2 := LoadMotifFile("../data/M3625_1.02.txt")
	err := InitDb("test.db")
	checkErr(err)
	db, err := OpenDb("test.db")
	checkErr(err)
	err = ml1.SaveToDb(db)
	checkErr(err)
	err = ml2.SaveToDb(db)
	checkErr(err)
	m3 := new(Motif)
	m3.LoadFromDb(db, "M3624_1.02")
	t.Log(m3.TxtEncode())
	ms := IterDb(db)
	for _, v := range ms {
		t.Log(v.TxtEncode())
		m := new(Motif)
		str, _ := v.TxtEncode()
		err := m.TxtDecode(str)
		checkErr(err)
		t.Log(m.TxtEncode())
	}

	os.Remove("test.db")

}

func TestMotifAln(t *testing.T) {
	bg := []float64{0.2, 0.3, 0.3, 0.2}
	ml1 := LoadMotifFile("../data/M3624_1.02.txt")
	ml2 := LoadMotifFile("../data/M3625_1.02.txt")
	score, offset := MotifJsAln(ml1.Pwm, ml2.Pwm, bg)
	t.Log(score, offset)
}

func TestMotifAln2(t *testing.T) {
	bg := []float64{0.2, 0.3, 0.3, 0.2}
	t.Log(m1)
	t.Log(m2)
	score, offset := MotifJsAln(m1, m2, bg)
	t.Log(score, offset)
	score, offset = MotifJsAln(m2, m1, bg)
	t.Log(MotifJsAlnReport(m2, m1, bg))
	t.Log(MotifJsAlnReport(m1, m2, bg))
	t.Log(score, offset)
}
