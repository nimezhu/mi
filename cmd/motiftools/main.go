package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/nimezhu/mi"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	//app.Version = mi.VERSION
	app.Name = "motiftools"
	app.Usage = "motiftools"
	app.EnableBashCompletion = true
	// global level flags
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Show more output",
		},
	}

	// Commands
	app.Commands = []cli.Command{
		{
			Name:   "align",
			Usage:  "align two motifs in database",
			Action: CmdAln,
		},
		{
			Name:   "loaddir2db",
			Usage:  "load dir to dbfile",
			Action: CmdLoadDirToDb,
		},
		{
			Name:   "divmat",
			Usage:  "divegence matrix",
			Action: CmdDivmat,
		},
	}
	app.Run(os.Args)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func CmdAln(c *cli.Context) error {
	if c.NArg() < 3 {
		log.Fatal("align MotifAName MotifBName file.db")
	}
	bg := []float64{0.2, 0.3, 0.3, 0.2}
	db, err := mi.OpenDb(c.Args().Get(2))
	if err != nil {
		panic(err)
	}
	m1 := new(mi.Motif)
	m2 := new(mi.Motif)
	err = m1.LoadFromDb(db, c.Args().Get(0))
	checkErr(err)
	err = m2.LoadFromDb(db, c.Args().Get(1))
	checkErr(err)
	fmt.Println(mi.MotifJsAlnReport(m1.Pwm, m2.Pwm, bg))
	return nil
}

func CmdLoadDirToDb(c *cli.Context) error {
	if c.NArg() < 2 {
		log.Fatal("loaddir dir file.db")
	}
	files, _ := ioutil.ReadDir(c.Args().Get(0))
	mi.InitDb(c.Args().Get(1))
	db, err := mi.OpenDb(c.Args().Get(1))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	for _, f := range files {
		m := mi.LoadMotifFile(path.Join(c.Args().Get(0), f.Name()))
		m.SaveToDb(db)
		log.Println("load " + m.Id)

	}
	return nil
}

func CmdDivmat(c *cli.Context) error {
	if c.NArg() < 1 {
		log.Fatal("divmat file.db file.txt[selected motif names]")
	}
	bg := []float64{0.2, 0.3, 0.3, 0.2}
	motifs := make([]*mi.Motif, 0)
	db, err := mi.OpenDb(c.Args().Get(0))
	if err != nil {
		panic(err)
	}

	if c.NArg() == 1 {
		motifs = mi.IterDb(db)
	} else {
		e, _ := ioutil.ReadFile(c.Args().Get(1))
		selected := strings.Split(string(e), "\n")
		for _, f := range selected {
			if f == "" {
				continue
			}
			m := new(mi.Motif)
			m.LoadFromDb(db, f)
			motifs = append(motifs, m)
		}
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
	return nil

}
