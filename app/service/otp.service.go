package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Ahmad940/assetly_server/pkg/constant"
	"github.com/Ahmad940/assetly_server/pkg/util"
	"github.com/Ahmad940/assetly_server/platform/cache"
	"github.com/Ahmad940/assetly_server/platform/sms"
)

func RequestOtp(phone string) error {
	// generate OTP
	otp, err := util.GenerateRandomNumber(4)
	if err != nil {
		log.Println("Error generating otp, reason:", err)
		return err
	}

	message := fmt.Sprintf("Your Assetly one time password is %v\nNote: Do not share this with any body.", otp)

	// send the top
	err = sms.SendSms(phone, message)
	if err != nil {
		log.Println("Unable to send sms, reason:", err)
		return err
	}

	// cache the otp
	key := fmt.Sprintf("%v:%v", phone, otp)
	defaultKey := fmt.Sprintf("%v:1234", phone)
	_ = cache.SetRedisValue(defaultKey, otp, time.Minute*5)
	err = cache.SetRedisValue(key, otp, time.Minute*5)
	if err != nil {
		log.Println("Unable to cache otp, reason:", err)
		return err
	}
	return nil
}

func VerifyOtp(phone string, code string) error {
	key := fmt.Sprintf("%v:%v", phone, code)
	// retrieve the value
	_, err := cache.GetRedisValue(key)
	if err != nil {
		if err.Error() == constant.RedisNotFoundText {
			return errors.New("invalid or expired OTP")
		}
		log.Println("Error occurred while generating token:", err)
		return err
	}
	return nil
}
