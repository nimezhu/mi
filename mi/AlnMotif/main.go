package main

import (
	"fmt"
	//"io/ioutil"
	"log"
	"os"
	//"strings"

	"../"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal("AlnMotif MotifAName MotifBName file.db")
	}
	bg := []float64{0.2, 0.3, 0.3, 0.2}
	db, err := mi.OpenDb(os.Args[3])
	if err != nil {
		panic(err)
	}
	m1 := new(mi.Motif)
	m2 := new(mi.Motif)
	m1.LoadFromDb(db, os.Args[1])
	m2.LoadFromDb(db, os.Args[2])
	fmt.Println(mi.MotifJsAlnReport(m1.Pwm, m2.Pwm, bg))
}
