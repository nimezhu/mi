package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/nimezhu/netio"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("DirToDivmat dir(all bed gzip files)")

	}
	files, _ := ioutil.ReadDir(os.Args[1])
	for _, f := range files {
		m := f.Name()
		ext := path.Ext(m)
		if ext == ".gz" || ext == ".bed" {
			log.Println("Processing", m)
			a, _ := netio.ReadAll(m)
			arr := strings.Split(string(a), "\n")
			b := make([]BedLine, len(arr)-1)
			for i, v := range arr {
				if v == "" {
					continue
				}
				b[i] = parseBed(v)
				//fmt.Println(i)
			}
			sort.Sort(BedSort(b))
		}
	}
}
