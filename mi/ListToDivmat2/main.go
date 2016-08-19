package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"../"
)

func main() {
	//files, _ := ioutil.ReadDir("./data")
	if len(os.Args) < 3 {
		log.Fatal("ListToDivmat2 file.list file.db")
	}
	bg := []float64{0.2, 0.3, 0.3, 0.2}
	motifs := make([]*mi.Motif, 0)
	e, _ := ioutil.ReadFile(os.Args[1])
	enriched := strings.Split(string(e), "\n")
	//fmt.Println(enriched)
	db, err := mi.OpenDb(os.Args[2])
	if err != nil {
		panic(err)
	}
	for _, f := range enriched {
		if f == "" {
			continue
		}
		m := new(mi.Motif)
		m.LoadFromDb(db, f)
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
