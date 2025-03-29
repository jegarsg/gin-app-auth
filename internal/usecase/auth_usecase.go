package usecase

import (
	"GreatThanosApp/internal/repository"
	"GreatThanosApp/utils"
	"errors"
	"time"

	"github.com/google/uuid"
)

type AuthUseCase interface {
	Login(email, password string) (string, error)
	RefreshNewTokens(refreshToken string) (string, string, error)
	Logout(userId uuid.UUID) error
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

	accessToken, err := utils.GenerateJWT(user.UserID, user.Email, user.UserName) // Fix here
	if err != nil {
		return "", errors.New("failed to generate access token")
	}

	refreshToken := ""

	err = u.authRepo.SaveUserLogin(user.UserID, email, accessToken, refreshToken, time.Now().Add(7*24*time.Hour))
	if err != nil {
		return "", errors.New("failed to save user login session")
	}

	return accessToken, nil
}

func (u *authUseCase) RefreshNewTokens(refreshToken string) (string, string, error) {
	claims, err := utils.ValidateJWT(refreshToken)
	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	userId := (*claims)["userId"].(string)
	email, emailOk := (*claims)["email"].(string)
	username, usernameOk := (*claims)["username"].(string)
	if !emailOk || !usernameOk {
		return "", "", errors.New("invalid token payload")
	}
	parsedUUID, err := uuid.Parse(userId)
	if err != nil {
		return "", "", errors.New("error parsing uuid")
	}

	newAccessToken, err := utils.GenerateJWT(parsedUUID, email, username)
	if err != nil {
		return "", "", errors.New("failed to generate new access token")
	}

	newRefreshToken, err := utils.GenerateRefreshToken(email, username)
	if err != nil {
		return "", "", errors.New("failed to generate new refresh token")
	}

	err = u.authRepo.SaveUserLogin(parsedUUID, email, newAccessToken, newRefreshToken, time.Now().Add(7*24*time.Hour))
	if err != nil {
		return "", "", errors.New("failed to save user login session")
	}
	return newAccessToken, newRefreshToken, nil
}

func (u *authUseCase) Logout(userId uuid.UUID) error {
	return u.authRepo.InvalidateToken(userId)
}
