package main

import (
	"fmt"
	"os"

	"github.com/nimezhu/mi"
	"github.com/urfave/cli"
)

const (
	uri = "http://compbio.mit.edu/encode-motifs/motifs.txt"
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
func CmdDivmat(c *cli.Context) error {
	bg := []float64{0.2, 0.3, 0.3, 0.2}
	motifs, err := mi.ReadEncodeMotifs(uri)
	checkErr(err)
	fmt.Printf("motifs")
	for _, m := range motifs {
		fmt.Printf("\t%s", m.Id)
	}
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
