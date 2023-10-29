package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/Ahmad940/assetly_server/app/model"
	"github.com/Ahmad940/assetly_server/pkg/util"
	"github.com/Ahmad940/assetly_server/platform/db"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

var invalidCred = "invalid email or password"

// Login
func Login(param model.Login) (model.AuthResponse, error) {
	var user model.User

	err := db.DB.Where("country_code = ? and phone_number = ?", param.CountryCode, param.PhoneNumber).First(&user).Error
	if err != nil {
		// if user not found
		if SqlErrorNotFound(err) {

			return model.AuthResponse{}, fmt.Errorf("phone number not registered")
		} else {
			fmt.Println("Error fetching credentials, reason:", err)
			return model.AuthResponse{}, err
		}
	}

	if passwordMatched := util.CompareHashedPassword(param.PassCode, user.PassCode); !passwordMatched {
		return model.AuthResponse{}, errors.New(invalidCred)
	}

	// generateToken
	token, err := util.GenerateToken(user)
	if err != nil {
		log.Println("Error occurred while generating token:", err)
		return model.AuthResponse{}, err
	}

	return model.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

// CreateAccount
func CreateAccount(param model.CreateUser) (model.AuthResponse, error) {
	var user model.User

	err := db.DB.Where("country_code = ? and phone_number = ?", param.CountryCode, param.PhoneNumber).First(&user).Error
	if SqlErrorIgnoreNotFound(err) != nil {
		return model.AuthResponse{}, err
	}

	// checking if user is registered or not
	if (user != model.User{}) {
		return model.AuthResponse{}, errors.New("phone number in use")
	}

	user_id := gonanoid.Must()
	txError := db.DB.Transaction(func(tx *gorm.DB) error {
		account_no, err := util.GenerateRandomNumber(10)
		if err != nil {
			log.Println("Error generating account number, reason:", err)
			return err
		}
		err = tx.Model(&user).Create(&model.User{
			ID:          user_id,
			FirstName:   param.FirstName,
			LastName:    param.LastName,
			CountryCode: param.CountryCode,
			PhoneNumber: param.PhoneNumber,
			Email:       param.Email,
			Wallet: &model.Wallet{
				ID:       gonanoid.Must(),
				WalletNo: account_no,
			},
			UserDetail: &model.UserDetail{
				ID:      gonanoid.Must(),
				Country: param.Country,
				State:   param.State,
				LGA:     param.LGA,
				Area:    param.Area,
				DOB:     param.DOB,
				Gender:  param.Gender,
			},
		}).Error
		if err != nil {
			log.Println("Error While creating user, reason:", err)
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	if txError != nil {
		log.Println("Transaction failed, reason:", err)
		return model.AuthResponse{}, err
	}

	err = db.DB.Preload("UserDetail").Preload("Wallet").First(&user, "id = ?", user_id).Error
	if err != nil {
		log.Println("Error While retrieving user account, reason:", err)
		return model.AuthResponse{}, err
	}

	// generateToken
	token, err := util.GenerateToken(user)
	if err != nil {
		log.Println("Error occurred while generating token:", err)
		return model.AuthResponse{}, err
	}

	return model.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}
