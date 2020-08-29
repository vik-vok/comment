// Package p contains an HTTP Cloud Function.
package p

import (
	"cloud.google.com/go/datastore"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// CommentGet function returns Comment with given id in JSON format
func CommentGet(w http.ResponseWriter, r *http.Request) {
	// 1. Write ID from request into struct d
	var d struct {
		ID int64 `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
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
	commentKey := datastore.IDKey(EntityName, d.ID, nil)
	err = client.Get(ctx, commentKey, &comment)
	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusNotFound)
		return
	}
	comment.CommentId = d.ID

	// 4. Cast Comment to JSON
	byteArray, err := json.Marshal(comment)
	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 5. Send response
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, string(byteArray))
}
