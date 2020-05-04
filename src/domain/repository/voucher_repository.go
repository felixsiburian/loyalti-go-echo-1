package repository

import (
	"database/sql"
	"github.com/labstack/gommon/log"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"
)

//Create Voucher
func CreateVoucher(voucher *model.Voucher) string {
	db := database.ConnectionDB()
	newvoucher := model.Voucher{
		Created:                  time.Now(),
		CreatedBy:                "Admin",
		Modified:                 time.Now(),
		ModifiedBy:               "Admin",
		Active:                   true,
		IsDeleted:                false,
		Deleted:                  time.Now(),
		DeletedBy:                "",
		VoucherName:              voucher.VoucherName,
		StartDate:                voucher.StartDate,
		EndDate:                  voucher.EndDate,
		VoucherDescription:       voucher.VoucherDescription,
		VoucherTermsAndCondition: voucher.VoucherTermsAndCondition,
		IsPushNotification:       voucher.IsPushNotification,
		IsGiveVoucher:            voucher.IsGiveVoucher,
		VoucherPeriod:            voucher.VoucherPeriod,
		RewardTermsAndCondition:  voucher.VoucherTermsAndCondition,
		BackgroundVoucherPattern: voucher.BackgroundVoucherPattern,
		BackgroundVoucherColour:  voucher.BackgroundVoucherColour,
		MerchantEmail: voucher.MerchantEmail,
		OutletId:     voucher.OutletId,
		ProgramId:    voucher.ProgramId,
		MerchantName: "a",
		OutletName:   "b",
		ProgramName:  "c",
	}
	db.Create(&newvoucher)
	defer db.Close()
	return "voucher berhasil dibuat"
}

//Update Voucher using program id
func UpdateVoucher(voucher *model.Voucher) string {
	db := database.ConnectionDB()
	db.Model(&voucher).Updates(map[string]interface{}{
		"voucher_name":                voucher.VoucherName,
		"voucher_description":         voucher.VoucherDescription,
		"voucher_terms_and_condition": voucher.VoucherTermsAndCondition,
		"is_push_notification":        voucher.IsPushNotification,
		"is_give_voucher":             voucher.IsGiveVoucher,
		"voucher_period":              voucher.VoucherPeriod,
		"reward_terms_and_condition":  voucher.RewardTermsAndCondition,
		"background_voucher":          voucher.BackgroundVoucherPattern,
		"background_voucher_colour":   voucher.BackgroundVoucherColour,
		"merchant_email":              voucher.MerchantEmail,
		"outlet_id":                   voucher.OutletId,
		"program_id":                  voucher.ProgramId,
	})
	voucher.Modified = time.Now()
	defer db.Close()
	return "Update Berhasil"
}

//Delete Vouocher using program id
func DeleteVoucher(voucher *model.Voucher) string {
	db := database.ConnectionDB()
	db.Model(&voucher).Where("program_id = ?", voucher.ProgramId).Update("active", false)
	db.Model(&voucher).Where("program_id = ?", voucher.ProgramId).Update("is_deleted", true)
	voucher.Deleted = time.Now()
	defer db.Close()
	return "berhasil dihapus"
}

//Get Voucher by Merchant_id and have sorting
func GetVoucher(page *int, size *int, sort *int, merchant_email *string) []model.Voucher {
	db := database.ConnectionDB()
	var voucher []model.Voucher
	var rows *sql.Rows
	var err error
	var total int

	if page == nil && size == nil && sort == nil && merchant_email == nil {
		rows, err = db.Find(&voucher).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil && sort != nil && merchant_email != nil {
		switch *sort {
		case 1:
			rows, err = db.Find(&voucher).Where("merchant_email = ?", merchant_email).Order("voucher_name asc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		case 2:
			rows, err = db.Find(&voucher).Where("merchant_email = ?", merchant_email).Order("voucher_name desc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if page != nil && size != nil && sort == nil && merchant_email != nil {
		rows, err = db.Find(&voucher).Where("merchant_email = ?", merchant_email).Order("voucher_name asc").Count(total).Limit(*size).Offset(*page).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil && sort == nil && merchant_email == nil {
		rows, err = db.Find(&voucher).Order("voucher_name asc").Count(total).Limit(*size).Offset(*page).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil && sort != nil && merchant_email == nil {
		switch *sort {
		case 1:
			rows, err = db.Find(&voucher).Order("voucher_name asc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		case 2:
			rows, err = db.Find(&voucher).Order("voucher_name desc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if page == nil && size == nil && sort == nil && merchant_email != nil {
		rows, err = db.Find(&voucher).Where("merchant_email = ?", merchant_email).Order("voucher_name asc").Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	result := make([]model.Voucher, 0)

	for rows.Next() {
		v := &model.Voucher{}

		err = rows.Scan(
			&v.Id,
			&v.Created,
			&v.CreatedBy,
			&v.Modified,
			&v.ModifiedBy,
			&v.Active,
			&v.IsDeleted,
			&v.Deleted,
			&v.DeletedBy,
			&v.VoucherName,
			&v.StartDate,
			&v.EndDate,
			&v.VoucherDescription,
			&v.VoucherTermsAndCondition,
			&v.IsPushNotification,
			&v.IsGiveVoucher,
			&v.VoucherPeriod,
			&v.RewardTermsAndCondition,
			&v.BackgroundVoucherPattern,
			&v.BackgroundVoucherColour,
			&v.MerchantEmail,
			&v.OutletId,
			&v.ProgramId,
			&v.MerchantName,
			&v.OutletName,
			&v.ProgramName,
		)

		merchant := new(model.Merchant)
		db.Table("merchants").
			Select("merchants.merchant_name").
			Where("merchant_email = ?", v.MerchantEmail).
			First(&merchant)
		v.MerchantName = merchant.MerchantName

		outlet := new(model.Outlet2)
		db.Table("outlet2").
			Select("outlet2.outlet_name").
			Where("id = ? ", v.OutletId).
			First(&outlet)
		v.OutletName = outlet.OutletName

		program := new(model.Program)
		db.Table("programs").
			Select("programs.program_name").
			Where("id = ? ", v.ProgramId).
			First(&program)
		v.ProgramName = program.ProgramName

		if err != nil {
			log.Fatal(err)
		}
		result = append(result, *v)
	}
	db.Close()
	return result

}
