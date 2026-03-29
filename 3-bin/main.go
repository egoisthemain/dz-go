package main

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	bins []Bin
}

func newBin(id string, private bool, name string) {}

func newBinList(bin Bin) {}

func main() {

}
