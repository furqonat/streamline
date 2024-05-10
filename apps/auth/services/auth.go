package services

import (
	"apps/auth/db"
	"apps/auth/dto"
	"apps/auth/utils"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	jwt    utils.Jwt
	logger utils.Logger
	db     utils.Database
	hash   utils.PwHash
}

func NewAuthService(jwt utils.Jwt, logger utils.Logger, db utils.Database, hash utils.PwHash) AuthService {
	return AuthService{
		jwt,
		logger,
		db,
		hash,
	}
}

func (a AuthService) SignIn(data dto.SignInDto) (*string, error) {
	return a.signInWithEmailOrUsername(data)
}

func (a AuthService) SignUp(data dto.SignUpDto) (*string, error) {

	username, _ := a.getUserFromDatabase(data.Username)
	if username != nil {
		return nil, errors.New("user with username already exists")
	}

	email, _ := a.getUserFromDatabase(data.Email)
	if email != nil {
		return nil, errors.New("user with email already exists")
	}

	user, err := a.createUser(data)

	if err != nil {
		return nil, err
	}

	return a.cretaeAuthToken(user, &data.SignInDto)
}

func (a AuthService) createUser(data dto.SignUpDto) (*db.UserModel, error) {
	return a.db.User.CreateOne(
		db.User.Name.Set(data.Name),
		db.User.Username.Set(data.Username),
		db.User.Email.Set(data.Email),
		db.User.Password.Set(data.Password),
		db.User.Dob.Set(data.Dob),
	).Exec(context.Background())
}

func (a AuthService) signInWithEmailOrUsername(data dto.SignInDto) (*string, error) {
	user, err := a.getUserFromDatabase(data.Username)
	if err != nil {
		return nil, err
	}
	if !a.hash.Compare(data.Password, user.Password) {
		return nil, errors.New("invalid password")
	}
	return a.cretaeAuthToken(user, &data)
}

func (a AuthService) getUserFromDatabase(usernameOrEmail string) (*db.UserModel, error) {
	userWithEmail, errWithEmail := a.db.User.FindUnique(
		db.User.Username.Equals(usernameOrEmail),
	).Exec(context.Background())

	userWithUsername, errWithUsername := a.db.User.FindUnique(
		db.User.Username.Equals(usernameOrEmail),
	).Exec(context.Background())

	if errWithEmail == nil && errWithUsername == nil {
		return nil, errors.New("multiple users found")
	}

	if errWithEmail != nil && errWithUsername != nil {
		return nil, errWithEmail
	}

	if errWithEmail == nil {
		return userWithEmail, nil
	}

	return userWithUsername, nil
}

func (a AuthService) cretaeAuthToken(user *db.UserModel, data *dto.SignInDto) (*string, error) {

	refreshToken, errRefreshToken := a.jwt.Encode(&utils.Claims{
		ID:   user.ID,
		Name: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})
	if errRefreshToken != nil {
		return nil, errRefreshToken
	}
	var roles []utils.Role
	for _, role := range user.Roles() {
		roles = append(roles, utils.Role{ID: role.ID, Name: role.Name})
	}
	token, err := a.jwt.Encode(&utils.Claims{
		ID:   user.ID,
		Name: user.Name,
		Role: roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	})
	if err != nil {
		return nil, err
	}

	if err := a.saveAuth(user, token, refreshToken, data); err != nil {
		return nil, err
	}
	return &token, nil
}

func (a AuthService) saveAuth(user *db.UserModel, token, refreshToken string, data *dto.SignInDto) error {
	_, err := a.db.Auth.CreateOne(
		db.Auth.User.Link(
			db.User.ID.Equals(user.ID),
		),
		db.Auth.Token.Set(token),
		db.Auth.RefreshToken.Set(refreshToken),
		db.Auth.IPAddress.SetIfPresent(data.IpAddress),
		db.Auth.UserAgent.SetIfPresent(data.UserAgent),
		db.Auth.Device.SetIfPresent(data.Device),
	).Exec(context.Background())

	if err != nil {
		return err
	}

	return nil
}
