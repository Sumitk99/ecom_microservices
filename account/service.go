package account

import (
	"context"
	"errors"
	"github.com/Sumitk99/ecom_microservices/account/helper"
	"github.com/go-playground/validator/v10"
	"github.com/segmentio/ksuid"
	"log"
	"time"
)

type Service interface {
	SignUp(ctx context.Context, name, password, email, phone, userType string) (*Account, error)
	Login(ctx context.Context, email, phone, password string) (*Account, error)
	GetAccount(ctx context.Context, id string) (*Account, error)
	GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
	Authentication(ctx context.Context, token string) (*Account, error)
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

type accountService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &accountService{r}
}

func (s *accountService) SignUp(ctx context.Context, name, password, email, phone, userType string) (*Account, error) {
	count, err := s.repository.ValidateNewAccount(ctx, email, phone)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("email or phone already exists")
	}

	acc := Account{
		Name:     name,
		Email:    email,
		Phone:    phone,
		UserType: userType,
		Password: password,
	}

	validationErr := validator.New().Struct(acc)
	if validationErr != nil {
		log.Println(validationErr)
		return nil, validationErr
	}
	acc.ID = ksuid.New().String()

	acc.Password, err = helper.HashPassword(acc.Password)
	if err != nil {
		return nil, err
	}

	acc.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	acc.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	acc.Token, acc.RefreshToken, err = helper.GenerateTokens(acc.Name, acc.Email, acc.Phone, acc.UserType, acc.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := s.repository.SignUp(ctx, acc); err != nil {
		log.Println(err)
		return nil, err
	}
	acc.Password = ""
	return &acc, nil
}

func (s *accountService) Login(ctx context.Context, email, phone, providedpassword string) (*Account, error) {
	if email == "" && phone == "" {
		return nil, errors.New("Either email or phone must be provided!")
	}

	acc, err := s.repository.GetAccountByCredentials(ctx, email, phone)
	if err != nil {
		return nil, err
	}
	hashedPassword := acc.Password
	check, msg := helper.VerifyPassword(hashedPassword, providedpassword)
	if check != true {
		return nil, errors.New(msg)
	}
	acc.Token, acc.RefreshToken, err = helper.GenerateTokens(acc.Name, acc.Email, acc.Phone, acc.UserType, acc.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	acc.Password = ""
	return acc, nil
}

func (s *accountService) Authentication(ctx context.Context, token string) (*Account, error) {
	if token == "" {
		return nil, errors.New("Token is required")
	}

	claims, err := helper.ValidateToken(token)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &Account{
		ID:       claims.UserID,
		Name:     claims.Name,
		Email:    claims.Email,
		Phone:    claims.Phone,
		UserType: claims.UserType,
	}, nil
}

func (s *accountService) GetAccount(ctx context.Context, id string) (*Account, error) {
	acc, err := s.repository.GetAccountByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (s *accountService) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}
	accs, err := s.repository.ListAccounts(ctx, skip, take)
	if err != nil {
		return nil, err
	}
	return accs, nil
}
