// Package p contains an HTTP Cloud Function.
package p

import (
	"cloud.google.com/go/datastore"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// CommentUpdate function returns Comment with given id in json format
func CommentUpdate(w http.ResponseWriter, r *http.Request) {
	// 1. Decode Request into Comment struct
	var req struct {
		ID   int64  `json:"commentId"`
		Text string `json:"text"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_, _ = fmt.Fprint(w, "Error While Parsing Request Body!")
		return
	}

	// 2. Connect to database
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, ProjectName)
	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 3. Get data
	var comment Comment
	commentKey := datastore.IDKey(EntityName, req.ID, nil)
	err = client.Get(ctx, commentKey, &comment)
	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// 4. Update and save
	comment.Text = req.Text
	if _, err := client.Put(ctx, commentKey, &comment); err != nil {
		log.Fatalf("tx.Put: %v", err)
	}

	// 5. Cast Comment to JSON
	byteArray, err := json.Marshal(comment)
	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 6. Send response
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, string(byteArray))
}
