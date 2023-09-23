package repositories

import (
	"database/sql"

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
			INSERT INTO user_account (id, user_name, email, first_name, last_name, encrypted_password, created_at, updated_at, deleted_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
			`

	_, err := acc.Database.Query(
		query,
		usrAcc.Username,
		usrAcc.Email,
		usrAcc.Password,
		usrAcc.FullName,
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

	return usrAcc, err
}

// Delete implements UserAccountRepository.
func (*UserAccountRepositoryImp) Delete(usrAcc *models.UserAccount) (error error) {
	panic("unimplemented")
}

// Update implements UserAccountRepository.
func (*UserAccountRepositoryImp) Update(usrAcc *models.UserAccount) (error error) {
	panic("unimplemented")
}

func scanIntoUserAccount(rows *sql.Rows) (*models.UserAccount, error) {
	usrAcc := &models.UserAccount{}
	err := rows.Scan(
		&usrAcc.ID,
		&usrAcc.Username,
		&usrAcc.Email,
		&usrAcc.Password,
		&usrAcc.FullName,
		&usrAcc.CreatedAt,
		&usrAcc.UpdatedAt,
		&usrAcc.DeletedAt,
	)

	return usrAcc, err
}
