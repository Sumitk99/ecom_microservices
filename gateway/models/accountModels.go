package models

type SignUpRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	UserType string `json:"user_type"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Name          string `json:"name"`
	ID            string `json:"id"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	UserType      string `json:"user_type"`
	JWT_Token     string `json:"jwt_token"`
	Refresh_Token string `json:"refresh_token"`
}
