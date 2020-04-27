package repository

import (
	"fmt"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"
)

type OutletRepository interface {
	CreateOutlet(newoutlet *model.Outlet) error
	UpdateOutlet(newoutlet *model.Outlet) error
	DeleteOutlet(newoutlet *model.Outlet) error
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

func (p *outlet_repo) DeleteOutlet(newoutlet *model.Outlet) error {
	db := database.ConnectionDB()

	err := db.Model(&newoutlet).Where("id = ?", newoutlet.Id).Update("active", false).Error
	if err == nil {
		fmt.Println("tidak ada error")
	}
	return err
}

func CreateOutlet(outlet *model.Outlet) string {
	db := database.ConnectionDB()
	outlet.Created = time.Now()
	outlet.Modified = time.Now()
	outletObj := *outlet

	db.Create(&outletObj)
	defer db.Close()
	return outletObj.OutletName
}

func UpdateOutlet(outlet *model.Outlet) string {
	db := database.ConnectionDB()
	db.Model(&outlet).Where("id = ?", outlet.Id).Update(&outlet)
	defer db.Close()
	return outlet.OutletName
}

func DeleteOutlet(outlet *model.Outlet) string {
	db := database.ConnectionDB()
	db.Model(&outlet).Where("id= ?", outlet.Id).Update("active", false)
	defer db.Close()
	return "berhasil dihapus"
}

func GetOutlet(page *int, size *int, id *int, email *string) []model.Outlet {
	fmt.Println("masuk ke get outlet")
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var outlet []model.Outlet
	//var result map[string]interface{}

	db.Find(&outlet)
	fmt.Println("lewat 1")
	if id == nil && size == nil && page == nil && email == nil {
		fmt.Println("1")
		db.Model(&outlet).Find(&outlet)
		fmt.Println("habis")
	}
	fmt.Println("lewat 2")
	if id == nil && size != nil && page != nil && email != nil {
		fmt.Println("2")
		db.Model(&outlet).Where("merchant_email = ?", email).Limit(*size).Offset(*page).Find(&outlet)
		db.Model(&outlet).Find(&outlet)
		pagination.Paging(&pagination.Param{
			DB:      db,
			Page:    *page,
			Limit:   *size,
			OrderBy: []string{"outlet_name desc"},
		}, &outlet)
	}
	fmt.Println("lewat 3")
	if id == nil && size != nil && page != nil && email == nil {
		fmt.Println("3")
		db.Model(&outlet).Find(&outlet)
		pagination.Paging(&pagination.Param{
			DB:      db,
			Page:    *page,
			Limit:   *size,
			OrderBy: []string{"outlet_name asc"},
		}, &outlet)
	}
	fmt.Println("lewat 4")
	if id != nil && size != nil && page != nil && email == nil {
		fmt.Println("4")
		db.Model(&outlet).Where("merchant_id =  ?", id).Find(&outlet)
		pagination.Paging(&pagination.Param{
			DB:      db,
			Page:    *page,
			Limit:   *size,
			OrderBy: []string{"outlet_name asc"},
		}, &outlet)
	}
	fmt.Println("lewat 5")
	//rows,err := db.Table("merchants").Select("merchants.merchant_name, outlets.outlet_name").Joins("left join outlets on outlets.merchant_email = merchants.merchant_email").Rows()
	//if err != nil {
	//	fmt.Println("Error join : ", err.Error())
	//	os.Exit(1)
	//}
	//
	//result := make([]model.Outlet, 0)
	//
	//for rows.Next() {
	//	fmt.Println("masuk ke perulangan")
	//	o := &model.Outlet{}
	//
	//	err = rows.Scan(
	//		&o.OutletName,
	//		&o.MerchantEmail,
	//		)
	//
	//	outlet := new (model.Outlet)
	//	db.Table("merchants").
	//		Select("merchants.merchant_name").
	//		Where("merchant_email = ?", o.MerchantEmail).
	//		First(&outlet)
	//
	//	if err != nil {
	//		fmt.Println("error join 2 : ", err.Error())
	//		log.Fatal(err)
	//	}
	//	result = append(result, *o)
	//	fmt.Println(result)
	//}
	//
	//res := db.Table("merchants").Select("merchants.merchant_name, outlets.outlet_name").Joins("left join outlets on outlets.merchant_email = merchants.merchant_email").Scan(&result)
	//fmt.Println(res)

	db.Close()
	return outlet
}
