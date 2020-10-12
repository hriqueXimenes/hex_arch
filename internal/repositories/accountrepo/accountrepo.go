package accountrepo

import (
	"database/sql"

	"hriqueXimenes/hexagonal/internal/core/domains"
)

// AccountRepository contains objects for database communication.
type AccountRepository struct {
	*sql.DB
}

//NewRelationalRepo Create Postgres Repos
func NewRelationalRepo(db *sql.DB) *AccountRepository {
	return &AccountRepository{db}
}

//Get return a specific account by ID.
func (r *AccountRepository) Get(id string) (*domains.Account, error) {

	var account *domains.Account
	err := r.QueryRow("Select * FROM Accounts WHERE id = ?", id).Scan(&account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

//Create Account
func (r *AccountRepository) Create(account *domains.Account) (*domains.Account, error) {

	return nil, nil
}
