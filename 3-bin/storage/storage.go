package storage

import (
	"encoding/json"
	"os"

	"github.com/egoisthemain/3-bin/bins"
)

func saveBin(bin bins.Bin, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	binBytes, err := bin.ToBytes()
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, binBytes, 0644)
	return nil
}

func ReadBinList(filename string) (bins.BinList, error) {
	datajson, err := os.ReadFile(filename)
	if err != nil {
		return bins.BinList{}, err
	}
	var binlist bins.BinList
	err = json.Unmarshal(datajson, &binlist)
	if err != nil {
		return bins.BinList{}, err
	}
	return binlist, nil
}

func main() {

}
