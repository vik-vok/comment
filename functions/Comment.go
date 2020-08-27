package p

import "time"

// Comment is good
type Comment struct {
	ID      int64  `datastore:"-"`
	VoiceID string `json:"voiceID"`
	UserID  string `json:"userID"`
	Text    string `json:"text"`
	Created         time.Time
}

// ProjectName is used for datastore.newClient()
const ProjectName string = "speech-similarity"

// EntityName is global constant which represents entity's (table) name in datastore
const EntityName string = "Comment"
