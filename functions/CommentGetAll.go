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

// CommentGetAll function returns Comment with given id in json format
func CommentGetAll(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Error While Parsing Request Body!")
		return
	}

	comment1 := Comment{
		VoiceID: "123",
		UserID:  "ab",
		Text:    "Best Voice1 Ever!!!"}

	comment2 := Comment{
		VoiceID: "456",
		UserID:  "cd",
		Text:    "Best Voice2 Ever!!!"}

	comment3 := Comment{
		VoiceID: "789",
		UserID:  "ef",
		Text:    "Best Voice3 Ever!!!"}

	comments := [3]Comment{comment1, comment2, comment3}

	byteArray, err := json.Marshal(comments)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(byteArray))
}
