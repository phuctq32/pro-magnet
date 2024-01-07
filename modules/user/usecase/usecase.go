package useruc

import (
	"golang.org/x/net/context"
	"pro-magnet/components/hasher"
	cardmodel "pro-magnet/modules/card/model"
	usermodel "pro-magnet/modules/user/model"
	wsmodel "pro-magnet/modules/workspace/model"
)

type UserRepository interface {
	FindById(ctx context.Context, id string) (*usermodel.User, error)
	UpdatePasswordById(ctx context.Context, id, password string) error
	UpdateById(ctx context.Context, id string, updateData *usermodel.UserUpdate) (*usermodel.User, error)
	FindSimpleUsersByIds(ctx context.Context, userIds []string) ([]usermodel.User, error)
	SearchUsersByEmail(ctx context.Context, emailSearchStr string, exceptedUserIds []string) ([]usermodel.User, error)
	UpdateSkills(ctx context.Context, userId string, skills []string) error
}

type CardRepository interface {
	FindById(ctx context.Context, id string) (*cardmodel.Card, error)
}

type BoardMemberRepository interface {
	IsBoardMember(ctx context.Context, boardId, userId string) (bool, error)
	FindMemberIdsByBoardId(ctx context.Context, boardId string) ([]string, error)
}

type WorkspaceRepository interface {
	FindById(ctx context.Context, id string) (*wsmodel.Workspace, error)
}

type WorkspaceMemberRepository interface {
	FindMemberIdsByWorkspaceId(ctx context.Context, workspaceId string) ([]string, error)
}

type userUseCase struct {
	userRepo     UserRepository
	cardRepo     CardRepository
	bmRepo       BoardMemberRepository
	wsRepo       WorkspaceRepository
	wsMemberRepo WorkspaceMemberRepository
	hasher       hasher.Hasher
}

func NewUserUseCase(
	userRepo UserRepository,
	cardRepo CardRepository,
	bmRepo BoardMemberRepository,
	wsRepo WorkspaceRepository,
	wsMemberRepo WorkspaceMemberRepository,
	hasher hasher.Hasher,
) *userUseCase {
	return &userUseCase{
		userRepo:     userRepo,
		cardRepo:     cardRepo,
		bmRepo:       bmRepo,
		wsRepo:       wsRepo,
		wsMemberRepo: wsMemberRepo,
		hasher:       hasher,
	}
}
