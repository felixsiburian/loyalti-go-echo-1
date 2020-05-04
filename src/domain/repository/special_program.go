package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

type SpecialProgramRepository interface {
	CreateSpecial(special *model.SpecialProgram) error
	UpdateSpecial(special *model.SpecialProgram) error
	DeleteSpecial(special *model.SpecialProgram) error
}

type special_repo struct {
	DB *gorm.DB
}

func CreateSpecialRepository(db *gorm.DB) SpecialProgramRepository {
	return &special_repo{
		DB: db,
	}
}

func (p *special_repo) CreateSpecial(special *model.SpecialProgram) error {
	db := database.ConnectionDB()
	specialObj := *special
	err := db.Create(&specialObj).Error
	return err
}

func (p *special_repo) UpdateSpecial(special *model.SpecialProgram) error {
	db := database.ConnectionDB()
	err := db.Model(&special).Where("program_name = ?", special.ProgramName).Update(&special).Error
	return err
}

func (p *special_repo) DeleteSpecial(special *model.SpecialProgram) error {
	db := database.ConnectionDB()
	err := db.Model(&special).Where("program_name = ?", special.ProgramName).Update("active", false).Error
	return err
}

//type SpecialRepository interface {
//	CreateSpecial (newspecial *model.SpecialProgram) error
//	UpdateSpecial (newspecial *model.SpecialProgram) error
//	DeleteSpecial (newspecial *model.SpecialProgram) error
//}
//
//type special_repo struct {
//	DB *gorm.DB
//}
//
//func CreateSpecialRepository (db *gorm.DB) SpecialRepository {
//	return  &special_repo{
//		DB:db,
//	}
//}
//
//func (p *special_repo) CreateSpecial (newspecial *model.SpecialProgram) error {
//	db:= database.ConnectionDB()
//	err := db.Create(&newspecial).Error
//	if err == nil {
//		fmt.Println("Tidak ada Error!")
//	}
//	return err
//}
//
//func (p *special_repo) UpdateSpecial (newspecial *model.SpecialProgram) error {
//	db := database.ConnectionDB()
//
//}

func CreateSpecial(special *model.SpecialProgram) string {
	db := database.ConnectionDB()
	specialObj := model.SpecialProgram{
		Created:               time.Now(),
		CreatedBy:             "Admin",
		Modified:              time.Now(),
		ModifiedBy:            "Admin",
		Active:                true,
		IsDeleted:             false,
		Deleted:               time.Time{},
		Deleted_by:            "",
		ProgramName:           special.ProgramName,
		ProgramImage:          special.ProgramImage,
		ProgramStartDate:      special.ProgramStartDate,
		ProgramEndDate:        special.ProgramEndDate,
		ProgramDescription:    special.ProgramDescription,
		Card:                  special.Card,
		OutletID:              special.OutletID,
		MerchantEmail:         special.MerchantEmail,
		CategoryId:            special.CategoryId,
		Benefit:               special.Benefit,
		TermsAndCondition:     special.TermsAndCondition,
		Tier:                  special.Tier,
		RedeemRules:           special.RedeemRules,
		RewardTarget:          special.RewardTarget,
		QRCodeId:              special.QRCodeId,
		IsReqBillNumber:       special.IsReqBillNumber,
		IsReqTotalTransaction: special.IsReqTotalTransaction,
		IsPushNotification:    special.IsPushNotification,
		IsLendCard:            special.IsLendCard,
		IsGiveCard:            special.IsGiveCard,
		IsWelcomeBonus:        special.IsWelcomeBonus,
		MerchantName:          "a",
		OutletName:            "b",
		CategoryName:          "c",
	}
	db.Create(&specialObj)
	defer db.Close()
	return specialObj.ProgramName
}

func UpdateSpecial(special *model.SpecialProgram) string {
	db := database.ConnectionDB()
	db.Model(&special).Updates(map[string]interface{}{
		"program_name":             special.ProgramName,
		"program_image":            special.ProgramImage,
		"program_start_date":       special.ProgramStartDate,
		"program_end_date":         special.ProgramEndDate,
		"program_description":      special.ProgramDescription,
		"card":                     special.Card,
		"outlet_id":                special.OutletID,
		"merchant_email":           special.MerchantEmail,
		"category_id":              special.CategoryId,
		"benefit":                  special.Benefit,
		"terms_and_condition":      special.TermsAndCondition,
		"tier":                     special.Tier,
		"redeem_rules":             special.RedeemRules,
		"reward_target":            special.RewardTarget,
		"qr_code_id":               special.QRCodeId,
		"is_req_bill_number":       special.IsReqBillNumber,
		"is_req_total_transaction": special.IsReqTotalTransaction,
		"is_push_notification":     special.IsPushNotification,
		"is_lend_card":             special.IsLendCard,
		"is_give_card":             special.IsGiveCard,
		"is_welcome_bonus":         special.IsWelcomeBonus,
	})
	special.Modified = time.Now()
	defer db.Close()
	return special.ProgramName
}

func DeleteSpecial(special *model.SpecialProgram) string {
	db := database.ConnectionDB()
	db.Model(&special).Where("id = ?", special.Id).Update("active", false)
	db.Model(&special).Where("id = ?", special.Id).Update("is_deleted", true)
	special.Deleted = time.Now()
	defer db.Close()
	return "berhasil dihapus"
}

func GetSpecialProgram(page *int, size *int, sort *int, category *int, email *string) []model.SpecialProgram {
	db := database.ConnectionDB()
	var program []model.SpecialProgram
	var rows *sql.Rows
	var err error
	var total int

	if page == nil && size == nil && sort == nil && category == nil && email == nil {
		fmt.Println("masuk if 1")
		rows, err = db.Find(&program).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil && sort != nil && category != nil && email != nil {
		fmt.Println("masuk if 2")
		switch *sort {
		case 1:
			rows, err = db.Find(&program).Where("category_id = ? AND merchant_email = ?", category, email).Order("merchant_name asc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		case 2:
			rows, err = db.Find(&program).Where("category_id = ? AND merchant_email = ?", category, email).Order("merchant_name desc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if page != nil && size != nil && sort != nil && category == nil && email != nil {
		fmt.Println("masuk if 3")
		switch *sort {
		case 1:
			rows, err = db.Find(&program).Where("merchant_email = ?", email).Order("merchant_name asc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		case 2:
			rows, err = db.Find(&program).Where("merchant_email = ?", email).Order("merchant_name desc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if page != nil && size != nil && sort != nil && category != nil && email == nil {
		fmt.Println("masuk if 4")
		switch *sort {
		case 1:
			rows, err = db.Find(&program).Where("category_id = ? ", category).Order("merchant_name asc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		case 2:
			rows, err = db.Find(&program).Where("category_id = ? ", category).Order("merchant_name desc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if page != nil && size != nil && sort != nil && category == nil && email == nil {
		fmt.Println("masuk if 5")
		switch *sort {
		case 1:
			rows, err = db.Find(&program).Order("merchant_name asc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		case 2:
			rows, err = db.Find(&program).Order("merchant_name desc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if page != nil && size != nil && sort == nil && category != nil && email == nil {
		fmt.Println("masuk if 6")
		rows, err = db.Find(&program).Where("category_id = ?", category).Order("merchant_name asc").Count(total).Limit(*size).Offset(*page).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil && sort == nil && category == nil && email != nil {
		fmt.Println("masuk if 7")
		rows, err = db.Find(&program).Where("merchant_email = ?", email).Order("merchant_name asc").Count(total).Limit(*size).Offset(*page).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil && sort == nil && category == nil && email == nil {
		fmt.Println("masuk if 8")
		rows, err = db.Find(&program).Order("merchant_name asc").Count(total).Limit(*size).Offset(*page).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page == nil && size == nil && sort == nil && category != nil && email == nil {
		fmt.Println("masuk if 9")
		rows, err = db.Find(&program).Where("category_id = ?", category).Order("merchant_name asc").Count(total).Limit(*size).Offset(*page).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page == nil && size == nil && sort == nil && category == nil && email != nil {
		fmt.Println("masuk if 9")
		rows, err = db.Find(&program).Where("merchant_email = ?", email).Order("merchant_name asc").Count(total).Limit(*size).Offset(*page).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	result := make([]model.SpecialProgram, 0)

	for rows.Next() {
		t := &model.SpecialProgram{}
		fmt.Println(t)

		err = rows.Scan(
			&t.Id,
			&t.Created,
			&t.CreatedBy,
			&t.Modified,
			&t.ModifiedBy,
			&t.Active,
			&t.IsDeleted,
			&t.Deleted,
			&t.Deleted_by,
			&t.ProgramName,
			&t.ProgramImage,
			&t.ProgramStartDate,
			&t.ProgramEndDate,
			&t.ProgramDescription,
			&t.Card,
			&t.OutletID,
			&t.MerchantEmail,
			&t.CategoryId,
			&t.Benefit,
			&t.TermsAndCondition,
			&t.Tier,
			&t.RedeemRules,
			&t.RewardTarget,
			&t.QRCodeId,
			&t.IsReqBillNumber,
			&t.IsReqTotalTransaction,
			&t.IsPushNotification,
			&t.IsLendCard,
			&t.IsGiveCard,
			&t.IsWelcomeBonus,
		)

		merchant := new(model.Merchant)
		db.Table("merchants").
			Select("merchants.merchant_name").
			Where("merchant_email = ?", t.MerchantEmail).
			First(&merchant)
		t.MerchantName = merchant.MerchantName

		category := new (model.MerchantCategory)

		db.Table("merchant_categories").
			Select("merchant_categories.category_name").
			Where("id = ? ", t.CategoryId).
			First(&category)
		t.CategoryName = category.CategoryName

		outlet := new (model.Outlet2)

		db.Table("outlets").
			Select("outlets.outlet_name").
			Where("id = ? ", t.OutletID).
			First(&outlet)
		t.OutletName = outlet.OutletName

		if err != nil {
			log.Fatal(err)
			return nil
		}
		result = append(result,*t)
	}
	db.Close()
	return result
}
