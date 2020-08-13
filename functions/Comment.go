package p

// Comment is good
type Comment struct {
	ID      int64
	VoiceID int64  `json:"voiceID"`
	UserID  int64  `json:"userID"`
	Text    string `json:"text"`
}

// ProjectName is used for datastore.newClient()
const ProjectName string = "speech-similarity"

// EntityName is global constant which represents entity's (table) name in datastore
const EntityName string = "Comment"
