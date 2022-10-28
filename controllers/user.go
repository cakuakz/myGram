package controllers

import (
	"final-project/helpers"
	"final-project/models"
	"final-project/database"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"net/http"
	"strconv"
)

var (
	appJSON = "application/json"
)

// register function
func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	User := models.User{}

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&User); err != nil {
			if strings.Contains(err.Error(), "ParseInt") {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Bad Request",
					"message": err.Error(),
				})
				return
			} else {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
		}
	} else {
		if err := c.ShouldBind(&User); err != nil {
			if strings.Contains(err.Error(), "ParseInt") {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Bad Request",
					"message": err.Error(),
				})
				return
			} else {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
		}
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"message": err.Error(),
			})
			return
		}
		
	c.JSON(http.StatusCreated, gin.H{
		"age":      User.Age,
		"email":    User.Email,
		"id":       User.ID,
		"username": User.Username,
	})
}



// login function
func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	User := models.User{}

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&User); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := c.ShouldBind(&User); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	password := User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"message": "Invalid Email",
			})
			return
		}
	} else if !helpers.ComparePassword([]byte(User.Password), []byte(password)) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "Invalid Password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}


// update function
func UserUpdate(c *gin.Context) {
	db := database.GetDB()
	auth := c.MustGet("auth").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	User := models.User{}
	UserId := uint(auth["id"].(float64))

	userId, e := strconv.Atoi(c.Param("userId"))
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"msg":   "Invalid user id",
		})
		return
	}

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&User); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := c.ShouldBind(&User); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	User.ID = UserId

	if userId != int(UserId) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"msg":   "You can't update other user",
		})
		return
	}

	err := db.Debug().Model(&User).Where("id = ?", userId).Updates(&User).Select("age").Scan(&User.Age).Error

	if err != nil {
		if strings.Contains(err.Error(), "users email have to be unique") {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Conflict",
				"msg":   "Email already exists",
			})
			return
		} else if strings.Contains(err.Error(), "users username have to be unique") {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Conflict",
				"msg":   "Username already exists",
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"msg":   err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         User.ID,
		"email":      User.Email,
		"username":   User.Username,
		"age":        User.Age,
		"updated_at": User.UpdatedAt,
	})
}



// delete function
func UserDelete(c *gin.Context) {
	db := database.GetDB()
	auth := c.MustGet("auth").(jwt.MapClaims)
	User := models.User{}
	UserId := uint(auth["id"].(float64))

	err := db.Debug().Model(&User).Where("id = ?", UserId).Delete(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}