package messages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Message struct {
	To      string    `json:"To"`
	From    string    `json:"From"`
	Date    time.Time `json:"Date"`
	Title   string    `json:"Title"`
	Content string    `json:"Content"`
}

func FromJson(path string) *Message {
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	// Must unmarshall the json object
	byteValue, _ := ioutil.ReadAll(jsonFile)

	m := new(Message)
	json.Unmarshal(byteValue, m)

	jsonFile.Close()

	return m
}
