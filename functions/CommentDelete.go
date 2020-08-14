// Package p contains an HTTP Cloud Function.
package p

import (
	"cloud.google.com/go/datastore"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// CommentDelete function deletes comment
func CommentDelete(w http.ResponseWriter, r *http.Request) {
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

	// 3. Delete comment
	commentKey := datastore.IDKey(EntityName, d.ID, nil)
	err = client.Delete(ctx, commentKey)
	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
