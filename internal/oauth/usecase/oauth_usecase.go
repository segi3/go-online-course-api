package oauth

import (
	"database/sql"
	"errors"
	"fmt"
	dto "online-course/internal/oauth/dto"
	entity "online-course/internal/oauth/entity"
	repository "online-course/internal/oauth/repository"
	useUseCase "online-course/internal/user/usecase"
	utils "online-course/pkg/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type OauthUseCase interface {
	Login(loginRequestBody dto.LoginRequestBody) (*dto.LoginResponse, error)
	Refresh(refreshTokenRequestBody dto.RefreshTokenRequestBody) (*dto.LoginResponse, error)
}

type OauthUseCaseImpl struct {
	oauthClientRepository       repository.OauthClientRepository
	oauthAccessTokenRepository  repository.OauthAccessTokenRepository
	oauthRefreshTokenRepository repository.OauthRefreshTokenRepository
	userUseCase                 useUseCase.UserUseCase
}

// Login implements OauthUseCase
func (usecase *OauthUseCaseImpl) Login(loginRequestBody dto.LoginRequestBody) (*dto.LoginResponse, error) {
	// check oauth client id and oauth client secret
	oauthClient, err := usecase.oauthClientRepository.FindByClientIdAndClientSecret(loginRequestBody.ClientId, loginRequestBody.ClientSecret)
	if err != nil {
		return nil, err
	}

	var user dto.UserResponse

	dataUser, err := usecase.userUseCase.FindByEmail(loginRequestBody.Email)
	fmt.Println(dataUser)
	if err != nil {
		return nil, errors.New("username or password is invalid")
	}

	user.ID = dataUser.ID
	user.Name = dataUser.Name
	user.Email = dataUser.Email
	user.Password = dataUser.Password

	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequestBody.Password))
	if err != nil {
		return nil, errors.New("incorrect password")
	}

	expirationTime := time.Now().Add(24 * 365 * time.Hour)

	claims := &dto.ClaimsResponse{
		ID:      user.ID,
		Name:    user.Email,
		Email:   user.Email,
		IsAdmin: false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return nil, err
	}

	// insert token to access token table
	oauthAccessTokenData := entity.OauthAccessToken{
		OauthClient:   &entity.OauthClient{},
		OauthClientID: &oauthClient.ID,
		UserID:        user.ID,
		Token:         tokenString,
		Scope:         "*",
		ExpiredAt: sql.NullTime{
			Time: expirationTime,
		},
	}

	oauthAccessToken, err := usecase.oauthAccessTokenRepository.Create(oauthAccessTokenData)
	if err != nil {
		return nil, err
	}

	// insert data to oauth refresh token table
	oauthRefreshTokenData := entity.OauthRefreshToken{
		OauthAccessTokenID: &oauthAccessToken.ID,
		UserID:             user.ID,
		Token:              utils.RandomString(128),
		ExpiredAt: sql.NullTime{
			Time: time.Now().Add(24 * 366 * time.Hour),
		},
	}

	oauthRefreshToken, err := usecase.oauthRefreshTokenRepository.Create(oauthRefreshTokenData)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  tokenString,
		RefreshToken: oauthRefreshToken.Token,
		Type:         "Bearer",
		ExpiredAt:    expirationTime.Format(time.RFC3339),
		Scope:        "*",
	}, nil
}

// Refresh implements OauthUseCase
func (*OauthUseCaseImpl) Refresh(refreshTokenRequestBody dto.RefreshTokenRequestBody) (*dto.LoginResponse, error) {
	panic("unimplemented")
}

func NewOauthUseCase(oauthClientRepository repository.OauthClientRepository,
	oauthAccessTokenRepository repository.OauthAccessTokenRepository,
	oauthRefreshTokenRepository repository.OauthRefreshTokenRepository,
	userUseCase useUseCase.UserUseCase) OauthUseCase {
	return &OauthUseCaseImpl{oauthClientRepository, oauthAccessTokenRepository, oauthRefreshTokenRepository, userUseCase}
}
