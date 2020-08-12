// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Comment is good
type Comment struct {
	VoiceID string `json:"voiceID"`
	UserID  string `json:"userID"`
	Text    string `json:"text"`
}

// Get function returns Comment with given id in json format
func Get(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Error While Parsing Request Body!")
		return
	}

	comment := Comment{
		VoiceID: "123-456-789",
		UserID:  "ab-cd-ef-gh",
		Text:    "Best Voice Ever!!!"}

	byteArray, err := json.Marshal(comment)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(byteArray))
}
