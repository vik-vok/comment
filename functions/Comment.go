package p

import "cloud.google.com/go/datastore"

// Comment is good
type Comment struct {
	//ID      int64
	ID *datastore.Key `datastore:"__key__"`
	VoiceID string `json:"voiceID"`
	UserID  string `json:"userID"`
	Text    string `json:"text"`
}

// ProjectName is used for datastore.newClient()
const ProjectName string = "speech-similarity"

// EntityName is global constant which represents entity's (table) name in datastore
const EntityName string = "Comment"
