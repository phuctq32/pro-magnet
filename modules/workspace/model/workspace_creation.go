package wsmodel

type WorkspaceCreation struct {
	Name   string `json:"name" validate:"required"`
	UserId string
}
