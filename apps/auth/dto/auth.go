package dto

import "time"

type SignInDto struct {
	Username  string  `json:"username" binding:"required"`
	Password  string  `json:"password" binding:"required"`
	UserAgent *string `json:"user_agent,omitempty"`
	IpAddress *string `json:"ip_address,omitempty"`
	Device    *string `json:"device,omitempty"`
}

type SignUpDto struct {
	SignInDto
	Name  string    `json:"name" binding:"required"`
	Email string    `json:"email" binding:"required"`
	Dob   time.Time `json:"dob" binding:"required"`
	Roles []string  `json:"role,omitempty"`
}
