package repository

import (
	"database/sql"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"
)

func CreateReview(review *model.Review) error {
	db := database.ConnectionDB()
	rev := model.Review{
		Created:       time.Now(),
		CreatedBy:     "Admin",
		Modified:      time.Now(),
		ModifiedBy:    "Admin",
		Active:        true,
		IsDeleted:     false,
		Deleted:       time.Now(),
		Deleted_by:    review.Deleted_by,
		CustomerName:  review.CustomerName,
		MerchantEmail: review.MerchantEmail,
		ProgramName:   review.ProgramName,
		Review:        review.Review,
		Rating:        review.Rating,
		MerchantName:  "a",
	}
	db.Create(&rev)
	db.Close()
	return nil
}

func UpdateReview(review *model.Review) error {
	db := database.ConnectionDB()
	db.Model(&review).Updates(map[string]interface{}{
		"customer_name":  review.CustomerName,
		"merchant_email": review.MerchantEmail,
		"program_name":   review.ProgramName,
		"review":         review.Review,
		"rating":         review.Review,
	})
	review.Modified = time.Now()
	db.Close()
	return nil
}

func DeleteReview(review *model.Review) error {
	db := database.ConnectionDB()
	db.Model(&review).Where("id = ?", review.Id).Update("active", false)
	db.Model(&review).Where("id = ?", review.Id).Update("is_deleted", true)
	review.Deleted = time.Now()
	defer db.Close()
	return nil
}

func GetReview(page *int, size *int, sort *int, program_name *string) []model.Review {
	fmt.Println("masuk ke get review")
	db := database.ConnectionDB()
	var review []model.Review
	var rows *sql.Rows
	var err error
	var total int

	if page == nil && size == nil && program_name == nil && sort == nil {
		rows, err = db.Find(&review).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil && sort != nil && program_name != nil {
		switch *sort {
		case 1:
			rows, err = db.Find(&review).Where("program_name = ? ", program_name).Order("rating desc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		case 2:
			rows, err = db.Find(&review).Where("program_name = ? ", program_name).Order("rating asc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		case 3:
			rows, err = db.Find(&review).Where("program_name = ? ", program_name).Order("program_name desc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		case 4:
			rows, err = db.Find(&review).Where("program_name = ? ", program_name).Order("program_name asc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if page != nil && size != nil && sort != nil && program_name == nil {
		switch *sort {
		case 1:
			rows, err = db.Find(&review).Order("rating desc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		case 2:
			rows, err = db.Find(&review).Order("rating asc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		case 3:
			rows, err = db.Find(&review).Order("program_name desc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		case 4:
			rows, err = db.Find(&review).Order("program_name asc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if page != nil  && size != nil && sort == nil && program_name == nil {
		rows, err = db.Find(&review).Order("id asc").Count(total).Limit(*size).Offset(*page).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page == nil  && size == nil && sort == nil && program_name != nil {
		rows, err = db.Find(&review).Where("program_name = ?", program_name).Order("id asc").Count(total).Limit(*size).Offset(*page).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	result := make([]model.Review, 0)

	for rows.Next() {
		r := &model.Review{}

		err = rows.Scan(
			&r.Id,
			&r.Created,
			&r.CreatedBy,
			&r.ModifiedBy,
			&r.ModifiedBy,
			&r.Active,
			&r.IsDeleted,
			&r.Deleted,
			&r.Deleted_by,
			&r.CustomerName,
			&r.MerchantEmail,
			&r.ProgramName,
			&r.Review,
			&r.Rating,
			&r.MerchantName,
		)

		merchant := new(model.Merchant)
		db.Table("merchants").
			Select("merchants.merchant_name").
			Where("merchant_email = ?", r.MerchantEmail).
			First(&merchant)
		r.MerchantName = merchant.MerchantName

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, *r)
	}

	db.Close()
	return result
}
