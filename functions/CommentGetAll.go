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
	var comments []Comment
	query := datastore.NewQuery(EntityName)
	ids, err := client.GetAll(ctx, query, &comments)
	// 2.1 Iterate and assign IDs to each comments
	for i, comment := range comments {
		comment.ID = ids[i].ID
	}

	// 3. Write into JSON
	byteArray, err := json.Marshal(comments)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// 5. Send response
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, string(byteArray))
}
