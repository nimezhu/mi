package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/nimezhu/mi"
)

var (
	program = "MotifDirToDb"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("MotifDirToDb dir(which store motif files) file.db")
	}
	files, _ := ioutil.ReadDir(os.Args[1])
	mi.InitDb(os.Args[2])
	db, err := mi.OpenDb(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer db.Close()
	for _, f := range files {
		m := mi.LoadMotifFile(path.Join(os.Args[1], f.Name()))
		m.SaveToDb(db)
		log.Println("load " + m.Id)

	}

}
