package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("DirToDivmat dir(all bed gzip files)")

	}
	files, _ := ioutil.ReadDir(os.Args[1])
	//mat := make([][]float64, len(files))

	bedFiles := make([]string, 0)
	for _, f := range files {
		m := f.Name()
		ext := path.Ext(m)
		if ext == ".bed" || ext == ".gz" {
			bedFiles = append(bedFiles, f.Name())
		}
	}
	l := len(bedFiles)
	bedMap := make([]BedMap, l)
	ids := make([]string, l)
	arr := make([]float64, l*l)
	for i, f := range bedFiles {
		bedMap[i] = ReadBedFile(f)
		n := strings.Replace(f, ".gz", "", -1)
		ids[i] = strings.Replace(n, ".bed", "", -1)
		if i%10 == 0 {
			log.Printf("reading %d %s", i, f)
		}
	}
	for i := 0; i < l; i++ {
		arr[i*l+i] = float64(1.0)
		log.Printf("calc %d %s\n", i, ids[i])
		//ba := ReadBedFile(bedFiles[i])
		for j := i + 1; j < l; j++ {
			//bb := ReadBedFile(bedFiles[j])
			if j%10 == 0 {
				log.Printf("%d ", j)
			}
			f := JaccardIndex(bedMap[i], bedMap[j])
			arr[i*l+j] = f
			arr[j*l+i] = f

		}
		log.Printf("\n")
	}
	fmt.Printf("IDs\t")
	fmt.Printf(strings.Join(ids, "\t"))
	fmt.Printf("\n")
	for i := 0; i < l; i++ {
		fmt.Printf("%s", ids[i])
		for j := 0; j < l; j++ {
			fmt.Printf("\t%f", arr[i*l+j])
		}
		fmt.Printf("\n")
	}
}
