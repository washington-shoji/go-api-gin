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

// Delete implements UserAccountRepository.
func (*UserAccountRepositoryImp) Delete(usrAcc *models.UserAccount) (error error) {
	panic("unimplemented")
}

// FindAll implements UserAccountRepository.
func (*UserAccountRepositoryImp) FindAll() (usrAcc []*models.UserAccount, error error) {
	panic("unimplemented")
}

// FindByID implements UserAccountRepository.
func (*UserAccountRepositoryImp) FindByID(id uuid.UUID) (usrAcc *models.UserAccount, error error) {
	panic("unimplemented")
}

// Update implements UserAccountRepository.
func (*UserAccountRepositoryImp) Update(usrAcc *models.UserAccount) (error error) {
	panic("unimplemented")
}
