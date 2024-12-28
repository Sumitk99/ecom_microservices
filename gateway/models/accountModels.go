package models

import "time"

type SignUpRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"email,required"`
	Phone    string `json:"phone" validate:"required"`
	UserType string `json:"user_type" validate:"required,eq=ADMIN|eq=BUYER|eq=SELLER"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password" validate:"required"`
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

type Account struct {
	ID           string    `json:"id"`
	Name         string    `json:"name" validate:"required,min=2,max=100"`
	Password     string    `json:"password" validate:"required,min=6"`
	Email        string    `json:"email" validate:"email,required"`
	Phone        string    `json:"phone" validate:"required"`
	Token        string    `json:"token"`
	UserType     string    `json:"user_type" validate:"required,eq=ADMIN|eq=BUYER|eq=SELLER"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
