package p

import "time"

// Comment is good
type Comment struct {
	CommentId int64  `json:"commentId" datastore:"-" `
	VoiceId   string `json:"voiceId"`
	UserId    string `json:"userId"`
	Text      string `json:"text"`
	Created   time.Time
}

// ProjectName is used for datastore.newClient()
const ProjectName string = "speech-similarity"

// EntityName is global constant which represents entity's (table) name in datastore
const EntityName string = "Comment"
