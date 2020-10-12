package ports

import (
	"hriqueXimenes/hexagonal/internal/core/domains"
)

//AccountService interface for account service
type AccountService interface {
	Get(string) (*domains.Account, error)
	Create(*domains.Account) (*domains.Account, error)
}
