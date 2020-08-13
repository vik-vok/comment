// Package p contains an HTTP Cloud Function.
package p

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/datastore"
)

// CommentGetAll function returns all Comments in json format
func CommentGetAll(w http.ResponseWriter, r *http.Request) {
	// 1. Connect to database
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, ProjectName)
	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 2. Get data from database
	query := datastore.NewQuery(EntityName)
	var comments []Comment
	_, err = client.GetAll(ctx, query, &comments)

	// 3. Write into JSON
	byteArray, err := json.Marshal(comments)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(byteArray))
}
