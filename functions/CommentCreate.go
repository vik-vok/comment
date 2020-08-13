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
	var comment Comment

	// 1. Decode Request into Comment struct
	err := json.NewDecoder(r.Body).Decode(&comment)
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
		w.Write([]byte("500 - " + err.Error()))
		return
	}

	// 3. Store comment entity in database
	key := datastore.IncompleteKey(EntityName, nil)
	key, err = client.Put(ctx, key, &comment)
	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - " + err.Error()))
		return
	}

	// 4. Return Status OK (at this point everything is good)
	fmt.Println(w, "Added new Entity with key="+key.String()) /* log new key */
}
