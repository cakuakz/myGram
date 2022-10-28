package routers

import (
	"final-project/controllers"
	"final-project/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/:userId", middleware.Authentication(), controllers.UserUpdate)
		userRouter.DELETE("/", middleware.Authentication(), controllers.UserDelete)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/photo", controllers.PhotoCreate)
		photoRouter.GET("/photo", controllers.PhotoGet)
		photoRouter.PUT("/:photoId", middleware.PhotoAuthorization(), controllers.PhotoUpdate)
		photoRouter.DELETE("/:photoId", middleware.PhotoAuthorization(), controllers.PhotoDelete)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/", controllers.CommentCreate)
		commentRouter.GET("/", controllers.CommentGet)
		commentRouter.PUT("/:commentId", middleware.CommentAuthorization(), controllers.CommentUpdate)
		commentRouter.DELETE("/:commentId", middleware.CommentAuthorization(), controllers.CommentDelete)
	}

	socialMediaRouter := router.Group("/socialmedias")
	{
		socialMediaRouter.Use(middleware.Authentication())
		socialMediaRouter.POST("/", controllers.SocialMediaCreate)
		socialMediaRouter.GET("/", controllers.SocialMediaGet)
		socialMediaRouter.PUT("/:socialMediaId", middleware.SocialMediaAuthorization(), controllers.SocialMediaUpdate)
		socialMediaRouter.DELETE("/:socialMediaId", middleware.SocialMediaAuthorization(), controllers.SocialMediaDelete)
	}
	return router
}