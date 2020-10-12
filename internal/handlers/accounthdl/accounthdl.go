package accounthdl

import (
	"hriqueXimenes/hexagonal/internal/core/domains"
	"hriqueXimenes/hexagonal/internal/core/ports"

	"github.com/gin-gonic/gin"
)

//Handler instance.
type Handler struct {
	accountService ports.AccountService
}

//New Create new instance of service
func New(service ports.AccountService) *Handler {
	return &Handler{
		accountService: service,
	}
}

//Get Return a specific account by ID
func (h *Handler) Get(c *gin.Context) {
	account, err := h.accountService.Get(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, account)
}

//Create account
func (h *Handler) Create(c *gin.Context) {

	req := &domains.Account{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	account, err := h.accountService.Create(req)
	if account != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, account)
}
