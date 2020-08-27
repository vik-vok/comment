// Package p contains an HTTP Cloud Function.
package p

import (
	"cloud.google.com/go/datastore"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// CommentCreate function returns Comment with given id in json format
func CommentCreate(w http.ResponseWriter, r *http.Request) {
	// 1. Decode Request into Comment struct
	var comment Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("500 - " + err.Error()))
		return
	}
	comment.Created = time.Now()
	comment.ID = nil;

	// 2. Connect to database
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, ProjectName)
	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("500 - " + err.Error()))
		return
	}

	// 3. Store comment entity in database
	key := datastore.IncompleteKey(EntityName, nil)
	key, err = client.Put(ctx, key, &comment)
	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("500 - " + err.Error()))
		return
	}

	// 4. Cast Comment to JSON
	comment.ID = key.ID
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
