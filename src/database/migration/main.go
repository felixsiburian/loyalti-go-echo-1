package main

import (
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"
)

func main() {
	//db := database.ConnectionDB()
	db := database.ConnectionDB()
	merchant := model.Merchant{
		Created:               time.Now(),
		CreatedBy:             "Admin",
		Modified:              time.Now(),
		ModifiedBy:            "Admin",
		Active:                true,
		IsDeleted:             false,
		Deleted:               nil,
		Deleted_by:            "",
		MerchantName:          "ABCD",
		MerchantEmail:         "abcd@gmail.com",
		MerchantPhoneNumber:   "1231231",
		MerchantProvince:      "DKI Jakarta",
		MerchantCity:          "Jakarta",
		MerchantAddress:       "Jalan Sudirman ",
		MerchantPostalCode:    "120212",
		MerchantCategoryId:    1,
		MerchantWebsite:       "www.abcd.com",
		MerchantMediaSocialId: 1,
		MerchantDescription:   "jual makanan enak",
		MerchantImageProfile:  "",
		MerchantGallery:       "",
	}
	outlet := model.Outlet{
		Created:          time.Now(),
		CreatedBy:        "Admin",
		Modified:         time.Now(),
		ModifiedBy:       "Admin",
		Active:           true,
		IsDeleted:        false,
		Deleted:          nil,
		Deleted_by:       "",
		OutletName:       "ABCD",
		OutletAddress:    "jalan jalan",
		OutletPhone:      "123123123",
		OutletCity:       "Jakarta",
		OutletProvince:   "DKI Jakarta",
		OutletPostalCode: "123123",
		OutletLongitude:  "",
		OutletLatitude:   "",
		OutletDay:        time.Time{},
		OutletHour:       time.Time{},
		MerchantEmail:    "abcd@gmail.com",
		Timezone:         "WIB",
	}

	db.Create(&outlet)
	db.Create(&merchant)
	//db.AutoMigrate(&model.City{}, &model.Employee{}, &model.MerchantCategory{}, &model.Merchant{}, &model.TransactionMerchant{},
	//		&model.Outlet{}, &model.Program{}, &model.SpecialProgram{}, &model.Province{}, &model.Reward{}, &model.MerchantSocialMedia{},
	//		&model.Voucher{})
}
