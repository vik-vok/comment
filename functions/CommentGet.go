// Package p contains an HTTP Cloud Function.
package p

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"cloud.google.com/go/datastore"
)

// CommentGet function returns Comment with given id in json format
func CommentGet(w http.ResponseWriter, r *http.Request) {
	var d struct {
		ID int64 `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		_, _ = fmt.Fprint(w, "Error While Parsing Request Body!")
		return
	}

	// 1. Connect to database
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, ProjectName)
	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 2. Get data from database
	var comment Comment
	commentKey := datastore.NameKey(EntityName, strconv.FormatInt(d.ID, 10), nil)
	err = client.Get(ctx, commentKey, &comment)
	comment.ID = d.ID

	byteArray, err := json.Marshal(comment)
	if err != nil {
		fmt.Println(err)
	}

	_, _ = fmt.Fprint(w, string(byteArray))
}
