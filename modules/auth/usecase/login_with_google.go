package authuc

import (
	"context"
	"pro-magnet/common"
	authmodel "pro-magnet/modules/auth/model"
	usermodel "pro-magnet/modules/user/model"
	"time"
)

func (uc *authUseCase) LoginWithGoogle(
	ctx context.Context, code string,
) (*authmodel.LoginResponse, error) {
	ggUser, err := uc.ggOauth.UserInfoFromCode(ctx, code)
	if err != nil {
		return nil, err
	}

	// If user existed:
	// - User is not google user, throw existed error, otherwise it's accepted.
	// If user not existed:
	// - Create a new Google user
	user, err := uc.userRepo.FindByEmail(ctx, ggUser.Email)
	if err != nil && err.Error() != usermodel.ErrUserNotFound.Error() {
		return nil, err
	}
	if user != nil {
		if *user.Type != usermodel.GoogleUser {
			return nil, common.NewBadRequestErr(authmodel.ErrUserExisted)
		}
	} else {
		// Create new Google user
		now := time.Now()
		isVerified := true
		user = &usermodel.User{
			CreatedAt:   &now,
			UpdatedAt:   &now,
			Email:       &ggUser.Email,
			Name:        ggUser.Name,
			Password:    nil,
			IsVerified:  &isVerified,
			Avatar:      ggUser.Avatar,
			PhoneNumber: ggUser.Phonenumber,
			Birthday:    ggUser.Birthday,
			Skills:      make([]string, 0),
		}
		userType := usermodel.GoogleUser
		user.Type = &userType

		userId, err := uc.userRepo.Create(ctx, user)
		if err != nil {
			return nil, err
		}

		user.Id = userId
	}

	accessToken, refreshToken, err := uc.generateTokenPair(ctx, user)
	if err != nil {
		return nil, err
	}

	return &authmodel.LoginResponse{
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
		User:         user,
	}, nil
}
