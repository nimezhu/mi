package mi

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/boltdb/bolt"
)

type Motif struct {
	Id   string
	Pwm  [][]float64
	Size int
}

const (
	Bucket = "Motif"
)

func (d *Motif) GobEncode() ([]byte, error) {
	w := new(bytes.Buffer)
	encoder := gob.NewEncoder(w)
	err := encoder.Encode(d.Id)
	if err != nil {
		return nil, err
	}
	err = encoder.Encode(d.Pwm)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func (d *Motif) GobDecode(buf []byte) error {
	r := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(r)
	err := decoder.Decode(&d.Id)
	if err != nil {
		return err
	}
	err = decoder.Decode(&d.Pwm)
	if err != nil {
		return err
	}
	d.Size = len(d.Pwm)
	return nil
}
func (d *Motif) TxtEncode() (string, error) {
	var buffer bytes.Buffer
	buffer.WriteString(d.Id)
	buffer.WriteString("\tA\tC\tG\tT\n")
	for i, r := range d.Pwm {
		buffer.WriteString(strconv.Itoa(int(i) + 1))
		for _, p := range r {
			buffer.WriteString(fmt.Sprintf("\t%f", p))
		}

		buffer.WriteString("\n")
	}
	return buffer.String(), nil
}
func (d *Motif) TxtDecode(a string) error {
	lines := strings.Split(a, "\n")
	size := len(lines) - 1
	d.Size = size
	d.Pwm = make([][]float64, size)
	pwm := make([]float64, 4*size)
	head := strings.Split(lines[0], "\t")
	d.Id = head[0]
	for i := 1; i < len(lines); i++ {
		iter := strings.Split(lines[i], "\t")
		_, values := iter[0], iter[1:]
		for j := 0; j < len(values); j++ {
			pwm[(i-1)*4+j], _ = strconv.ParseFloat(values[j], 64)
		}
	}
	for i := range d.Pwm {
		d.Pwm[i] = pwm[i*4 : (i+1)*4]
	}
	return nil
}

func (t *Motif) SaveToDb(db *bolt.DB) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(Bucket))
		v, _ := t.GobEncode()
		err := b.Put([]byte(t.Id), v)
		return err
	})
	return err
}

func (t *Motif) LoadFromDb(db *bolt.DB, Id string) error {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(Bucket))
		v := b.Get([]byte(Id))
		t.GobDecode(v)
		return nil
	})
	return nil
}

func OpenDb(fn string) (*bolt.DB, error) {
	return bolt.Open(fn, 0600, nil)
}

func InitDb(fn string) error {
	db, err := bolt.Open(fn, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(Bucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	return nil
}
func ListDb(db *bolt.DB) []string {
	s := make([]string, 0)
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(Bucket))
		c0 := b.Cursor()
		for k, _ := c0.First(); k != nil; k, _ = c0.Next() {
			s = append(s, string(k))
		}
		return nil
	})
	return s
}

func IterDb(db *bolt.DB) []*Motif {
	s := make([]*Motif, 0)
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(Bucket))
		c0 := b.Cursor()
		for k, v := c0.First(); k != nil; k, v = c0.Next() {
			m := new(Motif)
			m.GobDecode(v)
			s = append(s, m)
		}
		return nil
	})
	return s
}

func LoadMotifFile(file string) *Motif {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		log.Fatal("can't open file ", err)
	}
	m := Motif{}
	r := csv.NewReader(f)
	r.Comma = '\t'
	iter, err := r.ReadAll()
	if err != nil {
		log.Fatal("error in reading file ", err)
	}

	fn := path.Base(file)
	a := strings.Split(fn, ".")
	m.Id = strings.Join(a[0:len(a)-1], ".")
	m.Size = len(iter) - 1
	m.Pwm = make([][]float64, m.Size)
	pwm := make([]float64, 4*m.Size)
	for i := 1; i < len(iter); i++ {
		_, values := iter[i][0], iter[i][1:]
		for j := 0; j < len(values); j++ {
			pwm[(i-1)*4+j], _ = strconv.ParseFloat(values[j], 64)
		}
	}
	for i := range m.Pwm {
		m.Pwm[i] = pwm[i*4 : (i+1)*4]
	}

	return &m
}
