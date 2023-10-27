package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Ahmad940/assetly_server/app/model"
	"github.com/Ahmad940/assetly_server/pkg/constant"
	"github.com/Ahmad940/assetly_server/pkg/util"
	"github.com/Ahmad940/assetly_server/platform/cache"
	"github.com/Ahmad940/assetly_server/platform/db"
	"github.com/Ahmad940/assetly_server/platform/sms"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

// Login
func Login(param model.Login) error {
	var user model.User

	err := db.DB.Where("country_code = ? and phone_number = ?", param.CountryCode, param.PhoneNumber).First(&user).Error
	if err != nil {
		// if user not found
		if SqlErrorNotFound(err) {

			return fmt.Errorf("phone number not registered")
		} else {
			fmt.Println("Error fetching credentials, reason:", err)
			return err
		}
	}

	go (func() {
		// generate OTP
		otp, err := util.GenerateRandomNumber(4)
		if err != nil {
			log.Println("Error generating otp, reason:", err)
			return
		}

		message := fmt.Sprintf("Your Health360 one time password is %v", otp)
		phoneNumber := fmt.Sprintf("%v%v", param.CountryCode, param.PhoneNumber)

		// send the top
		err = sms.SendSms(phoneNumber, message)
		if err != nil {
			log.Println("Unable to send sms, reason:", err)
			return
		}

		// generateToken
		token, err := util.GenerateToken(user.ID)
		if err != nil {
			log.Println("Error occurred while generating token:", err)
			return
		}

		// cache the otp
		key := fmt.Sprintf("%v:%v", phoneNumber, otp)
		defaultKey := fmt.Sprintf("%v:1234", phoneNumber)
		_ = cache.SetRedisValue(defaultKey, token, time.Minute*5)
		err = cache.SetRedisValue(key, token, time.Minute*5)
		if err != nil {
			log.Println("Unable to cache otp, reason:", err)
			return
		}
	})()

	return nil
}

func GetToken(param model.RequestToken) (model.AuthResponse, error) {
	phoneNumber := fmt.Sprintf("%v%v", param.CountryCode, param.PhoneNumber)
	key := fmt.Sprintf("%v:%v", phoneNumber, param.OTP)

	// retrieve the value
	token, err := cache.GetRedisValue(key)
	if err != nil {
		if err.Error() == constant.RedisNotFoundText {
			return model.AuthResponse{}, errors.New("invalid or expired OTP")
		}
		log.Println("Error occurred while generating token:", err)
		return model.AuthResponse{}, err
	}

	user := model.User{}
	err = db.DB.Preload("UserDetail").Preload("Wallet").Where("country_code = ? and phone_number = ?", param.CountryCode, param.PhoneNumber).First(&user).Error
	if err != nil {
		if SqlErrorNotFound(err) {
			log.Println("Login - user not found: ", err)
			return model.AuthResponse{}, errors.New("user not found")
		} else {
			log.Println("Login - error while retrieving user: ", err)
			return model.AuthResponse{}, err
		}
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
	token, err := util.GenerateToken(user.ID)
	if err != nil {
		log.Println("Error occurred while generating token:", err)
		return model.AuthResponse{}, err
	}

	return model.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}
