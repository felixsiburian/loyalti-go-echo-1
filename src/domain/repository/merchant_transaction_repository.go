package repository

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"
)

type TransactionRepository interface {
	CreateTransaction(transaction *model.TransactionMerchant) error
}

type repoTransaction struct {
	DB *gorm.DB
}

func CreateRepositoryTransaction(db *gorm.DB) TransactionRepository {
	return &repoTransaction{
		DB: db,
	}
}

func (p *repoTransaction) CreateTransaction(transaction *model.TransactionMerchant) error {
	db := database.ConnectionDB()
	transactionObj := *transaction
	err := db.Create(&transactionObj).Error
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
	return err
}

func CreateTransaction(transaction *model.TransactionMerchant) string {
	db := database.ConnectionDB()
	transactionObj := model.TransactionMerchant{
		Created:          time.Now(),
		CreatedBy:        "Admin",
		Modified:         time.Now(),
		ModifiedBy:       "Admin",
		Active:           true,
		IsDeleted:        false,
		Deleted:          nil,
		Deleted_by:       "",
		MerchantEmail:    transaction.MerchantEmail,
		OutletId:         transaction.OutletId,
		TotalTransaction: transaction.TotalTransaction,
		PointTransaction: transaction.PointTransaction,
		BillNumber:       transaction.BillNumber,
		MerchantName:     "a",
		OutletName:       "b",
	}
	db.Create(&transactionObj)
	defer db.Close()
	return transactionObj.BillNumber
}

func UpdateTransaction(transaction *model.TransactionMerchant) string {
	db := database.ConnectionDB()
	db.Model(&transaction).Updates(map[string]interface{}{
		"merchant_email":    transaction.MerchantEmail,
		"outlet_id":         transaction.OutletId,
		"total_transaction": transaction.TotalTransaction,
		"point_transaction": transaction.PointTransaction,
	})
	defer db.Close()
	return transaction.BillNumber
}

func DeleteTransaction(transaction *model.TransactionMerchant) string {
	db := database.ConnectionDB()
	db.Model(&transaction).Where("id = ?", transaction.Id).Update("active", false)
	db.Model(&transaction).Where("id = ?", transaction.Id).Update("is_deleted", true)
	defer db.Close()
	return "berhasil dihapus"
}

func GetTransaction(page *int, size *int, sort *int, outletid *int) []model.TransactionMerchant {
	fmt.Println("Masuk ke get")
	db := database.ConnectionDB()
	var transaction []model.TransactionMerchant
	var rows *sql.Rows
	var err error
	var total int

	if page == nil && size == nil && sort == nil && outletid == nil {
		rows, err = db.Find(&transaction).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil && sort != nil && outletid == nil {
		switch *sort {
		case 1:
			rows, err = db.Find(&transaction).Order("created asc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		case 2:
			rows, err = db.Find(&transaction).Order("created desc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if page != nil && size != nil && sort != nil && outletid != nil {
		switch *sort {
		case 1:
			rows, err = db.Find(&transaction).Where("outlet_id = ?", outletid).Order("created asc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		case 2:
			rows, err = db.Find(&transaction).Where("outlet_id = ?", outletid).Order("created desc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if page != nil && size != nil && sort == nil && outletid != nil {
		rows, err = db.Find(&transaction).Where("outlet_id = ?", outletid).Count(total).Limit(*size).Offset(*page).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page == nil && size == nil && sort == nil && outletid != nil {
		rows, err = db.Find(&transaction).Where("outlet_id = ?", outletid).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	result := make([]model.TransactionMerchant, 0)

	for rows.Next() {
		tr := &model.TransactionMerchant{}
		err = rows.Scan(
			&tr.Id,
			&tr.Created,
			&tr.CreatedBy,
			&tr.Modified,
			&tr.ModifiedBy,
			&tr.Active,
			&tr.IsDeleted,
			&tr.Deleted,
			&tr.Deleted_by,
			&tr.MerchantEmail,
			&tr.OutletId,
			&tr.TotalTransaction,
			&tr.PointTransaction,
			&tr.BillNumber,
			&tr.MerchantName,
			&tr.OutletName,
		)

		merchant := new(model.Merchant)
		db.Table("merchants").
			Select("merchants.merchant_name").
			Where("merchant_email = ?", tr.MerchantEmail).
			First(&merchant)
		tr.MerchantName = merchant.MerchantName

		outlet := new(model.Outlet2)
		db.Table("outlet2").
			Select("outlet2.outlet_name").
			Where("id = ? ", tr.OutletId).
			First(&outlet)
		tr.OutletName = outlet.OutletName

		if err != nil {
			log.Fatal(err)
		}
		result = append(result, *tr)
	}

	db.Close()
	return result
}
