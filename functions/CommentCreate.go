// Package p contains an HTTP Cloud Function.
package p

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/datastore"
)

// CommentCreate function returns Comment with given id in json format
func CommentCreate(w http.ResponseWriter, r *http.Request) {
	var d struct {
		VoiceID string `json:"voiceID"`
		UserID  string `json:"userID"`
		Text    string `json:"text"`
	}

	// Decode Request into struct
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - " + err.Error()))
		return
	}

	comment := Comment{
		VoiceID: d.VoiceID,
		UserID:  d.UserID,
		Text:    d.Text}

	// 3. connect to database
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "speech-similarity")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - " + err.Error()))
		return
	}

	// store comment entity in database
	key := datastore.IncompleteKey("Comment", nil)
	key, err = client.Put(ctx, key, &comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
