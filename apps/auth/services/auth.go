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
	a.logger.Info(len(data.Password))
	if len(data.Password) < utils.MINIMUN_PASSWORD_LENGTH {
		return nil, errors.New("password must be at least 6 characters")
	}

	user, err := a.createUser(data)

	if err != nil {
		return nil, err
	}

	return a.cretaeAuthToken(user, &data.SignInDto)
}

func (a AuthService) findOrCreateRole(name string) (*db.RoleModel, error) {
	u, err := a.db.Role.FindFirst(
		db.Role.Name.Equals(name),
	).Exec(context.Background())

	if err != nil {
		return a.db.Role.CreateOne(
			db.Role.Name.Set(name),
		).Exec(context.Background())
	}
	return u, nil
}

func (a AuthService) createUser(data dto.SignUpDto) (*db.UserModel, error) {
	hashedPassword := a.hash.Hash(data.Password)
	role := []db.RoleModel{}
	if len(data.Roles) > 0 {
		for _, name := range data.Roles {
			r, _ := a.findOrCreateRole(name)
			role = append(role, *r)
		}
	} else {
		r, _ := a.findOrCreateRole(utils.ROLE_USER)
		role = append(role, *r)
	}

	user, err := a.db.User.CreateOne(
		db.User.Name.Set(data.Name),
		db.User.Username.Set(data.Username),
		db.User.Email.Set(data.Email),
		db.User.Password.Set(hashedPassword),
		db.User.Dob.Set(data.Dob),
	).With(
		db.User.Roles.Fetch(),
	).Exec(context.Background())
	if err != nil {
		return nil, err
	}

	for _, r := range role {
		_, err := a.db.User.FindUnique(
			db.User.ID.Equals(user.ID),
		).Update(
			db.User.Roles.Link(
				db.Role.ID.Equals(r.ID),
			),
		).Exec(context.Background())
		if err != nil {
			return nil, err
		}
	}
	return user, nil
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
		db.User.Email.Equals(usernameOrEmail),
	).With(
		db.User.Roles.Fetch(),
	).Exec(context.Background())

	userWithUsername, errWithUsername := a.db.User.FindUnique(
		db.User.Username.Equals(usernameOrEmail),
	).With(
		db.User.Roles.Fetch(),
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
