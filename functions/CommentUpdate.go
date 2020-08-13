// Package p contains an HTTP Cloud Function.
package p

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/datastore"
)

// CommentUpdate function returns Comment with given id in json format
func CommentUpdate(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	}

	// 1. Decode Request into Comment struct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - " + err.Error()))
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

	// 3. Get data from database
	// [START datastore_keys_only_query]
	query := datastore.NewQuery(EntityName)
	// [END datastore_keys_only_query]

	keys, err := client.GetAll(ctx, query, nil)
	fmt.Println(keys)

	// 3. Store comment entity in database
	// commentKey := datastore.NewQuery("Task")
	// fmt.Println(commentKey)

	// tx, err := client.NewTransaction(ctx)
	// if err != nil {
	// 	log.Fatalf("client.NewTransaction: %v", err)
	// }
	// var comment Comment
	// if err := tx.Get(commentKey, &comment); err != nil {
	// 	log.Fatalf("tx.Get: %v", err)
	// }
	// comment.Text = req.Text
	// if _, err := tx.Put(commentKey, &comment); err != nil {
	// 	log.Fatalf("tx.Put: %v", err)
	// }
	// if _, err := tx.Commit(); err != nil {
	// 	log.Fatalf("tx.Commit: %v", err)
	// }

	// 4. Return Status OK (at this point everything is good)
	w.WriteHeader(http.StatusOK)
}
