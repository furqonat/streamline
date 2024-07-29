package services

import (
	"apps/auth/db"
	"apps/auth/dto"
	"apps/auth/utils"
	"context"
	"errors"
)

type UserService struct {
	logger utils.Logger
	db     utils.Database
	hash   utils.PwHash
}

func NewUserService(logger utils.Logger, db utils.Database, hash utils.PwHash) UserService {
	return UserService{
		logger: logger,
		db:     db,
		hash:   hash,
	}
}

func (u UserService) GetUser(id string) (*db.UserModel, error) {
	return u.db.User.FindUnique(db.User.ID.Equals(id)).Select(
		db.User.ID.Field(),
		db.User.Name.Field(),
		db.User.Username.Field(),
		db.User.AvatarURL.Field(),
	).Exec(context.Background())
}

func (u UserService) UpdateUser(id string, data dto.UpdateUserDto) (*db.UserModel, error) {
	var password *string
	if data.Password != nil {
		psw := u.hash.Hash(*data.Password)
		password = &psw
	}

	if data.Username != nil {
		if err := u.checkUsernameExists(*data.Username); err != nil {
			return nil, err
		}
	}

	if data.Email != nil {
		if err := u.checkEmailExists(*data.Email); err != nil {
			return nil, err
		}
	}

	return u.db.User.FindUnique(db.User.ID.Equals(id)).Update(
		db.User.Name.SetIfPresent(data.Name),
		db.User.AvatarURL.SetIfPresent(data.AvatarURL),
		db.User.Username.SetIfPresent(data.Username),
		db.User.Dob.SetIfPresent(data.Dob),
		db.User.Password.SetIfPresent(password),
		db.User.Email.SetIfPresent(data.Email),
	).Exec(context.Background())
}

func (u UserService) checkUsernameExists(username string) error {
	existsUsername, _ := u.db.User.FindUnique(
		db.User.Username.Equals(username),
	).Exec(context.Background())

	if existsUsername != nil {
		return errors.New("username already exists")
	}
	return nil
}

func (u UserService) checkEmailExists(email string) error {
	existsEmail, _ := u.db.User.FindUnique(
		db.User.Email.Equals(email),
	).Exec(context.Background())

	if existsEmail != nil {
		return errors.New("email already exists")
	}
	return nil
}
