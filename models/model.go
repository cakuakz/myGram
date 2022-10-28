package models

import (
	"time"
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	// "gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `json:"username" gorm:"unique;not null" form:"username" valid:"required~Username is required"`
	Email        string `json:"email" gorm:"uniqueIndex;not null" form:"email" valid:"email"`
	Password     string `json:"password" gorm:"not null" form:"password" valid:"required~Password is required,minstringlength(6)~Password must be at least 6 characters"`
	Age          uint   `json:"age" gorm:"not null" form:"age" valid:"required~Age is required,range(8|100)~No Permission for under 8 years old"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
		if u.Email == "" {
			err = errors.New("required email")
			return
		} else if u.Username == "" {
			err = errors.New("required username")
			return
		} else if u.Email == "" && u.Username == "" {
			err = errors.New("email and password is required")
			return
		} else if !govalidator.IsEmail(u.Email) {
			err = errors.New("email is not valid")
			return
		}
	
		return
	}


type Photo struct {
	gorm.Model
	Title    string `json:"title" gorm:"not null" form:"title" valid:"required~Title is required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" gorm:"not null" form:"photo_url" valid:"required~Photo URL is required"`
	UserId   int    `json:"user_id" form:"user_id"`
	User     *User  `json:"user"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		return errCreate
	}

	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	if p.Title == "" && p.PhotoUrl == "" {
		err = errors.New("title and url is required")
		return
	} else if p.Title == "" {
		err = errors.New("title is required")
		return
	} else if p.PhotoUrl == "" {
		err = errors.New("url is required")
		return
	}

	return
}

type Comment struct {
	gorm.Model
	Message string `json:"message" gorm:"not null" form:"message" valid:"required~Required A Message"`
	UserId  int    `json:"user_id" form:"user_id"`
	User    *User  `json:"user"`
	PhotoId int    `json:"photo_id" form:"photo_id"`
	Photo   *Photo `json:"photo"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, e := govalidator.ValidateStruct(c)

	if e != nil {
		return e
	}

	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, e := govalidator.ValidateStruct(c)

	if e != nil {
		return e
	}

	return
}

type SocialMedia struct {
	gorm.Model
	Name           string `json:"name" gorm:"not null" form:"name" valid:"required~Name is required"`
	SocialMediaUrl string `json:"social_media_url" gorm:"not null" form:"social_media_url" valid:"required~Social Media URL is required"`
	UserId         int    `json:"user_id" form:"user_id"`
	User           *User  `json:"user"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		return errCreate
	}

	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, e := govalidator.ValidateStruct(s)

	if e != nil {
		return e
	}

	return
}
