package boardmodel

const (
	BoardCollectionName string = "boards"
)

type BoardStatus uint8

const (
	Deleted BoardStatus = 0
	Active  BoardStatus = 1
)
