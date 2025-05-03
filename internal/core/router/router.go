package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gui-laranjeira/livreria/configs"
	"github.com/gui-laranjeira/livreria/internal/books/handler"
	"github.com/gui-laranjeira/livreria/internal/books/repository"
	"github.com/gui-laranjeira/livreria/internal/books/service"
	"github.com/gui-laranjeira/livreria/internal/core/infrastructure/database"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Pong")
	})

	api := r.Group("/v1")
	injectDependencies(api)
}

func injectDependencies(r *gin.RouterGroup) {
	dbConfigs, err := configs.LoadDBConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}
	db, err := database.OpenConnection(dbConfigs)
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	bookRepository := repository.NewBookRepositoryAdapter(db)
	bookService := service.NewBookServiceAdapter(bookRepository)
	bookHandler := handler.NewBookHandlerAdapter(bookService)

	r.GET("/books/:id", bookHandler.FindByID)
}
