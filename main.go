package main

import (
	"hriqueXimenes/hexagonal/internal/core/service/accountsrv"
	"hriqueXimenes/hexagonal/internal/handlers/accounthdl"
	"hriqueXimenes/hexagonal/internal/repositories/accountrepo"

	"github.com/gin-gonic/gin"
)

func main() {
	accountrepo := accountrepo.NewRelationalRepo(nil)
	accountsrv := accountsrv.New(accountrepo)
	accounthdl := accounthdl.New(accountsrv)

	router := gin.New()
	router.GET("/account/:id", accounthdl.Get)
	router.POST("/account", accounthdl.Create)
	router.Run(":9000")
}
