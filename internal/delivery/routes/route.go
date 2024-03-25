package routes

import (
	commenthandler "go-final-project/internal/delivery/handlers/comment"
	"go-final-project/internal/delivery/handlers/photo"
	socialmedia "go-final-project/internal/delivery/handlers/social_media"
	"go-final-project/internal/delivery/handlers/user"
	"go-final-project/internal/delivery/middlewares/jwt"
	commentService "go-final-project/internal/domain/comment/service"
	photoService "go-final-project/internal/domain/photo/service"
	socialMediaService "go-final-project/internal/domain/social_media/service"
	"go-final-project/internal/domain/user/service"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	// Initialize services
	userService := service.NewInstanceUserService()
	photoService := photoService.NewInstancePhotoService()
	commentService := commentService.NewCommentService()
	socialMediaService := socialMediaService.NewSocialMediaService()

	// Initialize handlers
	userHandler := user.NewUserHandler(userService)
	photoHandler := photo.NewPhotoHandler(photoService)
	commentHandler := commenthandler.NewCommentHandler(commentService)
	socialMediaHandler := socialmedia.NewSocialMediaHandler(socialMediaService)

	// API version 1 routes
	v1 := router.Group("/api/v1")
	{
		userV1 := v1.Group("/users")
		{
			userV1.POST("/login", userHandler.UserLogin)
			userV1.POST("/register", userHandler.UserRegister)
			userV1.POST("/delete", jwt.ValidateJWT(userHandler.DeleteUser))
		}

		photoV1 := v1.Group("/photos")
		{
			photoV1.POST("/", jwt.ValidateJWT(photoHandler.CreatePhoto))
			photoV1.GET("/", jwt.ValidateJWT(photoHandler.GetPhoto))
			photoV1.GET("/:id", jwt.ValidateJWT(photoHandler.GetPhotoByID))
			photoV1.PUT("/:id", jwt.ValidateJWT(photoHandler.UpdatePhoto))
			photoV1.DELETE("/:id", jwt.ValidateJWT(photoHandler.DeletePhoto))
		}

		commentV1 := v1.Group("/comments")
		{
			commentV1.POST("/", jwt.ValidateJWT(commentHandler.CreateComment))
			commentV1.GET("/", jwt.ValidateJWT(commentHandler.GetComments))
			commentV1.PUT("/:id", jwt.ValidateJWT(commentHandler.UpdateComment))
			commentV1.DELETE("/:id", jwt.ValidateJWT(commentHandler.DeleteComment))
		}

		socialmediaV1 := v1.Group("/socialmedias")
		{
			socialmediaV1.POST("/", jwt.ValidateJWT(socialMediaHandler.CreateSocialMedia))
			socialmediaV1.GET("/", jwt.ValidateJWT(socialMediaHandler.GetSocialMedias))
			socialmediaV1.PUT("/:id", jwt.ValidateJWT(socialMediaHandler.UpdateSocialMedia))
			socialmediaV1.DELETE("/:id", jwt.ValidateJWT(socialMediaHandler.DeleteSocialMedia))
		}
	}
}
