package p

// Comment is good
type Comment struct {
	ID      int64
	VoiceID string `json:"voiceID"`
	UserID  string `json:"userID"`
	Text    string `json:"text"`
}

// ProjectName is used for datastore.newClient()
const ProjectName string = "speech-similarity"

// EntityName is global constant which represents entity's (table) name in datastore
const EntityName string = "Comment"
