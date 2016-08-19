package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"../"
)

func main() {
	//files, _ := ioutil.ReadDir("./data")
	bg := []float64{0.2, 0.3, 0.3, 0.2}
	motifs := make([]*mi.Motif, 0)
	e, _ := ioutil.ReadFile(os.Args[1])
	enriched := strings.Split(string(e), "\n")
	fmt.Println(enriched)
	for _, f := range enriched {
		if f == "" {
			continue
		}
		m := mi.LoadMotifFile(path.Join("./data", f+".txt"))
		motifs = append(motifs, m)
	}
	l := len(motifs)
	fmt.Println(l)
	fmt.Printf("motifs")
	for _, m := range motifs {
		fmt.Printf("\t%s", m.Id)
	}
	fmt.Println("")
	for _, im := range motifs {
		fmt.Printf("%s", im.Id)
		for _, jm := range motifs {
			d := mi.MotifJsDiv(im.Pwm, jm.Pwm, bg)
			fmt.Printf("\t%f", d)
		}
		fmt.Println("")
	}
}
