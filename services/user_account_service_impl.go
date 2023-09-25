package services

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/models"
	"github.com/washington-shoji/gin-api/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserAccountServiceImp struct {
	UserAccountRepository repositories.UserAccountRepository
	Validate              *validator.Validate
}

func NewUserAccountService(usrAccRep repositories.UserAccountRepository, validate *validator.Validate) UserAccountService {
	return &UserAccountServiceImp{
		UserAccountRepository: usrAccRep,
		Validate:              validate,
	}
}

// Create implements UserAccountService.
func (acc *UserAccountServiceImp) Create(usrAcc *models.CreateUserAccountRequest) (error error) {
	err := acc.Validate.Struct(usrAcc)
	if err != nil {
		return err
	}

	id := uuid.New()
	time := time.Now().UTC()
	encryptPw, err := bcrypt.GenerateFromPassword([]byte(usrAcc.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	usrAccModel := models.UserAccount{
		ID:        id,
		Username:  usrAcc.Username,
		Email:     usrAcc.Email,
		Password:  string(encryptPw),
		FullName:  usrAcc.FullName,
		CreatedAt: time,
	}

	if err := acc.UserAccountRepository.Create(&usrAccModel); err != nil {
		return err
	}
	return nil
}

// Delete implements UserAccountService.
func (acc *UserAccountServiceImp) Delete(id uuid.UUID) (error error) {
	if _, err := acc.UserAccountRepository.FindByID(id); err != nil {
		return err
	}

	time := time.Now().UTC()
	usrAccModel := models.UserAccount{
		ID:        id,
		DeletedAt: &time,
	}

	if err := acc.UserAccountRepository.Delete(&usrAccModel); err != nil {
		return err
	}

	return nil
}

// FindAll implements UserAccountService.
func (acc *UserAccountServiceImp) FindAll() (usrAcc []*models.UserAccountResponse, error error) {
	result, err := acc.UserAccountRepository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := usrAcc
	for _, rst_item := range result {
		resp = append(resp, &models.UserAccountResponse{
			ID:       rst_item.ID,
			Username: rst_item.Username,
			Email:    rst_item.Email,
			FullName: rst_item.FullName,
		})
	}

	return resp, nil
}

// FindByID implements UserAccountService.
func (acc *UserAccountServiceImp) FindByID(id uuid.UUID) (usrAcc *models.UserAccountResponse, error error) {

	result, err := acc.UserAccountRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	resp := &models.UserAccountResponse{
		ID:       result.ID,
		Username: result.Username,
		Email:    result.Email,
		FullName: result.FullName,
	}

	return resp, nil
}

// Update implements UserAccountService.
func (acc *UserAccountServiceImp) Update(id uuid.UUID, usrAcc *models.UpdateUserAccountRequest) (error error) {
	err := acc.Validate.Struct(usrAcc)
	if err != nil {
		return err
	}

	time := time.Now().UTC()
	encryptPw, err := bcrypt.GenerateFromPassword([]byte(usrAcc.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	usrAccModel := models.UserAccount{
		ID:        id,
		Username:  usrAcc.Username,
		Email:     usrAcc.Email,
		Password:  string(encryptPw),
		FullName:  usrAcc.FullName,
		UpdatedAt: &time,
	}

	if err := acc.UserAccountRepository.Update(&usrAccModel); err != nil {
		return err
	}
	return nil
}
