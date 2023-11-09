package models

type File struct {
	Filename  string `json:"filename"`
	Extension string `json:"extension"`
	Url       string `json:"url"`
	Data      []byte `json:"-"`
	Folder    string `json:"-"`
}
