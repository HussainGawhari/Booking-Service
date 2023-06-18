package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Routings(route *gin.Engine, db *sql.DB) {

	// V1 group routes
	// v1 := route.Group("/v1")

	//Articles
	// article := v1.Group("/article").Use(middlewares.Auth())
	// {
	// 	article.GET("/", apiService.Articles)
	// 	article.GET("/:id", apiService.Article)
	// 	article.POST("/", apiService.CreateArticle)
	// 	article.POST("/upload", apiService.UploadToBucket)
	// 	article.PUT("/:id", apiService.ModifyArticle)
	// 	article.PUT("/publish/:id", apiService.ModifyBlogStatus)
	// }

	//User
	// user := v1.Group("/user")
	// user.POST("/", apiService.CreateUser)
	// user.POST("/login", apiService.UserLogin)

}
