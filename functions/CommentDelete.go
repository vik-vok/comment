// Package p contains an HTTP Cloud Function.
package p

import (
	"cloud.google.com/go/datastore"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// CommentDelete function deletes comment
func CommentDelete(w http.ResponseWriter, r *http.Request) {
	// 1. Write ID from request into struct d
	var d struct {
		ID int64 `json:"commentId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		idStr := r.FormValue("commentId")
		if idStr == "" {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(w, "Error While Parsing Request Body!\n URL: "+r.URL.Path)
			return
		}else {
			if n, err := strconv.Atoi(idStr); err == nil {
				d.ID = int64(n)
			} else {
				fmt.Println(err) /* log error */
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = fmt.Fprint(w, idStr, " is not an integer.")
				return
			}
		}
	}

	// 2. Connect to database
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, ProjectName)
	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(w, err)
		return
	}

	// 3. Delete comment
	commentKey := datastore.IDKey(EntityName, d.ID, nil)
	err = client.Delete(ctx, commentKey)
	if err != nil {
		fmt.Println(err) /* log error */
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Comment was successfully deleted")

}
