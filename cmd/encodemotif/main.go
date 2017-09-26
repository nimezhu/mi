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
	fmt.Printf("\n")
	l := len(motifs)
	matarray := make([]float64, l*l)
	for i := 0; i < len(motifs); i++ {
		matarray[i*l+i] = float64(0.0)
		for j := i + 1; j < len(motifs); j++ {
			v := mi.MotifJsDiv(motifs[i].Pwm, motifs[j].Pwm, bg)
			matarray[i*l+j] = v
			matarray[j*l+i] = v
		}
	}
	for i := 0; i < len(motifs); i++ {
		fmt.Printf("%s", motifs[i].Id)
		for j := 0; j < len(motifs); j++ {
			fmt.Printf("\t%f", matarray[i*l+j])
		}
		fmt.Printf("\n")
	}
	return nil

}
