package boardmodel

type BoardCreation struct {
	WorkspaceId string `json:"workspaceId" validate:"required"`
	Name        string `json:"name" validate:"required"`
	UserId      string
}
