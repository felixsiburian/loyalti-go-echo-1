package main

import (
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

func main() {
	db := database.ConnectionDB()
	db.AutoMigrate(&model.City{}, &model.Employee{}, &model.MerchantCategory{}, &model.Merchant{}, &model.TransactionMerchant{},
			&model.Outlet{}, &model.Program{}, &model.SpecialProgram{}, &model.Province{}, &model.Reward{}, &model.MerchantSocialMedia{},
			&model.Voucher{})
}
