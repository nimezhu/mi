package main

import (
	"strconv"
	"strings"
)

type BedLine struct {
	chr   string
	start int
	end   int
	line  string
}
type BedSort []BedLine

func (a BedSort) Len() int { return len(a) }
func (a BedSort) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a BedSort) Less(i, j int) bool {
	if a[i].Chr() < a[j].Chr() {
		return true
	}
	if a[i].Chr() > a[j].Chr() {
		return false
	}
	if a[i].Start() < a[j].Start() {
		return true
	}
	if a[i].Start() > a[j].Start() {
		return false
	}
	if a[i].End() < a[j].End() {
		return true
	}
	if a[i].End() > a[j].End() {
		return false
	}
	return false
}
func (a BedLine) Start() int {
	return a.start
}
func (a BedLine) End() int {
	return a.end
}
func (a BedLine) Chr() string {
	return a.chr
}

type BedI interface {
	Chr() string
	Start() int
	End() int
}

func parseBed(line string) BedLine {
	a := strings.Split(line, "\t")
	start, _ := strconv.Atoi(a[1])
	end, _ := strconv.Atoi(a[2])
	return BedLine{a[0], start, end, line}
}
