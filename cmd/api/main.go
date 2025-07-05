package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gui-laranjeira/livreria/internal/core/router"
)

func main() {
	r := gin.Default()
	router.SetupRoutes(r)
	if err := r.Run(); err != nil {
		panic(err)
	}
}
