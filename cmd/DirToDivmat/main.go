package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/nimezhu/mi"
)

const (
	program = "DirToDivmat"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("DirToDivmat dir(which store motif files)")

	}
	files, _ := ioutil.ReadDir(os.Args[1])
	bg := []float64{0.2, 0.3, 0.3, 0.2}
	motifs := make([]*mi.Motif, 0)
	for _, f := range files {
		m := mi.LoadMotifFile(path.Join(os.Args[1], f.Name()))
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
