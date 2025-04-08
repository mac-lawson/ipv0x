package chain

import (
	"encoding/json"
	"log"
	"os"
)

func saveToJSON(filename string, data any) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func loadFromJSON(filename string) any {
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var data any
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
