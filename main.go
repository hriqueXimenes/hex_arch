package main

import (
	"hriqueXimenes/hexagonal/db"
	"hriqueXimenes/hexagonal/internal/core/service/accountsrv"
	"hriqueXimenes/hexagonal/internal/handlers/accounthdl"
	"hriqueXimenes/hexagonal/internal/repositories/accountrepo"

	"github.com/gin-gonic/gin"
)

func main() {

	dbConfig := db.InitConfig()
	db, err := db.New(dbConfig)
	if err != nil {
		panic(err)
	}

	db.RunFileQuery("migrations/schema.sql")

	accountrepo := accountrepo.NewRelationalRepo(db)
	accountsrv := accountsrv.New(accountrepo)
	accounthdl := accounthdl.New(accountsrv)

	router := gin.New()
	router.GET("/account/:id", accounthdl.Get)
	router.POST("/account", accounthdl.Create)
	router.Run(":9000")
}
