package service

import (
	"github.com/Ahmad940/assetly_server/app/model"
	"github.com/Ahmad940/assetly_server/platform/db"
	"gorm.io/gorm/clause"
)

func GetAUser(id string) (model.User, error) {
	users := model.User{}

	err := db.DB.Where("id = ?", id).First(&users).Error
	if err != nil {
		return model.User{}, err
	}

	return users, nil
}

func GetAllUsers() ([]model.User, error) {
	users := []model.User{}

	err := db.DB.Find(&users).Error
	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}

// UpdateUser update user password
func UpdateUser(param model.UpdateUser) (model.User, error) {
	user := model.User{
		ID: param.ID,
	}

	err := db.DB.Model(&user).Clauses(clause.Returning{}).Updates(param).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func UpdateUserAdmin(param model.UpdateUserAdmin) (model.User, error) {
	user := model.User{
		ID:   param.ID,
		Role: model.UserRoleAdmin,
	}

	err := db.DB.Model(&user).Clauses(clause.Returning{}).Updates(param).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
