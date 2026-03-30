package files

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func WriteFile() {

}

func ReadFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func ReadJsonFile(filename string) bool {
	data, err := os.ReadFile(filename)
	if err != nil {
		return false
	}
	ext := filepath.Ext(filename)
	if ext == ".json" && json.Valid(data) {
		return true
	} else {
		return false
	}
}

func main() {

}
