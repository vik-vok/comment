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

// CommentGet function returns Comment with given id in JSON format
func CommentOriginalVoicesGet(w http.ResponseWriter, r *http.Request) {
	// 1. Write ID from request into struct d
	var d struct {
		ID string `json:"voiceId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		d.ID = r.FormValue("voiceId")
		if d.ID == "" {
			_, _ = fmt.Fprint(w, "Error While Parsing Request Body!\n URL: "+r.URL.String())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
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
	var comments []Comment
	query := datastore.NewQuery(EntityName).Filter("VoiceID =", d.ID)
	ids, err := client.GetAll(ctx, query, &comments)
	// 2.1 Iterate and assign IDs to each comments
	for i, _ := range comments {
		comments[i].CommentId = ids[i].ID
	}
	// 2.2 Sort with created date
	sort.Slice(comments, func(i, j int) bool {
		return comments[i].Created.Before(comments[j].Created)
	})

	// 4. Cast Comment to JSON
	byteArray, err := json.Marshal(comments)
	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 5. Send response
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, string(byteArray))
}
