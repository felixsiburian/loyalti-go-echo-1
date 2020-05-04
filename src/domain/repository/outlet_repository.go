package repository

import (
	"database/sql"
	"fmt"
	"github.com/labstack/gommon/log"
	"time"

	//"github.com/biezhi/gorm-paginator/pagination"

	//"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

type OutletRepository interface {
	CreateOutlet(newoutlet *model.Outlet2) error
	UpdateOutlet(newoutlet *model.Outlet2) error
	DeleteOutlet(newoutlet *model.Outlet2) error
}

type outlet_repo struct {
	DB *gorm.DB
	//database.Connection_Interface
}

//func (p *outlet_repo) CreateOutlet(newoutlet *model.Outlet) error {
//	fmt.Println("masuk fungsi")
//	outlet := model.Outlet{
//		Created:          time.Now(),
//		CreatedBy:        "",
//		Modified:         time.Now(),
//		ModifiedBy:       "",
//		Active:           true,
//		IsDeleted:        false,
//		Deleted:          nil,
//		Deleted_by:       "",
//		OutletName:       newoutlet.OutletName,
//		OutletAddress:    newoutlet.OutletAddress,
//		OutletPhone:      newoutlet.OutletPhone,
//		OutletCity:       newoutlet.OutletCity,
//		OutletProvince:   newoutlet.OutletProvince,
//		OutletPostalCode: newoutlet.OutletPostalCode,
//		OutletLongitude:  newoutlet.OutletLongitude,
//		OutletLatitude:   newoutlet.OutletLatitude,
//		OutletDay:        time.Time{},
//		OutletHour:       time.Time{},
//		MerchantId:       []model.Merchant{},
//	}
//
//	db := database.ConnectionDB()
//	err := db.Create(&outlet).Error
//	if err != nil {
//		fmt.Println("Tak ada error")
//	}
//	return err
//}

//func CreateOutletRepository(db *gorm.DB) OutletRepository {
//	return &outlet_repo{
//		DB: db,
//	}
//}

//func (p *outlet_repo) UpdateOutlet(newoutlet *model.Outlet) error {
//	db := database.ConnectionDB()
//	outlet := model.Outlet{
//		Created:          time.Time{},
//		CreatedBy:        "",
//		Modified:         time.Time{},
//		ModifiedBy:       "",
//		Active:           false,
//		IsDeleted:        false,
//		Deleted:          nil,
//		Deleted_by:       "",
//		OutletName:       newoutlet.OutletName,
//		OutletAddress:    newoutlet.OutletAddress,
//		OutletPhone:      newoutlet.OutletPhone,
//		OutletCity:       newoutlet.OutletCity,
//		OutletProvince:   newoutlet.OutletProvince,
//		OutletPostalCode: newoutlet.OutletPostalCode,
//		OutletLongitude:  newoutlet.OutletLongitude,
//		OutletLatitude:   newoutlet.OutletLatitude,
//		OutletDay:        time.Time{},
//		OutletHour:       time.Time{},
//		MerchantId:       newoutlet.MerchantId,
//	}
//	err := db.Model(&outlet).Where("merchant_id = ?", outlet.MerchantId).Update(&outlet).Error
//	return err
//}

func (p *outlet_repo) DeleteOutlet(newoutlet *model.Outlet2) error {
	db := database.ConnectionDB()

	err := db.Model(&newoutlet).Where("id = ?", newoutlet.Id).Update("active", false).Error
	if err == nil {
		fmt.Println("tidak ada error")
	}
	return err
}

func CreateOutlet(outlet *model.Outlet2) string {
	db := database.ConnectionDB()
	outletObj := model.Outlet2{
		Created:          time.Now(),
		CreatedBy:        "Admin",
		Modified:         time.Now(),
		ModifiedBy:       "Admin",
		Active:           true,
		IsDeleted:        false,
		Deleted:          time.Now(),
		Deleted_by:       "",
		OutletName:       outlet.OutletName,
		OutletAddress:    outlet.OutletAddress,
		OutletPhone:      outlet.OutletPhone,
		OutletCity:       outlet.OutletCity,
		OutletProvince:   outlet.OutletProvince,
		OutletPostalCode: outlet.OutletPostalCode,
		OutletLongitude:  outlet.OutletLongitude,
		OutletLatitude:   outlet.OutletLatitude,
		OutletDay:        outlet.OutletDay,
		OutletHour:       outlet.OutletHour,
		MerchantEmail:    outlet.MerchantEmail,
		Timezone:         outlet.Timezone,
		MerchantName:     "a",
	}

	db.Create(&outletObj)
	defer db.Close()
	return outletObj.OutletName
}

func UpdateOutlet(outlet *model.Outlet2) string {
	db := database.ConnectionDB()
	db.Model(&outlet).Updates(map[string]interface{}{
		"outlet_name": outlet.OutletName,
		"outlet_address":outlet.OutletAddress,
		"outlet_phone":outlet.OutletPhone,
		"outlet_city":outlet.OutletCity,
		"outlet_province":outlet.OutletProvince,
		"outlet_postal_code": outlet.OutletPostalCode,
		"outlet_longitude": outlet.OutletLongitude,
		"outlet_latitude": outlet.OutletLatitude,
		"outlet_day":outlet.OutletDay,
		"outlet_hour":outlet.OutletHour,
		"merchant_email":outlet.MerchantEmail,
		"timezone":outlet.Timezone,
	})
	outlet.Modified = time.Now()
	defer db.Close()
	return outlet.OutletName
}

func DeleteOutlet(outlet *model.Outlet2) string {
	db := database.ConnectionDB()
	db.Model(&outlet).Where("id= ?", outlet.Id).Update("active", false)
	db.Model(&outlet).Where("id= ?", outlet.Id).Update("is_deleted", true)
	outlet.Deleted = time.Now()
	defer db.Close()
	return "berhasil dihapus"
}

func GetOutlet(page *int, size *int, id *int, email *string) []model.Outlet2 {
	fmt.Println("masuk ke get outlet")
	db := database.ConnectionDB()
	var outlet []model.Outlet2
	var rows *sql.Rows
	var err error
	var total int

	if page == nil && size == nil && id == nil && email == nil {
		rows, err = db.Find(&outlet).Order("outlet_name asc").Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil && id != nil && email != nil {
		rows, err = db.Find(&outlet).Where("merchant_email = ? AND id = ?", email, id).Order("outlet_name asc").Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil && email != nil && id == nil {
		rows, err = db.Find(&outlet).Where("merchant_email = ?", email).Order("outlet_name asc").Order(total).Limit(*size).Offset(*page).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil && email == nil && id != nil {
		rows, err = db.Find(&outlet).Where("id = ? ", id).Order("outlet_name asc").Order(total).Limit(*size).Offset(*page).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil && email == nil && id == nil {
		rows, err = db.Find(&outlet).Order("outlet_name asc").Count(total).Limit(*size).Offset(*page).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page == nil && size == nil && id == nil && email != nil {
		rows, err = db.Find(&outlet).Where("merchant_email = ?", email).Order("outlet_name asc").Count(total).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

		result := make([]model.Outlet2, 0)
	fmt.Println("lewat")
	for rows.Next() {
		fmt.Println("masuk")
		o := &model.Outlet2{}
		fmt.Println(o)

		err = rows.Scan(
			&o.Id,
			&o.Created,
			&o.CreatedBy,
			&o.Modified,
			&o.ModifiedBy,
			&o.Active,
			&o.IsDeleted,
			&o.Deleted,
			&o.Deleted_by,
			&o.OutletName,
			&o.OutletAddress,
			&o.OutletPhone,
			&o.OutletCity,
			&o.OutletProvince,
			&o.OutletPostalCode,
			&o.OutletLongitude,
			&o.OutletLatitude,
			&o.OutletDay,
			&o.OutletHour,
			&o.MerchantEmail,
			&o.Timezone,
			&o.MerchantName,
		)

		merchant := new(model.Merchant)
		db.Table("merchants").Select("merchants.merchant_name").
			Where("merchant_email = ? ", o.MerchantEmail).First(&merchant)
		o.MerchantName = merchant.MerchantName

		if err != nil {
			log.Fatal(err)
		}
		result = append(result, *o)
	}
	db.Close()
	return result
}
