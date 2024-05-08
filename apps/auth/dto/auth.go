package dto

type SignInDto struct {
	Username  string  `json:"username" binding:"required"`
	Password  string  `json:"password" binding:"required"`
	UserAgent *string `json:"user_agent,omitempty"`
	IpAddress *string `json:"ip_address,omitempty"`
	Device    *string `json:"device,omitempty"`
}

type SignUpDto struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
