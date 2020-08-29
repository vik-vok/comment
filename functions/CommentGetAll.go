// Package p contains an HTTP Cloud Function.
package p

import (
	"cloud.google.com/go/datastore"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
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
	for i, _ := range comments {
		comments[i].commentId = ids[i].ID
	}
	// 2.2 Sort with created date
	sort.Slice(comments, func(i, j int) bool {
		return comments[i].Created.Before(comments[j].Created)
	})
	fmt.Printf("%+v\n", comments)


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
