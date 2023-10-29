package service

import (
	"errors"
	"log"

	"github.com/Ahmad940/assetly_server/app/model"
	"github.com/Ahmad940/assetly_server/platform/db"
)

func GetWalletHistory(wallet_no string) ([]model.TransactionHistory, error) {
	var history []model.TransactionHistory = []model.TransactionHistory{}

	err := db.DB.Find(&history, "wallet_no = ?", wallet_no).Error
	if err != nil {
		if SqlErrorNotFound(err) {
			return []model.TransactionHistory{}, errors.New("wallet not found")
		}
		log.Println("Error retrieving history:", err)
		return []model.TransactionHistory{}, err
	}

	return history, nil
}
