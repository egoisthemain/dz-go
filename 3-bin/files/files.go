package files

import (
	"encoding/json"
	"os"
)

func WriteFile() {

}

func ReadJsonFile(filename string) bool {
	data, err := os.ReadFile(filename)
	if err != nil {
		return false
	}
	return json.Valid(data)
}

func main() {

}
