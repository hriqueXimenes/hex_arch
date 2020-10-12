package cmd

import (
	"github.com/hriqueXimenes/hexagonal/internal/repositories/accountsrepo"
	"github.com/hriqueXimenes/hexagonal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	accountRepository := accountsrepo.New()
	accountService := service.New(gamesRepository, uidgen.New())
	gamesHandler := gamehdl.NewHTTPHandler(gamesService)

	router := gin.New()
	router.GET("/account/:id", gamesHandler.Get)
	router.POST("/account", gamesHandler.Create)

	router.Run(":8080")
}
