package api

import (
	"effective-test/internal/api/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run() {
	r := gin.Default()
	persons := r.Group("/person")
	{
		persons.POST("", handlers.CreatePerson)
		persons.DELETE("/:id", handlers.DeletePerson)
		persons.PUT("/:id", handlers.UpdatePerson)
		persons.GET("/:id", handlers.GetPerson)
		persons.GET("/page", handlers.GetPersonsWithPaging)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
