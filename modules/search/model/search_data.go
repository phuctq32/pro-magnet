package searchmodel

import (
	boardmodel "pro-magnet/modules/board/model"
	cardmodel "pro-magnet/modules/card/model"
	wsmodel "pro-magnet/modules/workspace/model"
)

type SearchData struct {
	Workspaces []wsmodel.Workspace `json:"workspaces"`
	Boards     []boardmodel.Board  `json:"boards"`
	Cards      []cardmodel.Card    `json:"cards"`
}
