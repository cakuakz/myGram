package controllers

import (
	"final-project/database"
	"final-project/models"
	"final-project/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func CommentCreate(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)


	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	Comment := models.Comment{}

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&Comment); err != nil {
			if strings.Contains(err.Error(), "ParseInt") {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Bad Request",
					"message": err.Error(),
				})
				return
			}
		}
	} else {
		if err := c.ShouldBind(&Comment); err != nil {
			if strings.Contains(err.Error(), "ParseInt") {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Bad Request",
					"message": err.Error(),
				})
				return
			}
		}
	}

	Comment.UserId = userId

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoId,
		"user_id":    Comment.UserId,
		"created_at": Comment.CreatedAt,
	})
}

func CommentGet(c *gin.Context) {
	db := database.GetDB()

	var (
		Comments []models.Comment
		Data     []interface{}
	)

	err := db.Debug().Preload("User").Preload("Photo").Find(&Comments).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message":   err.Error(),
		})
		return
	}

	for i := range Comments {
		User := gin.H{
			"id":       Comments[i].User.ID,
			"email":    Comments[i].User.Email,
			"username": Comments[i].User.Username,
		}

		Photo := gin.H{
			"id":        Comments[i].Photo.ID,
			"title":     Comments[i].Photo.Title,
			"caption":   Comments[i].Photo.Caption,
			"photo_url": Comments[i].Photo.PhotoUrl,
			"user_id":   Comments[i].Photo.UserId,
		}

		Data = append(Data, gin.H{
			"id":         Comments[i].ID,
			"message":    Comments[i].Message,
			"photo_id":   Comments[i].PhotoId,
			"user_id":    Comments[i].UserId,
			"updated_at": Comments[i].UpdatedAt,
			"created_at": Comments[i].CreatedAt,
			"User":       User,
			"Photo":      Photo,
		})
	}

	c.JSON(http.StatusOK, Data)
}

func CommentUpdate(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	Comment := models.Comment{}

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&Comment); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := c.ShouldBind(&Comment); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	Comment.ID = uint(commentId)
	Comment.UserId = userId

	err := db.Debug().Model(&Comment).Preload("User").Preload("Photo").Where("id = ?", commentId).Updates(models.Comment{Message: Comment.Message}).Select("photo_id").Scan(&Comment.PhotoId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoId,
		"user_id":    Comment.UserId,
		"updated_at": Comment.UpdatedAt,
	})
}

func CommentDelete(c *gin.Context) {
	db := database.GetDB()


	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	Comment := models.Comment{}

	err := db.Debug().Where("id = ? AND user_id = ?", commentId, userId).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your comment has been successfully deleted",
	})
}