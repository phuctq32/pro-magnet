package boarduc

import (
	"golang.org/x/net/context"
	"pro-magnet/components/asyncgroup"
	boardmodel "pro-magnet/modules/board/model"
	cardmodel "pro-magnet/modules/card/model"
	camodel "pro-magnet/modules/cardattachment/model"
	columnmodel "pro-magnet/modules/column/model"
	labelmodel "pro-magnet/modules/label/model"
	"time"
)

type CardRepo interface {
	FindByColumnId(ctx context.Context, cardStatus cardmodel.CardStatus, columnId string, cardIdsOrder []string) ([]cardmodel.Card, error)
}

type CardAttachmentRepo interface {
	CountByCardId(ctx context.Context, caStatus camodel.CardAttachmentStatus, cardId string) (int, error)
}

type LabelRepo interface {
	FindByBoardId(ctx context.Context, labelStatus labelmodel.LabelStatus, boardId string) ([]labelmodel.Label, error)
}

type ColumnRepo interface {
	FindByBoardId(ctx context.Context, status columnmodel.ColumnStatus, boardId string, columnIdsOrder []string) ([]columnmodel.Column, error)
}

type boardAggregator struct {
	asyncg               asyncgroup.AsyncGroup
	colRepo              ColumnRepo
	cardRepo             CardRepo
	caRepo               CardAttachmentRepo
	labelRepo            LabelRepo
	labelMap             map[string]labelmodel.Label
	getBoardLabelsDoneCh chan bool
}

func NewBoardAggregator(
	asyncg asyncgroup.AsyncGroup,
	colRepo ColumnRepo,
	cardRepo CardRepo,
	caRepo CardAttachmentRepo,
	labelRepo LabelRepo,
) *boardAggregator {
	return &boardAggregator{
		asyncg:               asyncg,
		colRepo:              colRepo,
		cardRepo:             cardRepo,
		caRepo:               caRepo,
		labelRepo:            labelRepo,
		labelMap:             make(map[string]labelmodel.Label),
		getBoardLabelsDoneCh: make(chan bool, 1),
	}
}

func (ba *boardAggregator) Aggregate(ctx context.Context, board *boardmodel.Board) error {
	return ba.asyncg.ProcessWithTimeout(
		ctx,
		time.Second*1000,
		ba.aggregateLabels(board),
		ba.aggregateColumns(board),
	)
}

func (ba *boardAggregator) aggregateLabels(board *boardmodel.Board) func(context.Context) error {
	return func(ctx context.Context) error {
		labels, err := ba.labelRepo.FindByBoardId(ctx, labelmodel.Active, *board.Id)
		if err != nil {
			return err
		}
		board.Labels = labels
		for i := 0; i < len(labels); i++ {
			ba.labelMap[*labels[i].Id] = labels[i]
		}
		ba.getBoardLabelsDoneCh <- true

		return nil
	}
}

func (ba *boardAggregator) aggregateColumns(board *boardmodel.Board) func(context.Context) error {
	return func(ctx context.Context) error {
		cols, err := ba.colRepo.FindByBoardId(ctx, columnmodel.Active, *board.Id, board.OrderedColumnIds)
		if err != nil {
			return err
		}
		board.Columns = cols

		// wait for get labels task done
		<-ba.getBoardLabelsDoneCh
		aggColumnTasks := make([]func(context.Context) error, 0)
		for i := 0; i < len(board.Columns); i++ {
			aggColumnTasks = append(aggColumnTasks, ba.aggregateCards(&board.Columns[i]))
		}

		return ba.asyncg.Process(ctx, aggColumnTasks...)
	}
}

func (ba *boardAggregator) aggregateCards(col *columnmodel.Column) func(context.Context) error {
	return func(ctx context.Context) error {
		cards, err := ba.cardRepo.FindByColumnId(ctx, cardmodel.Active, *col.Id, col.OrderedCardIds)
		if err != nil {
			return err
		}
		col.Cards = cards

		aggCardTasks := make([]func(context.Context) error, 0)
		for i := 0; i < len(col.Cards); i++ {
			aggCardTasks = append(aggCardTasks, ba.aggregateCard(&col.Cards[i]))
		}

		if err = ba.asyncg.Process(ctx, aggCardTasks...); err != nil {
			return err
		}

		return nil
	}
}

func (ba *boardAggregator) aggregateCard(card *cardmodel.Card) func(context.Context) error {
	return func(ctx context.Context) error {
		//memberCount := len(card.MemberIds)
		//card.MemberCount = &memberCount
		//
		//commentCount := len(card.Comments)
		//card.CommentCount = &commentCount

		attachmentCount, err := ba.caRepo.CountByCardId(ctx, camodel.Active, *card.Id)
		if err != nil {
			return nil
		}
		card.AttachmentCount = &attachmentCount

		card.Labels = make([]labelmodel.Label, 0)
		for i := 0; i < len(card.LabelIds); i++ {
			if _, ok := ba.labelMap[card.LabelIds[i]]; ok {
				card.Labels = append(card.Labels, ba.labelMap[card.LabelIds[i]])
			}
		}

		return nil
	}
}
