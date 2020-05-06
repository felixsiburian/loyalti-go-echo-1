package repository

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"
)

type Repository interface {
	CreateMerchant(newmerchant *model.Merchant) error
	UpdateMerchant(newmerchant *model.Merchant) error
	DeleteMerchant(newmerchant *model.Merchant) error
}

type repo struct {
	DB *gorm.DB
}

func (p *repo) CreateMerchant(newmerchant *model.Merchant) error {
	merchant := model.Merchant{
		Created:               time.Now(),
		CreatedBy:             "",
		Modified:              time.Now(),
		ModifiedBy:            "",
		Active:                true,
		IsDeleted:             false,
		Deleted:               time.Time{},
		Deleted_by:            "",
		MerchantName:          newmerchant.MerchantName,
		MerchantEmail:         newmerchant.MerchantEmail,
		MerchantPhoneNumber:   newmerchant.MerchantPhoneNumber,
		MerchantProvince:      newmerchant.MerchantProvince,
		MerchantCity:          newmerchant.MerchantCity,
		MerchantAddress:       newmerchant.MerchantAddress,
		MerchantPostalCode:    newmerchant.MerchantPostalCode,
		MerchantCategoryId:    newmerchant.MerchantCategoryId,
		MerchantWebsite:       newmerchant.MerchantWebsite,
		MerchantMediaSocialId: newmerchant.MerchantMediaSocialId,
		MerchantDescription:   newmerchant.MerchantDescription,
		MerchantImageProfile:  newmerchant.MerchantImageProfile,
		MerchantGallery:       newmerchant.MerchantGallery,
	}
	db := database.ConnectionDB()
	err := db.Create(&merchant).Error
	if err == nil {
		fmt.Println("tak ada error")
	}
	return err
}

func CreateRepository(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func CreateMerchant2(newmerchant *model.Merchant) string {
	merchant := model.Merchant{
		Created:               time.Now(),
		CreatedBy:             "Admin",
		Modified:              time.Now(),
		ModifiedBy:            "Admin",
		Active:                true,
		IsDeleted:             false,
		Deleted:               time.Now(),
		Deleted_by:            "",
		MerchantName:          newmerchant.MerchantName,
		MerchantEmail:         newmerchant.MerchantEmail,
		MerchantPhoneNumber:   newmerchant.MerchantPhoneNumber,
		MerchantProvince:      newmerchant.MerchantProvince,
		MerchantCity:          newmerchant.MerchantCity,
		MerchantAddress:       newmerchant.MerchantAddress,
		MerchantPostalCode:    newmerchant.MerchantPostalCode,
		MerchantCategoryId:    newmerchant.MerchantCategoryId,
		MerchantWebsite:       newmerchant.MerchantWebsite,
		MerchantMediaSocialId: newmerchant.MerchantMediaSocialId,
		MerchantDescription:   newmerchant.MerchantDescription,
		MerchantImageProfile:  newmerchant.MerchantImageProfile,
		MerchantGallery:       newmerchant.MerchantGallery,
		CategoryName:          "a",
		Rating:                0,
	}

	db := database.ConnectionDB()
	db.Create(&merchant)
	defer db.Close()
	return merchant.MerchantEmail
}

func (p *repo) UpdateMerchant(newmerchant *model.Merchant) error {
	db := database.ConnectionDB()
	merchant := model.Merchant{
		Created:               time.Time{},
		CreatedBy:             "",
		Modified:              time.Now(),
		ModifiedBy:            "",
		Active:                true,
		IsDeleted:             false,
		Deleted:               time.Now(),
		Deleted_by:            "",
		MerchantName:          newmerchant.MerchantName,
		MerchantEmail:         newmerchant.MerchantEmail,
		MerchantPhoneNumber:   newmerchant.MerchantPhoneNumber,
		MerchantProvince:      newmerchant.MerchantProvince,
		MerchantCity:          newmerchant.MerchantCity,
		MerchantAddress:       newmerchant.MerchantAddress,
		MerchantPostalCode:    newmerchant.MerchantPostalCode,
		MerchantCategoryId:    newmerchant.MerchantCategoryId,
		MerchantWebsite:       newmerchant.MerchantWebsite,
		MerchantMediaSocialId: newmerchant.MerchantMediaSocialId,
		MerchantDescription:   newmerchant.MerchantDescription,
		MerchantImageProfile:  newmerchant.MerchantImageProfile,
		MerchantGallery:       newmerchant.MerchantGallery,
	}
	err := db.Model(&merchant).Where("merchant_email = ?", merchant.MerchantEmail).Update(&merchant).Error
	return err
}

func UpdateMerchant2(newmerchant *model.Merchant) string {
	db := database.ConnectionDB()
	merchant := model.Merchant{
		Created:               time.Time{},
		CreatedBy:             "Admin",
		Modified:              time.Now(),
		ModifiedBy:            "",
		Active:                true,
		IsDeleted:             false,
		Deleted:               time.Now(),
		Deleted_by:            "",
		MerchantName:          newmerchant.MerchantName,
		MerchantEmail:         newmerchant.MerchantEmail,
		MerchantPhoneNumber:   newmerchant.MerchantPhoneNumber,
		MerchantProvince:      newmerchant.MerchantProvince,
		MerchantCity:          newmerchant.MerchantCity,
		MerchantAddress:       newmerchant.MerchantAddress,
		MerchantPostalCode:    newmerchant.MerchantPostalCode,
		MerchantCategoryId:    newmerchant.MerchantCategoryId,
		MerchantWebsite:       newmerchant.MerchantWebsite,
		MerchantMediaSocialId: newmerchant.MerchantMediaSocialId,
		MerchantDescription:   newmerchant.MerchantDescription,
		MerchantImageProfile:  newmerchant.MerchantImageProfile,
		MerchantGallery:       newmerchant.MerchantGallery,
	}
	db.Model(&merchant).Where("merchant_email = ?", merchant.MerchantEmail).Update(&merchant)
	defer db.Close()
	return merchant.MerchantEmail
}

func (p *repo) DeleteMerchant(newmerchant *model.Merchant) error {
	db := database.ConnectionDB()
	merchant := model.Merchant{
		Created:               time.Time{},
		CreatedBy:             "",
		Modified:              time.Now(),
		ModifiedBy:            "",
		Active:                true,
		IsDeleted:             false,
		Deleted:               time.Time{},
		Deleted_by:            "",
		MerchantName:          newmerchant.MerchantName,
		MerchantEmail:         newmerchant.MerchantEmail,
		MerchantPhoneNumber:   newmerchant.MerchantPhoneNumber,
		MerchantProvince:      newmerchant.MerchantProvince,
		MerchantCity:          newmerchant.MerchantCity,
		MerchantAddress:       newmerchant.MerchantAddress,
		MerchantPostalCode:    newmerchant.MerchantPostalCode,
		MerchantCategoryId:    newmerchant.MerchantCategoryId,
		MerchantWebsite:       newmerchant.MerchantWebsite,
		MerchantMediaSocialId: newmerchant.MerchantMediaSocialId,
		MerchantDescription:   newmerchant.MerchantDescription,
		MerchantImageProfile:  newmerchant.MerchantImageProfile,
		MerchantGallery:       newmerchant.MerchantGallery,
	}
	err := db.Model(&merchant).Where("merchant_email = ?", merchant.MerchantEmail).Update("active", false).Error
	merchant.Modified = time.Now()
	if err == nil {
		fmt.Println("Tidak ada error")
	}
	return err
}

func DeleteMerchant2(newmerchant *model.Merchant) string {
	db := database.ConnectionDB()
	merchant := model.Merchant{
		Created:               time.Time{},
		CreatedBy:             "",
		Modified:              time.Now(),
		ModifiedBy:            "",
		Active:                true,
		IsDeleted:             false,
		Deleted:               time.Now(),
		Deleted_by:            "",
		MerchantName:          newmerchant.MerchantName,
		MerchantEmail:         newmerchant.MerchantEmail,
		MerchantPhoneNumber:   newmerchant.MerchantPhoneNumber,
		MerchantProvince:      newmerchant.MerchantProvince,
		MerchantCity:          newmerchant.MerchantCity,
		MerchantAddress:       newmerchant.MerchantAddress,
		MerchantPostalCode:    newmerchant.MerchantPostalCode,
		MerchantCategoryId:    newmerchant.MerchantCategoryId,
		MerchantWebsite:       newmerchant.MerchantWebsite,
		MerchantMediaSocialId: newmerchant.MerchantMediaSocialId,
		MerchantDescription:   newmerchant.MerchantDescription,
		MerchantImageProfile:  newmerchant.MerchantImageProfile,
		MerchantGallery:       newmerchant.MerchantGallery,
	}
	db.Model(&merchant).Where("merchant_email = ?", merchant.MerchantEmail).Update("active", false)
	db.Model(&merchant).Where("merchant_email = ?", merchant.MerchantEmail).Update("is_deleted", true)
	merchant.Deleted = time.Now()
	defer db.Close()
	return "berhasil dihapus"
}

func GetMerchant(page *int, size *int, sort *int, email *string) []model.Merchant {
	fmt.Println("masuk ke Fungsi Get")
	db := database.ConnectionDB()
	var merchant []model.Merchant
	var rows *sql.Rows
	var err error
	var total int

	if (page == nil && size == nil && sort == nil && email == nil) {
		rows, err = db.Find(&merchant).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if (page != nil && size != nil && sort != nil && email == nil) {
		switch *sort {
		case 1:
			rows, err = db.Find(&merchant).Order("merchant_name asc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		case 2:
			rows, err = db.Find(&merchant).Order("merchant_name desc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if (page != nil && size != nil && sort != nil && email != nil) {
		switch *sort {
		case 1:
			rows, err = db.Find(&merchant).Where("merchant_email = ?", email).Order("merchant_name asc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		case 2:
			rows, err = db.Find(&merchant).Where("merchant_email = ?", email).Order("merchant_name asc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if (page != nil && size != nil && sort != nil && email == nil) {
		switch *sort {
		case 1:
			rows, err = db.Find(&merchant).Order("merchant_name asc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		case 2:
			rows, err = db.Find(&merchant).Order("merchant_name desc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if page != nil && size != nil {
		rows, err = db.Find(&merchant).Count(total).Limit(*size).Offset(*page).Order("merchant_name asc").Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if (page == nil && size == nil && sort == nil && email != nil) {
		rows, err = db.Find(&merchant).Where("merchant_email = ?", email).Order("merchant_name asc").Count(total).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	result := make([]model.Merchant, 0)
	fmt.Println("lewat")
	for rows.Next() {
		fmt.Println("masuk")
		m := &model.Merchant{}
		err = rows.Scan(
			&m.Id,
			&m.Created,
			&m.CreatedBy,
			&m.Modified,
			&m.ModifiedBy,
			&m.Active,
			&m.IsDeleted,
			&m.Deleted,
			&m.Deleted_by,
			&m.MerchantName,
			&m.MerchantEmail,
			&m.MerchantPhoneNumber,
			&m.MerchantProvince,
			&m.MerchantCity,
			&m.MerchantAddress,
			&m.MerchantPostalCode,
			&m.MerchantCategoryId,
			&m.MerchantWebsite,
			&m.MerchantMediaSocialId,
			&m.MerchantDescription,
			&m.MerchantImageProfile,
			&m.MerchantGallery,
			&m.CategoryName,
			&m.Rating,
		)

		category := new(model.MerchantCategory)
		db.Table("merchant_categories").Select("merchant_categories.category_name").
			Where("id = ? ", m.MerchantCategoryId).First(&category)
		m.CategoryName = category.CategoryName

		profile := new(model.ImageProfile)
		db.Table("image_profiles").Select("image_profiles.link").
			Where("merchant_email = ?", m.MerchantEmail).First(&profile)
		m.MerchantImageProfile = profile.Link

		var gall string
		var gallery []model.Gallery
		db.Table("galleries").Select("galleries.link").
			Where("merchant_email = ?", m.MerchantEmail).Find(&gallery)
		fmt.Println("Galeri : ", gallery)
		for _, val := range gallery {
			gall = gall + " , " + val.Link
		}
		fmt.Println("Gall : ", gall)
		m.MerchantGallery = gall

		//var rev int
		var review []model.Review
		var rate float32
		var total float32
		db.Table("reviews").Select("reviews.rating").
			Where("merchant_email = ? ", m.MerchantEmail).Find(&review)
		fmt.Println("Review : ", review)
		for _, value  :=range review {
			rate = rate + value.Rating
			fmt.Println("value : ", value.Rating)
			fmt.Println("rates : ", rate)
			leng := float32(len(review))
			total = rate / leng
		}
		totals := fmt.Sprintf("%.1f", total)
		rates ,err := strconv.ParseFloat(totals, 64)
		ratess := float32(rates)
		if err != nil {
			log.Fatal(err)
		}
		m.Rating = ratess
		fmt.Println("total : ", m.Rating)
		fmt.Println("Rate : ", rate)

		if err != nil {
			log.Fatal(err)
		}
		result = append(result, *m)
	}

	defer db.Close()
	return result
}
