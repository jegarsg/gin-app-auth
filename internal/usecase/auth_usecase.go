package usecase

import (
	"GreatThanosApp/internal/repository"
	"GreatThanosApp/utils"
	"errors"
)

type AuthUseCase interface {
	Login(email, password string) (string, error)
	RefreshNewTokens(refreshToken string) (string, string, error)
}

type authUseCase struct {
	authRepo repository.AuthRepository
}

func NewAuthUseCase(authRepo repository.AuthRepository) AuthUseCase {
	return &authUseCase{authRepo: authRepo}
}

func (u *authUseCase) Login(email, password string) (string, error) {
	user, err := u.authRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	accessToken, err := utils.GenerateJWT(user.Email, user.UserName) // Fix here
	if err != nil {
		return "", errors.New("failed to generate access token")
	}

	return accessToken, nil
}

func (u *authUseCase) RefreshNewTokens(refreshToken string) (string, string, error) {
	claims, err := utils.ValidateJWT(refreshToken)
	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	email, emailOk := (*claims)["email"].(string)
	username, usernameOk := (*claims)["username"].(string)
	if !emailOk || !usernameOk {
		return "", "", errors.New("invalid token payload")
	}

	newAccessToken, err := utils.GenerateJWT(email, username)
	if err != nil {
		return "", "", errors.New("failed to generate new access token")
	}

	newRefreshToken, err := utils.GenerateRefreshToken(email, username)
	if err != nil {
		return "", "", errors.New("failed to generate new refresh token")
	}

	return newAccessToken, newRefreshToken, nil
}
