package repositories

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/washington-shoji/gin-api/models"
)

type UserAccountRepositoryImp struct {
	Database *sql.DB
}

func NewUserAccountRePository(Db *sql.DB) UserAccountRepository {
	return &UserAccountRepositoryImp{
		Database: Db,
	}
}

// FindAll implements UserAccountRepository.
func (acc *UserAccountRepositoryImp) FindAll() (usrAcc []*models.UserAccount, error error) {
	query := `SELECT * FROM user_account WHERE deleted_at IS NULL`

	rows, err := acc.Database.Query(query)
	if err != nil {
		return nil, err
	}

	accs := []*models.UserAccount{}

	for rows.Next() {
		ac, err := scanIntoUserAccount(rows)
		if err != nil {
			return nil, err
		}
		accs = append(accs, ac)
	}

	return accs, nil
}

// Create implements UserAccountRepository.
func (acc *UserAccountRepositoryImp) Create(usrAcc *models.UserAccount) (error error) {
	query := `
			INSERT INTO user_account (id, user_name, email, full_name, encrypted_password, created_at)
			VALUES ($1, $2, $3, $4, $5, $6);
			`
	_, err := acc.Database.Query(
		query,
		usrAcc.ID,
		usrAcc.Username,
		usrAcc.Email,
		usrAcc.FullName,
		usrAcc.Password,
		usrAcc.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

// FindByID implements UserAccountRepository.
func (acc *UserAccountRepositoryImp) FindByID(id uuid.UUID) (usrAcc *models.UserAccount, error error) {

	query := `SELECT * FROM user_account WHERE id = $1 AND deleted_at IS NULL`

	rows, err := acc.Database.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		acc, err := scanIntoUserAccount(rows)
		if err != nil {
			return nil, err
		}

		usrAcc = acc
	}
	return usrAcc, fmt.Errorf("account with number [%s] not found", id)
}

// Update implements UserAccountRepository.
func (acc *UserAccountRepositoryImp) Update(usrAcc *models.UserAccount) (error error) {
	query := `
	UPDATE user_account
	SET 
	user_name = $2, 
	email = $3,
	encrypted_password = $4,
	full_name = $5,
	updated_at = $6
	WHERE id = $1
	`

	_, err := acc.Database.Query(query,
		usrAcc.ID,
		usrAcc.Username,
		usrAcc.Email,
		usrAcc.Password,
		usrAcc.FullName,
		usrAcc.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

// Delete implements UserAccountRepository.
func (acc *UserAccountRepositoryImp) Delete(usrAcc *models.UserAccount) (error error) {
	query := `
	UPDATE user_account
	SET deleted_at = $1 WHERE id = $2
	`

	_, err := acc.Database.Query(query, usrAcc.DeletedAt, usrAcc.ID)
	if err != nil {
		return err
	}
	return nil
}

// FindByUsername implements UserAccountRepository.
func (acc *UserAccountRepositoryImp) FindByUsername(username string) (usrAcc *models.UserAccount, error error) {

	query := `SELECT * FROM user_account WHERE user_name = $1 AND deleted_at IS NULL`

	rows, err := acc.Database.Query(query, username)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		acc, err := scanIntoUserAccount(rows)
		if err != nil {
			return nil, err
		}

		usrAcc = acc
	}

	if usrAcc == nil {
		return nil, fmt.Errorf("account with number [%s] not found", username)
	}
	return usrAcc, nil
}

func scanIntoUserAccount(rows *sql.Rows) (*models.UserAccount, error) {

	usrAcc := &models.UserAccount{}
	err := rows.Scan(
		&usrAcc.ID,
		&usrAcc.Username,
		&usrAcc.Email,
		&usrAcc.FullName,
		&usrAcc.Password,
		&usrAcc.CreatedAt,
		&usrAcc.UpdatedAt,
		&usrAcc.DeletedAt,
	)

	return usrAcc, err
}
