package messages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type messageString struct {
	To      string `json:"To"`
	From    string `json:"From"`
	Date    string `json:"Date"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
}

type Message struct {
	To      string
	From    string
	Date    time.Time
	Title   string
	Content string
}

func newMessage(mS messageString) *Message {
	const longForm = "January 2, 2006 3:04pm (MST)"
	date, err := time.Parse(longForm, mS.Date)

	if err != nil {
		fmt.Println("Invalid time")
	}

	return &Message{mS.To, mS.From, date, mS.Title, mS.Content}
}

func FromJson(path string) *Message {
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	// Must unmarshall the json object
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var mS messageString
	json.Unmarshal(byteValue, &mS)

	jsonFile.Close()

	return newMessage(mS)
}

func (m Message) String() string {
	return fmt.Sprintf("\nTo: %s\nFrom: %s\nDate: %s\nTitle: %s\nContent: %s\n\n",
		m.To, m.From, m.Date.String(), m.Title, m.Content)
}
