package entities

import (
	"jobsPortal/helpers"

	"github.com/asaskevich/govalidator"

	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"not null" valid:"required"`
	Password string `gorm:"not null" form:"passwowrd" valid:"required,length(6|50)"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
