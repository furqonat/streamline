package dto

import "time"

type UpdateUserDto struct {
	Name      *string    `json:"name,omitempty"`
	AvatarURL *string    `json:"avatar_url,omitempty"`
	Username  *string    `json:"username,omitempty"`
	Dob       *time.Time `json:"dob,omitempty"`
	Password  *string    `json:"password,omitempty"`
	Email     *string    `json:"email,omitempty"`
}
