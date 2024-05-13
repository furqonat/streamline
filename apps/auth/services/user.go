package services

import (
	"apps/auth/db"
	"apps/auth/dto"
	"apps/auth/utils"
	"context"
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
	return u.db.User.FindUnique(db.User.ID.Equals(id)).Exec(context.Background())
}

func (u UserService) UpdateUser(id string, data dto.UpdateProfileDto) (*db.UserModel, error) {
	var password *string
	if data.Password != nil {
		psw := u.hash.Hash(*data.Password)
		password = &psw
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
