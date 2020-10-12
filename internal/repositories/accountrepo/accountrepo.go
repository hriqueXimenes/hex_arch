package accountrepo

import (
	"database/sql"
	"fmt"
	"hriqueXimenes/hexagonal/db"
	"hriqueXimenes/hexagonal/internal/core/domains"
)

// AccountRepository contains objects for database communication.
type AccountRepository struct {
	*db.Database
}

//NewRelationalRepo Create Postgres Repos
func NewRelationalRepo(db *db.Database) *AccountRepository {
	return &AccountRepository{db}
}

//Get return a specific account by ID.
func (r *AccountRepository) Get(id string) (*domains.Account, error) {

	account := &domains.Account{}
	query := fmt.Sprintf("Select * FROM Accounts WHERE id = %v", id)

	err := r.QueryRow(query).Scan(&account.ID, &account.Username, &account.Password)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return account, nil
}

//Create Account
func (r *AccountRepository) Create(account *domains.Account) (*domains.Account, error) {

	id := 0
	err := r.QueryRow("INSERT INTO Accounts (username,password) VALUES($1, $2) RETURNING id", account.Username, account.Password).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("error")
	}

	account.ID = id
	return account, nil
}
