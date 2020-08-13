// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CommentCreate function returns Comment with given id in json format
func CommentCreate(w http.ResponseWriter, r *http.Request) {
	var d struct {
		VoiceID string `json:"voiceID"`
		UserID  string `json:"userID"`
		Text    string `json:"text"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Error While Parsing Request Body!")
		return
	}

	comment := Comment{
		VoiceID: d.VoiceID,
		UserID:  d.UserID,
		Text:    d.Text}

	byteArray, err := json.Marshal(comment)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(byteArray))
}
