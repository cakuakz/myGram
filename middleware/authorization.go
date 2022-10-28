package middleware

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"final-project/database"
	"final-project/models"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		photoId, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"message": "Invalid photo",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := int(userData["id"].(float64))
		Photo := models.Photo{}

		err = db.Debug().Select("user_id").First(&Photo, photoId).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Not Found",
				"messsage": "Photo doesn't exist",
			})
			return
		}

		if Photo.UserId != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"meesage": "You are not authorized to access this photo",
			})
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		commentId, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"message": "Invalid comment id",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := int(userData["id"].(float64))
		Comment := models.Comment{}

		err = db.Debug().Select("user_id").First(&Comment, commentId).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Not Found",
				"meesage": "Comment doesn't exist",
			})
			return
		}

		if Comment.UserId != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"message": "You are not authorized to access this comment",
			})
			return
		}

		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"message": "Invalid social media id",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := int(userData["id"].(float64))
		SocialMedia := models.SocialMedia{}

		err = db.Debug().Select("user_id").First(&SocialMedia, socialMediaId).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Not Found",
				"message": "Social media doesn't exist",
			})
			return
		}

		if SocialMedia.UserId != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"message": "You are not authorized to access this social media",
			})
			return
		}

		c.Next()
	}
}