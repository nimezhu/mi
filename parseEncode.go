package mi

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/nimezhu/netio"
)

func parse(s string) Motif {
	l := strings.Split(s, "\n")
	l1 := strings.Split(l[0], "\t")
	l2 := strings.Split(l1[0], " ")
	m := Motif{}
	m.Id = l2[0]
	m.Size = len(l) - 2
	m.Pwm = make([][]float64, m.Size)
	pwm := make([]float64, 4*m.Size)
	for i := 1; i < len(l)-1; i++ {
		//_, values := iter[i][0], iter[i][1:]
		values := strings.Split(l[i], " ")
		for j := 1; j < len(values); j++ {
			pwm[(i-1)*4+(j-1)], _ = strconv.ParseFloat(values[j], 64)
		}
	}
	for i := range m.Pwm {
		m.Pwm[i] = pwm[i*4 : (i+1)*4]
	}
	return m
}

/*ReadEncodeMotifs Read and Parse motifs from
http://compbio.mit.edu/encode-motifs/
*/
func ReadEncodeMotifs(uri string) ([]Motif, error) {
	f, err := netio.NewReadSeeker(uri)
	if err != nil {
		log.Fatal("can't open file ", err)
		return nil, err
	}
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("can't read file ", err)
		return nil, err
	}
	l := strings.Split(string(buf), ">")
	l = l[1:]
	m := make([]Motif, len(l))
	for i, s := range l {
		m[i] = parse(s)
	}
	return m, nil
}
