package services

import (
	"jobsPortal/entities"

	"gorm.io/gorm"
)

type JobsPortalService struct {
	DB *gorm.DB
}

func (js *JobsPortalService) CreateUser(user entities.User) (entities.User, error) {
	if err := js.DB.Debug().Create(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (js *JobsPortalService) Login(user entities.User) (entities.User, error) {
	if err := js.DB.Debug().Where("username = ?", user.Username).Take(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}
