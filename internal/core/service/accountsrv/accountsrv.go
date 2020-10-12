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

	if isValidID(id) {
		return &domains.Account{}, fmt.Errorf("error")
	}

	obj, err := s.accountRepository.Get(id)
	if err != nil {
		return &domains.Account{}, fmt.Errorf("error")
	}

	return obj, nil
}

//Create Account
func (s *Service) Create(account *domains.Account) (*domains.Account, error) {
	return nil, nil
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
