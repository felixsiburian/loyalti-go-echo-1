package repository

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"
)

const (
	MemberTypeCard = "Member"
	GoldTier       = "Gold"
	SilverTier     = "Silver"
	PlatinumTier   = "Platinum"
)

type MemberTier string

const (
	Silver   MemberTier = "1"
	Gold     MemberTier = "2"
	Platinum MemberTier = "3"
)

type CardRepository interface {
	CreateCard(newcard *model.ProgramCard) error
	DeleteCard(newcard *model.ProgramCard) error
	UpdateCard(newcard *model.ProgramCard) error
}

type card_repo struct {
	DB *gorm.DB
}

func (p *card_repo) CreateCard(newcard *model.ProgramCard) error {
	kartu := model.ProgramCard{
		Created:           time.Now(),
		CreatedBy:         "System",
		Modified:          time.Now(),
		ModifiedBy:        "System",
		Active:            true,
		IsDeleted:         false,
		Deleted:           time.Now(),
		DeletedBy:         "",
		Title:             newcard.Title,
		Description:       newcard.Description,
		FontColor:         newcard.FontColor,
		TemplateColor:     newcard.TemplateColor,
		IconImage:         newcard.IconImage,
		TermsAndCondition: newcard.TermsAndCondition,
		Benefit:           newcard.Benefit,
		ValidUntil:        newcard.ValidUntil,
		CurrentPoint:      newcard.CurrentPoint,
		IsValid:           true,
		ProgramId:         newcard.ProgramId,
		CardType:          newcard.CardType,
		IconImageStamp:    newcard.IconImageStamp,
		MerchantEmail:     newcard.MerchantEmail,
		Tier:              "",
	}
	db := database.ConnectionDB()
	err := db.Create(&kartu).Error
	if err == nil {
		fmt.Println("Error")
	}
	return err
}

func CreateCardRepository(db *gorm.DB) CardRepository {
	return &card_repo{
		DB: db,
	}
}

func (p *card_repo) DeleteCard(newcard *model.ProgramCard) error {
	db := database.ConnectionDB()
	err := db.Model(&newcard).Where("id = ?", newcard.Id).Update("active", false).Error
	if err == nil {
		fmt.Println("tidak ada error")
	}
	return err
}

func (p *card_repo) UpdateCard(newcard *model.ProgramCard) error {
	db := database.ConnectionDB()
	err := db.Model(&newcard).Where("id = ?", newcard.Id).Update(&newcard).Error

	return err
}

//func CreateCardMerchant(card *model.ProgramCard) error {
//	db := database.ConnectionDB()
//
//	for i := 0; i <= 2; i++ {
//		cards := model.ProgramCard{
//			Created:           time.Now(),
//			CreatedBy:         "",
//			Modified:          time.Now(),
//			ModifiedBy:        "",
//			Active:            true,
//			IsDeleted:         false,
//			Deleted:           nil,
//			DeletedBy:         "",
//			Title:             card.Title,
//			Description:       card.Description,
//			FontColor:         card.FontColor,
//			TemplateColor:     card.TemplateColor,
//			IconImage:         card.IconImage,
//			TermsAndCondition: card.TermsAndCondition,
//			Benefit:           card.Benefit,
//			ValidUntil:        time.Now(),
//			CurrentPoint:      card.CurrentPoint,
//			IsValid:           card.IsValid,
//			ProgramId:         card.ProgramId,
//			CardType:          card.CardType,
//			IconImageStamp:    card.IconImageStamp,
//			MerchantEmail:        card.MerchantEmail,
//		}
//		if (i == 0) {
//			fmt.Println("masuk ke if == 0")
//			fmt.Println("isi enum silver", domain.EnumMember.Silver)
//			cards.Tier = domain.EnumMember.Silver
//			db.Create(&cards)
//		} else if (i == 1) {
//			fmt.Println("masuk ke if == 1")
//			fmt.Println("isi enum silver", domain.EnumMember.Gold)
//			cards.Tier = domain.EnumMember.Gold
//			db.Create(&cards)
//		} else {
//			fmt.Println("masuk ke if == 2")
//			fmt.Println("isi enum silver", domain.EnumMember.Platinum)
//			cards.Tier = domain.EnumMember.Platinum
//			db.Create(&cards)
//		}
//	}
//	return nil
//}

func CreateCardMerchant(card *model.ProgramCard) string {
	db := database.ConnectionDB()
	cards := model.ProgramCard{
		Created:           time.Now(),
		CreatedBy:         "Admin",
		Modified:          time.Now(),
		ModifiedBy:        "Admin",
		Active:            true,
		IsDeleted:         false,
		Deleted:           time.Now(),
		DeletedBy:         "",
		Title:             card.Title,
		Description:       card.Description,
		FontColor:         card.FontColor,
		TemplateColor:     card.TemplateColor,
		IconImage:         card.IconImage,
		TermsAndCondition: card.TermsAndCondition,
		Benefit:           card.Benefit,
		ValidUntil:        card.ValidUntil,
		CurrentPoint:      card.CurrentPoint,
		IsValid:           card.IsValid,
		ProgramId:         card.ProgramId,
		CardType:          card.CardType,
		IconImageStamp:    card.IconImageStamp,
		MerchantEmail:     card.MerchantEmail,
		ProgramName:       "a",
		MerchantName:      "b",
	}
	if (cards.CardType != "Member") {
		fmt.Println("masuk kesini")
		db.Create(&cards)
	}

	if cards.CardType == "Member" {
		for i := 0; i <= 2; i++ {
			cards2 := model.ProgramCard{
				Created:           time.Now(),
				CreatedBy:         "Admin",
				Modified:          time.Now(),
				ModifiedBy:        "Admin",
				Active:            true,
				IsDeleted:         false,
				Deleted:           time.Now(),
				DeletedBy:         "",
				Title:             card.Title,
				Description:       card.Description,
				FontColor:         card.FontColor,
				TemplateColor:     card.TemplateColor,
				IconImage:         card.IconImage,
				TermsAndCondition: card.TermsAndCondition,
				Benefit:           card.Benefit,
				ValidUntil:        card.ValidUntil,
				CurrentPoint:      card.CurrentPoint,
				IsValid:           card.IsValid,
				ProgramId:         card.ProgramId,
				CardType:          card.CardType,
				IconImageStamp:    card.IconImageStamp,
				MerchantEmail:     card.MerchantEmail,
				ProgramName:       "a",
				MerchantName:      "b",
			}


			if (i == 0) {
				fmt.Println("masuk ke if == 0")
				fmt.Println("isi enum silver", domain.EnumMember.Silver)
				cards2.Tier = domain.EnumMember.Silver
				db.Create(&cards2)
			} else if (i == 1) {
				fmt.Println("masuk ke if == 1")
				fmt.Println("isi enum silver", domain.EnumMember.Gold)
				cards2.Tier = domain.EnumMember.Gold
				db.Create(&cards2)
			} else {
				fmt.Println("masuk ke if == 2")
				fmt.Println("isi enum silver", domain.EnumMember.Platinum)
				cards2.Tier = domain.EnumMember.Platinum
				db.Create(&cards2)
			}
		}
	}
	//db.Create(&cards)
	db.Close()
	fmt.Println("Selesai")
	return cards.Description
}

func UpdateCardMerchant(card *model.ProgramCard) error {
	fmt.Println("masuk ke update")
	db := database.ConnectionDB()
	card.Modified = time.Now()
	db.Model(&card).Updates(map[string]interface{}{
		"title": card.Title,
		"description": card.Description,
		"font_color": card.FontColor,
		"template_color": card.TemplateColor,
		"icon_image": card.IconImage,
		"terms_and_condition": card.TermsAndCondition,
		"benefit": card.Benefit,
		"valid_until": card.ValidUntil,
		"current_point":card.CurrentPoint,
		"is_valid": card.IsValid,
		"program_id": card.ProgramId,
		"card_type" : card.CardType,
		"icon_image_stamp": card.IconImageStamp,
		"merchant_email": card.MerchantEmail,
		"tier": card.Tier,
		"template_pattern": card.TemplatePattern,
	})

	//db.Model(&card).Update(&card)
	db.Close()
	return nil
}

func DeleteCardMerchant(card *model.ProgramCard) error {
	fmt.Println("masuk ke delete")
	db := database.ConnectionDB()
	card.Deleted = time.Now()
	db.Model(&card).Where("id = ?", card.Id).Update("active", false)
	db.Model(&card).Where("id = ?", card.Id).Update("is_deleted", true)

	//if err != nil {
	//	fmt.Println("error : ", err.Error)
	//	log.Fatal(err)
	//}
	db.Close()
	return nil
}

func GetCardMerchant(page *int, size *int, id *int, card_type *string) []model.ProgramCard {

	db := database.ConnectionDB()
	var kartu []model.ProgramCard
	var rows *sql.Rows
	var err error
	var total int

	if page == nil && size == nil && id == nil && card_type == nil {
		rows, err = db.Find(&kartu).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil && id != nil && card_type != nil {
		rows, err = db.Find(&kartu).Where("card_type = ? ", card_type).Order("title asc").Count(total).Limit(*size).Offset(*page).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil {
		rows, err = db.Find(&kartu).Order("title asc").Count(total).Limit(*size).Offset(*page).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	result := make([]model.ProgramCard, 0)

	for rows.Next() {
		o := new(model.ProgramCard)
		var err = rows.Scan(
			&o.Id,
			&o.Created,
			&o.CreatedBy,
			&o.Modified,
			&o.ModifiedBy,
			&o.Active,
			&o.IsDeleted,
			&o.Deleted,
			&o.DeletedBy,
			&o.Title,
			&o.Description,
			&o.FontColor,
			&o.TemplateColor,
			&o.IconImage,
			&o.TermsAndCondition,
			&o.Benefit,
			&o.ValidUntil,
			&o.CurrentPoint,
			&o.IsValid,
			&o.ProgramId,
			&o.CardType,
			&o.IconImageStamp,
			&o.MerchantEmail,
			&o.Tier,
			&o.TemplatePattern,
			&o.ProgramName,
			&o.MerchantName,
		)
		//add tier
		//if o.CardType == MemberTypeCard{
		//	if o.CurrentPoint >= 0 && o.CurrentPoint <= 100{
		//		o.Tier = SilverTier
		//	}else if o.CurrentPoint > 100 && o.CurrentPoint <= 230{
		//		o.Tier = GoldTier
		//	}else if o.CurrentPoint > 230 && o.CurrentPoint <= 400{
		//		o.Tier = PlatinumTier
		//	}
		//}else{
		//	o.Tier = "Ga ada tier"
		//}

		merchant := new(model.Merchant)
		db.Table("merchants").
			Select("merchants.merchant_name").
			Where("merchant_email = ?", o.MerchantEmail).
			First(&merchant)
		o.MerchantName = merchant.MerchantName

		fmt.Println("temp : ", o.MerchantName)

		program := new(model.Program)
		db.Table("programs").
			Select("programs.program_name").
			Where("id = ?", o.ProgramId).
			First(&program)
		o.ProgramName = program.ProgramName

		fmt.Println("temp2 : ", o.ProgramName)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, *o)
	}

	db.Close()
	return result
}

func GetCardMember(program_id int) []model.ProgramCard {
	db := database.ConnectionDB()
	card := &model.ProgramCard{}
	db.Model("card").Where("program_id = ?", program_id).Order("tier asc").First(&card)

	tier := make([]model.ProgramCard, 0)
	tier = append(tier, *card)

	fmt.Println("tier : ", tier)

	return tier
}
