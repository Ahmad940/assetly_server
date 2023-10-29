package service

import (
	"errors"
	"log"

	"github.com/Ahmad940/assetly_server/app/model"
	"github.com/Ahmad940/assetly_server/pkg/util"
	"github.com/Ahmad940/assetly_server/platform/db"
	"gorm.io/gorm/clause"
)

func GetAUser(id string) (model.User, error) {
	user := model.User{}

	err := db.DB.Preload("UserDetail").Preload("Wallet").Where("id = ?", id).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func GetAPlainUser(id string) (model.User, error) {
	user := model.User{}

	err := db.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func GetAllUsers() ([]model.User, error) {
	user := []model.User{}

	err := db.DB.Find(&user).Error
	if err != nil {
		return []model.User{}, err
	}

	return user, nil
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

func UpdateUserPassCode(userID string, param model.UpdateUserPassCode) error {
	user, err := GetAPlainUser(userID)
	if err != nil {
		log.Println("Error retrieving user:", err)
		return err
	}

	if user.PassCode != "" && param.CurrentPassCode == "" {
		return errors.New("current pass code field required")
	}

	if param.CurrentPassCode == param.NewPassCode {
		return errors.New("current and new pass code are same")
	}

	if user.PassCode != "" {
		if passwordMatched := util.CompareHashedPassword(param.CurrentPassCode, user.PassCode); !passwordMatched {
			return errors.New("incorrect current pass code")
		}
	}

	hashedPassCode, err := util.HashPassword(param.NewPassCode)
	if err != nil {
		log.Println("Error hashing new pass code:", err)
		return err
	}

	err = db.DB.Model(&user).Clauses(clause.Returning{}).Update("pass_code", hashedPassCode).Error
	if err != nil {
		log.Println("Error updating user:", err)
		return err
	}

	return nil
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
