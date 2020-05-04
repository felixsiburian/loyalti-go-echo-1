package repository

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

type ProgramRepository interface {
	CreateProgram(program *model.Program) error
	UpdateProgram(program *model.Program) error
	DeleteProgram(program *model.Program) error
	GetProgram(page *int, size *int, sort *int, category *int, id *int) error
}

type repoProgram struct {
	DB *gorm.DB
}

func (p *repoProgram) CreateProgram(newprogram *model.Program) error {
	fmt.Println("masuk fungsi")
	program := model.Program{
		Created:            time.Now(),
		CreatedBy:          "Admin",
		Modified:           time.Now(),
		ModifiedBy:         "",
		Active:             true,
		IsDeleted:          false,
		Deleted:            time.Time{},
		Deleted_by:         "",
		ProgramName:        newprogram.ProgramName,
		ProgramImage:       newprogram.ProgramImage,
		ProgramStartDate:   newprogram.ProgramStartDate,
		ProgramEndDate:     newprogram.ProgramEndDate,
		ProgramDescription: newprogram.ProgramDescription,
		Card:               newprogram.Card,
		OutletID:           newprogram.OutletID,
		CategoryId:         newprogram.CategoryId,
		Benefit:            newprogram.Benefit,
		TermsAndCondition:  newprogram.TermsAndCondition,
		Tier:               newprogram.Tier,
		RedeemRules:        newprogram.RedeemRules,
		RewardTarget:       newprogram.RewardTarget,
		QRCodeId:           newprogram.QRCodeId,
		ProgramPoint:       newprogram.ProgramPoint,
		MinPayment:         newprogram.MinPayment,
	}

	db := database.ConnectionDB()
	err := db.Create(&program).Error
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
	return err
}

//func CreateRepositoryProgram(db *gorm.DB) ProgramRepository {
//	return &repoProgram{
//		DB:db,
//	}
//}

func CreateProgram(program *model.Program) string {
	db := database.ConnectionDB()
	programObj := model.Program{
		Created:               time.Now(),
		CreatedBy:             "Admin",
		Modified:              time.Now(),
		ModifiedBy:            "Admin",
		Active:                true,
		IsDeleted:             false,
		Deleted:               time.Time{},
		Deleted_by:            "",
		ProgramName:           program.ProgramName,
		ProgramImage:          program.ProgramImage,
		ProgramStartDate:      program.ProgramStartDate,
		ProgramEndDate:        program.ProgramEndDate,
		ProgramDescription:    program.ProgramDescription,
		Card:                  program.Card,
		OutletID:              program.OutletID,
		MerchantEmail:         program.MerchantEmail,
		CategoryId:            program.CategoryId,
		Benefit:               program.Benefit,
		TermsAndCondition:     program.TermsAndCondition,
		Tier:                  program.Tier,
		RedeemRules:           program.RedeemRules,
		RewardTarget:          program.RewardTarget,
		QRCodeId:              program.QRCodeId,
		ProgramPoint:          program.ProgramPoint,
		MinPayment:            program.MinPayment,
		IsReqBillNumber:       program.IsReqBillNumber,
		IsReqTotalTransaction: program.IsReqTotalTransaction,
		IsPushNotification:    program.IsPushNotification,
		IsLendCard:            program.IsLendCard,
		IsGiveCard:            program.IsGiveCard,
		IsWelcomeBonus:        program.IsWelcomeBonus,
		MerchantName:          "a",
		OutletName:            "b",
		CategoryName:          "c",
	}
	db.Create(&programObj)
	defer db.Close()
	return programObj.ProgramName
}

//func (p *repoProgram) UpdateProgram  (program *model.Program) error {
//	db := database.ConnectionDB()
//	err := db.Model(&program).Where("merchant_id = ?", program.MerchantId).Update(&program).Error
//	if err != nil {
//		fmt.Println("Error : ", err.Error())
//	}
//	defer db.Close()
//	return err
//}

func UpdateProgram(program *model.Program) string {
	db := database.ConnectionDB()
	db.Model(&program).Updates(map[string]interface{}{
		"program_name":             program.ProgramName,
		"program_image":            program.ProgramImage,
		"program_start_date":       program.ProgramStartDate,
		"program_end_date":         program.ProgramEndDate,
		"program_description":      program.ProgramDescription,
		"card":                     program.Card,
		"outlet_id":                program.OutletID,
		"merchant_email":           program.MerchantEmail,
		"category_id":              program.CategoryId,
		"benefit":                  program.Benefit,
		"terms_and_condition":      program.TermsAndCondition,
		"tier":                     program.Tier,
		"redeem_rules":             program.RedeemRules,
		"reward_target":            program.RewardTarget,
		"qr_code_id":               program.QRCodeId,
		"program_point":            program.ProgramPoint,
		"min_payment":              program.MinPayment,
		"is_req_bill_number":       program.IsReqBillNumber,
		"is_req_total_transaction": program.IsReqTotalTransaction,
		"is_push_notification":     program.IsPushNotification,
		"is_lend_card":             program.IsLendCard,
		"is_give_card":             program.IsGiveCard,
		"is_welcome_bonus":         program.IsWelcomeBonus,
	})
	program.Modified = time.Now()
	defer db.Close()
	return "Berhasil diUpdate"
}

func (p *repoProgram) DeleteProgram(program *model.Program) error {
	db := database.ConnectionDB()
	err := db.Model(&program).Where("id = ?", program.Id).Update("active", false).Error
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
	return err
}

func DeleteProgram(program *model.Program) string {
	db := database.ConnectionDB()
	db.Model(&program).Where("id= ?", program.Id).Update("active", false)
	db.Model(&program).Where("id= ?", program.Id).Update("is_deleted", true)
	program.Deleted = time.Now()
	defer db.Close()
	return "berhasil dihapus"
}

func TotalPoint(id int, pay int, pin string, outletid int, cardtype string) []model.TotalPoint {
	db := database.ConnectionDB()
	employee := model.Employee{}
	totalpoint := []model.TotalPoint{}
	program := model.Program{}
	card := model.ProgramCard{}
	db.Model(&program).Where("id = ?", id).Find(&program)
	db.Model(&employee).Where("employee_pin = ?", pin).Find(&employee)
	db.Model(&program).Where("outlet_id = ?", outletid).Find(&program)
	db.Model(&card).Where("card_type = ?", cardtype).Find(&card)

	if cardtype != card.CardType {
		fmt.Println("Customer tidak mendapatkan point dengan kartu ini")
		return nil
	}

	if outletid != program.OutletID {
		fmt.Println("Anda tidak berada di outlet yang tepat")
		return nil
	}

	if pin != employee.EmployeePin {
		fmt.Println("Pin Anda Salah. Silahkan Coba Lagi ")
		return nil
	}
	if outletid == employee.OutletId && pin == employee.EmployeePin {
		if pay < (program.MinPayment) {
			fmt.Println("Customer tidak mendapatkan poin ")
			return nil
		}
		var total = pay * (program.ProgramPoint) / (program.MinPayment)
		t := &model.TotalPoint{}
		t.Total = total
		updatepoint := append(totalpoint, *t)
		fmt.Printf("Customer mendapatkan %d poin \n", total)
		return updatepoint
	}
	defer db.Close()
	return nil
}

func TotalChop(id int, pay int, pin string, outletid int, cardtype string) []model.TotalChop {
	db := database.ConnectionDB()
	employee := model.Employee{}
	totalchop := []model.TotalChop{}
	program := model.Program{}
	db.Model(&program).Where("id = ?", id).Find(&program)
	db.Model(&employee).Where("employee_pin = ?", pin).Find(&employee)
	db.Model(&program).Where("outlet_id = ?", outletid).Find(&program)
	db.Model(&program).Where("card = ?", cardtype).Find(&program)

	if cardtype != program.Card {
		fmt.Println("Customer tidak mendapatkan chop dengan kartu ini")
		return nil
	}

	if outletid != program.OutletID {
		fmt.Println("Anda tidak memiliki akses ")
		return nil
	}
	if pin != employee.EmployeePin {
		fmt.Println("Pin Anda Salah. Silahkan Coba Lagi ")
		return nil
	}
	if outletid == employee.OutletId && pin == employee.EmployeePin {
		if pay < (program.MinPayment) {
			fmt.Println("Customer tidak mendapatkan tambahan chop ")
			return nil
		}

		if pay >= (program.MinPayment) {
			var total = pay / pay * 1
			t := &model.TotalChop{}
			t.Total = total
			updatechop := append(totalchop, *t)
			fmt.Printf("Customer mendapatkan %d chop \n", total)
			return updatechop
		}
	}
	defer db.Close()
	return nil
}

//func (p *repoProgram) GetProgram(page *int, size *int, sort *int, category *int, id *int) error{
//	db := database.ConnectionDB()
//	var program []model.Program
//	var rows *sql.Rows
//	var err error
//	var total int
//
//	if sort != nil {
//		switch *sort {
//		case 1:
//			if page != nil && size != nil && category == nil{
//				rows, err = db.Find(&program).Order("created asc").Count(total).Limit(*size).Offset(*page).Rows()
//				if err != nil {
//					panic(err)
//				}
//			}
//			if category != nil && page != nil && size != nil{
//				rows, err = db.Where("category_id = ?", category).Find(&program).Order("created asc").Count(total).Limit(*size).Offset(*page).Rows()
//				if err != nil {
//					panic(err)
//				}
//			}
//		case 2:
//			if page != nil && size != nil && category == nil{
//				rows, err = db.Find(&program).Order("created desc").Count(total).Limit(*size).Offset(*page).Rows()
//				if err != nil {
//					panic(err)
//				}
//			}
//			if category != nil && page != nil && size != nil{
//				rows, err = db.Where("category_id = ?", category).Find(&program).Order("created desc").Count(total).Limit(*size).Offset(*page).Rows()
//				if err != nil {
//					panic(err)
//				}
//			}
//		case 3:
//			if page != nil && size != nil && category == nil{
//				rows, err = db.Find(&program).Order("program_name asc").Count(total).Limit(*size).Offset(*page).Rows()
//				if err != nil {
//					panic(err)
//				}
//			}
//			if category != nil && page != nil && size != nil{
//				rows, err = db.Where("category_id = ?", category).Find(&program).Order("program_name asc").Count(total).Limit(*size).Offset(*page).Rows()
//				if err != nil {
//					panic(err)
//				}
//			}
//		case 4:
//			if page != nil && size != nil && category == nil{
//				rows, err = db.Find(&program).Order("program_name desc").Count(total).Limit(*size).Offset(*page).Rows()
//				if err != nil {
//					panic(err)
//				}
//			}
//			if category != nil && page != nil && size != nil{
//				rows, err = db.Where("category_id = ?", category).Find(&program).Order("program_name desc").Count(total).Limit(*size).Offset(*page).Rows()
//				if err != nil {
//					panic(err)
//				}
//			}
//		}
//	}else {
//		if page != nil && size != nil {
//			rows, err = db.Find(&program).Order("created desc").Count(total).Limit(*size).Offset(*page).Rows()
//			if err != nil {
//				panic(err)
//			}
//		} else{
//			rows, err = db.Find(&program).Rows()
//			if err != nil {
//				panic(err)
//			}
//		}
//	}
//	if id != nil {
//		rows, err = db.Where("id = ?", id).First(&program).Rows()
//		if err != nil{
//			panic(err)
//		}
//	}
//
//	result := make([]model.Program, 0)
//
//	for rows.Next() {
//		t := &model.Program{}
//		fmt.Println(t)
//		benefitmemory := t.Benefit
//		fmt.Println(&benefitmemory)
//
//
//		err = rows.Scan(
//			&t.Id,
//			&t.Created,
//			&t.CreatedBy,
//			&t.Modified,
//			&t.ModifiedBy,
//			&t.Active,
//			&t.IsDeleted,
//			&t.Deleted,
//			&t.Deleted_by,
//			&t.ProgramName,
//			&t.ProgramImage,
//			&t.ProgramStartDate,
//			&t.ProgramEndDate,
//			&t.ProgramDescription,
//			&t.Card,
//			//&t.OutletID,
//			&t.MerchantId,
//			&t.CategoryId,
//			&t.Benefit,
//			&t.TermsAndCondition,
//			&t.Tier,
//			&t.RedeemRules,
//			&t.RewardTarget,
//			&t.QRCodeId,
//			&t.ProgramPoint,
//			&t.MinPayment,
//		)
//
//		//add alert
//		merchant := new  (model.Merchant)
//
//		db.Table("merchants").
//			Select("merchants.merchant_name").
//			Where("id = ?", t.MerchantId).
//			First(&merchant)
//		//t.MerchantName = merchant.MerchantName
//		if err != nil {
//			logrus.Error(err)
//			return nil
//		}
//		result = append(result,*t)
//	}
//	db.Close()
//	return err
//}

func GetProgram(page *int, size *int, sort *int, category *int, email *string) []model.Program {
	fmt.Println("masuk ke Get Program")
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var program []model.Program
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

	result := make([]model.Program, 0)

	for rows.Next() {
		t := &model.Program{}
		fmt.Println(t)
		benefitmemory := t.Benefit
		fmt.Println(&benefitmemory)

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
			&t.ProgramPoint,
			&t.MinPayment,
			&t.IsReqBillNumber,
			&t.IsReqTotalTransaction,
			&t.IsPushNotification,
			&t.IsLendCard,
			&t.IsGiveCard,
			&t.IsWelcomeBonus,
			&t.MerchantName,
			&t.OutletName,
			&t.CategoryName,
		)
		//add alert
		merchant := new(model.Merchant)

		db.Table("merchants").
			Select("merchants.merchant_name").
			Where("merchant_email = ?", t.MerchantEmail).
			First(&merchant)
		t.MerchantName = merchant.MerchantName

		category := new(model.MerchantCategory)

		db.Table("merchant_categories").
			Select("merchant_categories.category_name").
			Where("id = ? ", t.CategoryId).
			First(&category)
		t.CategoryName = category.CategoryName

		outlet := new(model.Outlet2)

		db.Table("outlet2").
			Select("outlet2.outlet_name").
			Where("id = ? ", t.OutletID).
			First(&outlet)
		t.OutletName = outlet.OutletName

		if err != nil {
			logrus.Error(err)
			return nil
		}
		result = append(result, *t)
	}
	db.Close()
	return result
}
