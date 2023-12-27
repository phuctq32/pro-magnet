package wsmembermodel

type WorkspaceMember struct {
	Id          *string `json:"_id,omitempty" bson:"_id,omitempty"`
	WorkspaceId string  `json:"workspaceId" bson:"workspaceId"`
	UserId      string  `json:"userId" bson:"userId"`
}
