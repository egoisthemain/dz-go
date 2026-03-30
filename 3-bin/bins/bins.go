package bins

import (
	"errors"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	bins []Bin
}

func newBin(id string, private bool, name string) (*Bin, error) {
	if id == "" {
		return nil, errors.New("wrong id")
	}
	if name == "" {
		return nil, errors.New("wrong name")
	}
	bin := Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
	return &bin, nil
}

func newBinList(bin Bin) (*BinList, error) {
	binList := BinList{
		bins: []Bin{bin},
	}
	return &binList, nil
}
