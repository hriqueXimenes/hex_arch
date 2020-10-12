package ports

import (
	"hriqueXimenes/hexagonal/internal/core/domains"
)

//AccountRepository interface for account repository
type AccountRepository interface {
	Get(string) (*domains.Account, error)
	Create(*domains.Account) (*domains.Account, error)
}
