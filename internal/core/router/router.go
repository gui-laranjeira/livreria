package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gui-laranjeira/livreria/configs"
	"github.com/gui-laranjeira/livreria/internal/books"
	"github.com/gui-laranjeira/livreria/internal/core/infrastructure/database"
	"github.com/gui-laranjeira/livreria/internal/publisher"
)

func SetupRoutes(r *gin.Engine) {
	gin.SetMode(gin.DebugMode)
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Pong")
	})

	api := r.Group("/v1")
	injectDependencies(api)
}

func injectDependencies(r *gin.RouterGroup) {
	_, err := configs.LoadDBConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	db, err := database.OpenConnection()
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	publisherRepository := publisher.NewPublisherRepositoryAdapter(db)
	publisherService := publisher.NewPublisherServiceAdapter(publisherRepository)
	publisherHandler := publisher.NewPublisherHandlerAdapter(publisherService)

	bookRepository := books.NewBookRepositoryAdapter(db)
	bookService := books.NewBookServiceAdapter(bookRepository)
	bookHandler := books.NewBookHandlerAdapter(bookService, publisherService)

	r.GET("/book/:id", bookHandler.FindByID)
	r.POST("/book", bookHandler.Create)

	r.GET("/publisher/:id", publisherHandler.FindByID)
	r.GET("/publisher/name/:name", publisherHandler.FindByName)
	r.POST("/publisher", publisherHandler.Create)
}
