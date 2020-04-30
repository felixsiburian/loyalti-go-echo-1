package repository

import (
	"database/sql"
	"github.com/labstack/gommon/log"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"
)

func CreateReward(reward *model.Reward) string {
	db := database.ConnectionDB()
	rewardobj := model.Reward{
		Created:           time.Now(),
		CreatedBy:         "Admin",
		Modified:          time.Now(),
		ModifiedBy:        "Admin",
		Active:            true,
		IsDeleted:         false,
		Deleted:           nil,
		DeletedBy:         "",
		RedeemPoints:      reward.RedeemPoints,
		RewardName:        reward.RewardName,
		RedeemRules:       reward.RedeemRules,
		TermsAndCondition: reward.TermsAndCondition,
		ProgramId:         reward.ProgramId,
		MerchantEmail:     reward.MerchantEmail,
		Outletid:          reward.Outletid,
		MerchantName:      "a",
		ProgramName:       "b",
		OutletName:        "c",
	}
	db.Create(&rewardobj)
	defer db.Close()
	return "reward berhasil dibuat"
}

func UpdateReward(reward *model.Reward) string {
	db := database.ConnectionDB()
	db.Model(&reward).Updates(map[string]interface{}{
		"redeem_points":       reward.RedeemPoints,
		"reward_name":         reward.RewardName,
		"redeem_rules":        reward.RedeemRules,
		"terms_and_condition": reward.TermsAndCondition,
		"program_id":          reward.ProgramId,
		"merchant_email":      reward.MerchantEmail,
		"outletid":            reward.Outletid,
	})
	defer db.Close()
	return "Update Berhasil"
}

func DeleteReward(reward *model.Reward) string {
	db := database.ConnectionDB()
	db.Model(&reward).Where("id = ?", reward.Id).Update("active", false)
	db.Model(&reward).Where("id = ?", reward.Id).Update("is_deleted", true)

	defer db.Close()
	return "Berhasil dihapus"
}

func GetReward(page *int, size *int, sort *int, merchant_email *string) []model.Reward {
	db := database.ConnectionDB()
	var reward []model.Reward
	var rows *sql.Rows
	var err error
	var total int

	if page == nil && size == nil && merchant_email == nil {
		rows, err = db.Find(&reward).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil && sort != nil && merchant_email != nil {
		switch *sort {
		case 1:
			rows, err = db.Find(&reward).Where("merchant_email = ?", merchant_email).Order("reward_name asc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		case 2:
			rows, err = db.Find(&reward).Where("merchant_email = ?", merchant_email).Order("reward_name desc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if merchant_email != nil {
		rows, err = db.Find(&reward).Where("merchant_email = ? ", merchant_email).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	result := make([]model.Reward, 0)

	for rows.Next() {
		r := &model.Reward{}

		err = rows.Scan(
			&r.Id,
			&r.Created,
			&r.CreatedBy,
			&r.ModifiedBy,
			&r.ModifiedBy,
			&r.Active,
			&r.IsDeleted,
			&r.Deleted,
			&r.DeletedBy,
			&r.RedeemPoints,
			&r.RewardName,
			&r.RedeemRules,
			&r.TermsAndCondition,
			&r.ProgramId,
			&r.MerchantEmail,
			&r.Outletid,
			&r.MerchantName,
			&r.ProgramName,
			&r.OutletName,
		)

		merchant := new(model.Merchant)
		db.Table("merchants").
			Select("merchants.merchant_name").
			Where("merchant_email = ?", r.MerchantEmail).
			First(&merchant)
		r.MerchantName = merchant.MerchantName

		program := new(model.Program)
		db.Table("programs").
			Select("programs.program_name").
			Where("id = ?", r.ProgramId).
			First(&program)
		r.ProgramName = program.ProgramName

		outlet := new(model.Outlet2)
		db.Table("outlet2").
			Select("outlet2.outlet_name").
			Where("id = ? ", r.Outletid).
			First(&outlet)
		r.OutletName = outlet.OutletName

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, *r)
	}

	db.Close()
	return result
}
