package bins

import (
	"encoding/json"
	"errors"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func newBin(id string, private bool, name string) (*Bin, error) {
	if id == "" {
		return nil, errors.New("wrong id")
	}
	if name == "" {
		return nil, errors.New("wrong name")
	}
	bin := Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
	return &bin, nil
}

func newBinList(bin Bin) (*BinList, error) {
	binList := BinList{
		Bins: []Bin{bin},
	}
	return &binList, nil
}

func (bin *Bin) ToBytes() ([]byte, error) {
	file, err := json.Marshal(bin)
	if err != nil {
		return nil, err
	}
	return file, nil
}
