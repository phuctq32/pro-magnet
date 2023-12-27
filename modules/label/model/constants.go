package labelmodel

const (
	LabelCollectionName = "labels"
)

type LabelStatus uint8

const (
	Deleted LabelStatus = 0
	Active  LabelStatus = 1
)
