package accountsrv

import (
	"fmt"
	"hriqueXimenes/hexagonal/internal/core/domains"
	"hriqueXimenes/hexagonal/internal/core/ports"
	"strconv"
)

type Service struct {
	accountRepository ports.AccountRepository
}

//New Create new instance of service
func New(repository ports.AccountRepository) *Service {
	return &Service{
		accountRepository: repository,
	}
}

//Get Return a specific account by ID
func (s *Service) Get(id string) (*domains.Account, error) {

	if !isValidID(id) {
		return &domains.Account{}, fmt.Errorf("error")
	}

	return s.accountRepository.Get(id)
}

//Create Account
func (s *Service) Create(account *domains.Account) (*domains.Account, error) {

	if !isValidCreateParameters(account) {
		return &domains.Account{}, fmt.Errorf("error")
	}

	return s.accountRepository.Create(account)
}

func isValidID(id string) bool {
	num, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	if num <= 0 {
		return false
	}

	return true
}

func isValidCreateParameters(account *domains.Account) bool {
	if account == nil {
		return false
	}

	if account.Username == "" || account.Password == "" {
		return false
	}

	return true
}
