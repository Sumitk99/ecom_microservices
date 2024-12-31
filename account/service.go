package account

import (
	"context"
	"errors"
	"fmt"
	"github.com/Sumitk99/ecom_microservices/account/helper"
	"github.com/Sumitk99/ecom_microservices/account/models"
	"github.com/go-playground/validator/v10"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

type Service interface {
	SignUp(ctx context.Context, name, password, email, phone, userType string) (*Account, error)
	Login(ctx context.Context, email, phone, password string) (*Account, error)
	GetAccount(ctx context.Context) (*Account, error)
	GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
	Authentication(ctx context.Context) (*Account, error)
	AddAddress(ctx context.Context, address models.AddAddressRequest) (*models.Address, error)
	GetAddresses(ctx context.Context) ([]*models.Address, error)
	DeleteAddress(ctx context.Context, addressID string) error
	GetAddress(ctx context.Context, addressId string) (*models.Address, error)
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
	if err = s.repository.SignUp(ctx, acc); err != nil {
		log.Println(err)
		return nil, err
	}

	acc.Token, acc.RefreshToken, err = helper.GenerateTokens(acc.Name, acc.Email, acc.Phone, acc.UserType, acc.ID)
	if err != nil {
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

func (s *accountService) Authentication(ctx context.Context) (*Account, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("metadata not found")
		return nil, errors.New("no user metadata found in context")
	}

	token := md.Get("authorization")
	if token == nil {
		return nil, errors.New(fmt.Sprintf("No authorization token not found in context"))
	}
	clientToken := token[0]
	claims, err := helper.ValidateToken(clientToken)

	if err != nil || claims == nil {
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

func (s *accountService) GetAccount(ctx context.Context) (*Account, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("metadata not found")
		return nil, errors.New("no user metadata found in context")
	}

	id := md.Get("UserID")
	if id == nil {
		return nil, errors.New(fmt.Sprintf("UserID not found in context"))
	}
	acc, err := s.repository.GetAccountByID(ctx, id[0])
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

func (s *accountService) AddAddress(ctx context.Context, add models.AddAddressRequest) (*models.Address, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("metadata not found")
		return nil, errors.New("no user metadata found in context")
	}

	id := md.Get("UserID")
	if id == nil || len(id) == 0 {
		return nil, errors.New(fmt.Sprintf("UserID not found in context"))
	}
	address := models.Address{
		Name:          add.Name,
		Phone:         add.Phone,
		Street:        add.Street,
		City:          add.City,
		State:         add.State,
		ZipCode:       add.ZipCode,
		Country:       add.Country,
		IsDefault:     add.IsDefault,
		ApartmentUnit: add.ApartmentUnit,
		UserID:        id[0],
		AddressID:     ksuid.New().String(),
		CreatedAt:     time.Now().String(),
	}
	fmt.Println("Address Side: ", address)
	err := s.repository.AddAddress(ctx, &address)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &address, nil
}

func (s *accountService) GetAddresses(ctx context.Context) ([]*models.Address, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("metadata not found")
		return nil, errors.New("no user metadata found in context")
	}

	id := md.Get("UserID")
	if id == nil || len(id) == 0 {
		return nil, errors.New(fmt.Sprintf("UserID not found in context"))
	}
	accountID := id[0]
	addresses, err := s.repository.GetAddresses(ctx, accountID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return addresses, nil
}

func (s *accountService) DeleteAddress(ctx context.Context, addressID string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("metadata not found")
		return errors.New("no user metadata found in context")
	}

	id := md.Get("UserID")
	if id == nil || len(id) == 0 {
		return errors.New(fmt.Sprintf("UserID not found in context"))
	}
	accountID := id[0]

	err := s.repository.DeleteAddress(ctx, addressID, accountID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *accountService) GetAddress(ctx context.Context, addressId string) (*models.Address, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("metadata not found")
		return nil, errors.New("no user metadata found in context")
	}
	AccountId := md.Get("UserID")
	if AccountId == nil || len(AccountId) == 0 {
		return nil, errors.New(fmt.Sprintf("UserID not found in context"))
	}

	address, err := s.repository.GetAddress(ctx, addressId, AccountId[0])
	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("Error getting address: %s", err))
	}
	return &models.Address{
		Name:          address.Name,
		Phone:         address.Phone,
		Street:        address.Street,
		City:          address.City,
		State:         address.State,
		ZipCode:       address.ZipCode,
		Country:       address.Country,
		IsDefault:     address.IsDefault,
		ApartmentUnit: address.ApartmentUnit,
		AddressID:     address.AddressID,
		UserID:        address.UserID,
		CreatedAt:     address.CreatedAt,
	}, nil
}
