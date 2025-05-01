package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gui-laranjeira/livreria/configs"
	"github.com/gui-laranjeira/livreria/internal/core/infrastructure/database"
	"github.com/gui-laranjeira/livreria/internal/core/router"
)

func main() {
	r := gin.Default()
	router.SetupRoutes(r)

	dbConfig, err := configs.LoadDBConfig()
	if err != nil {
		panic(err)
	}

	_, err = database.OpenConnection(dbConfig)
	if err != nil {
		panic(err)
	}

	if err := r.Run(); err != nil {
		panic(err)
	}
}
