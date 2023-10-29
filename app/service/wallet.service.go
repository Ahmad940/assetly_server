package service

import (
	"errors"
	"log"

	"github.com/Ahmad940/assetly_server/app/model"
	"github.com/Ahmad940/assetly_server/platform/db"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

func GetWalletByNo(walletNo string) (model.Wallet, error) {
	var wallet model.Wallet
	err := db.DB.Preload("User", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("id", "first_name", "last_name")
	}).First(&wallet, "wallet_no = ?", walletNo).Select("wallet_no").Error
	if err != nil {
		log.Println("Error while fetching error:", err)
		return model.Wallet{}, err
	}

	return wallet, nil
}

func TopUpWallet(walletNo string, amount float64) (model.Wallet, error) {
	var wallet model.Wallet

	err := db.DB.First(&wallet, "wallet_no = ?", walletNo).Error
	if err != nil {
		if SqlErrorNotFound(err) {
			return model.Wallet{}, errors.New("wallet not found")
		}
		log.Println("Error while fetching wallet:", err)
		return model.Wallet{}, err
	}

	wallet.Balance += amount

	err = db.DB.Save(&wallet).Error
	if err != nil {
		log.Println("Error during toping up wallet:", err)
		return model.Wallet{}, err
	}

	err = db.DB.Create(&model.TransactionHistory{
		ID:                  gonanoid.Must(),
		WalletNo:            walletNo,
		Amount:              amount,
		Narration:           "Top up",
		TransactionType:     model.TransactionTypeCredit,
		TransactionStatus:   model.TransactionStatusSuccess,
		TransactionCategory: model.TransactionCategoryTopUp,
	}).Error
	if err != nil {
		log.Println("Error inserting history:", err)
	}

	return wallet, nil
}

func Transfer() (model.Wallet, error) {
	return model.Wallet{}, nil
}
