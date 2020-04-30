package main

import (
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

func main() {
	db := database.ConnectionDB()

	db.AutoMigrate(&model.Voucher{})

	//db.AutoMigrate(&model.Reward{})
	//reward := model.Reward{
	//	Created:           time.Now(),
	//	CreatedBy:         "Admin",
	//	Modified:          time.Now(),
	//	ModifiedBy:        "Admin",
	//	Active:            true,
	//	IsDeleted:         false,
	//	Deleted:           nil,
	//	DeletedBy:         "",
	//	RedeemPoints:      100,
	//	RewardName:        "Adidas+ Reward",
	//	RedeemRules:       "Bawa ke kasir",
	//	TermsAndCondition: "Tukarkan reward ini dengan kaos kaki gratis dari kami",
	//	ProgramId:         2,
	//	MerchantEmail:     "contact@adidas.com",
	//	Outletid:          2,
	//}
	//db.Create(&reward)

	//db.AutoMigrate(&model.Voucher{})
	//voucher := model.Voucher{
	//	Created:                  time.Now(),
	//	CreatedBy:                "Admin",
	//	Modified:                 time.Now(),
	//	ModifiedBy:               "Admin",
	//	Active:                   true,
	//	IsDeleted:                false,
	//	Deleted:                  nil,
	//	DeletedBy:                "",
	//	VoucherName:              "Voucher Adidas+",
	//	StartDate:                time.Date(2020, time.May, 1, 00,00,00,00000, time.UTC),
	//	EndDate:                  time.Date(2020, time.May, 23, 59,59,59,99999, time.UTC),
	//	VoucherDescription:       "Voucher untuk Adidas+",
	//	VoucherTermsAndCondition: "Belanja dulu, kasih ke kasir selanjutnya",
	//	IsPushNotification:       true,
	//	IsGiveVoucher:            true,
	//	VoucherPeriod:            "14 days",
	//	RewardTermsAndCondition:  "kasih ke kasir saja",
	//	BackgroundVoucherPattern: "asdada",
	//	BackgroundVoucherColour:  "#346656",
	//	MerchantEmail:            "contact@adidas.com",
	//	OutletId:                 2,
	//	ProgramId:                2,
	//}
	//db.Create(&voucher)
	//
	//db.AutoMigrate(&model.ProgramCard{})
	//card := model.ProgramCard{
	//	Id:                2,
	//	Created:           time.Now(),
	//	CreatedBy:         "Admin",
	//	Modified:          time.Now(),
	//	ModifiedBy:        "Admin",
	//	Active:            true,
	//	IsDeleted:         false,
	//	Deleted:           nil,
	//	DeletedBy:         "",
	//	Title:             "Adidas+",
	//	Description:       "Untuk program Adidas+",
	//	FontColor:         "#632366",
	//	TemplateColor:     "#ffffff",
	//	IconImage:         "asda",
	//	TermsAndCondition: "belanja banyak makin miskin",
	//	Benefit:           "Tidak Ada",
	//	ValidUntil:        time.Date(2020, time.May, 15, 23, 59, 59, 99999, time.UTC),
	//	CurrentPoint:      0,
	//	IsValid:           true,
	//	ProgramId:         2,
	//	CardType:          "Member",
	//	IconImageStamp:    "asdaas",
	//	MerchantEmail:     "contact@adidas.com",
	//	Tier:              "Platinum",
	//	TemplatePattern:   "asdaaf",
	//}
	//db.Create(&card)

	//db.AutoMigrate(&model.TransactionMerchant{})
	//trans := model.TransactionMerchant{
	//	Created:          time.Now(),
	//	CreatedBy:        "Admin",
	//	Modified:         time.Now(),
	//	ModifiedBy:       "Admin",
	//	Active:           true,
	//	IsDeleted:        false,
	//	Deleted:          nil,
	//	Deleted_by:       "",
	//	MerchantEmail:    "contact@adidas.com",
	//	OutletId:         2,
	//	TotalTransaction: 15000000,
	//	PointTransaction: 0,
	//	BillNumber:       "8880880",
	//}
	//db.Create(&trans)
	//
	////db.AutoMigrate(&model.Employee{})
	//employee := model.Employee{
	//	Created:       time.Now(),
	//	CreatedBy:     "Admin",
	//	Modified:      time.Now(),
	//	ModifiedBy:    "Admin",
	//	Active:        true,
	//	IsDeleted:     false,
	//	Deleted:       nil,
	//	Deleted_by:    "",
	//	EmployeeName:  "bambang",
	//	EmployeeEmail: "bambang@gmail.com",
	//	EmployeePin:   "4321",
	//	EmployeeRole:  "Kasir",
	//	OutletId:      2,
	//}
	//db.Create(&employee)
	//db.AutoMigrate(&model.SpecialProgram{})
	//special := model.SpecialProgram{
	//	Created:               time.Now(),
	//	CreatedBy:             "Admin",
	//	Modified:              time.Now(),
	//	ModifiedBy:            "Admin",
	//	Active:                true,
	//	IsDeleted:             false,
	//	Deleted:               nil,
	//	Deleted_by:            "",
	//	ProgramName:           "Adidas++",
	//	ProgramImage:          "asdaf",
	//	ProgramStartDate:      time.Date(2020, time.May, 1, 00, 00,00,000000, time.UTC),
	//	ProgramEndDate:        time.Date(2020, time.May, 15, 23, 59,59,999999, time.UTC),
	//	ProgramDescription:    "Beli banyak makin miskin",
	//	Card:                  "Member",
	//	OutletID:              2,
	//	MerchantEmail:         "contact@adidas.com",
	//	CategoryId:            1,
	//	Benefit:               "Tidak Ada",
	//	TermsAndCondition:     "Beli minimum 45 juta",
	//	Tier:                  "Platinum",
	//	RedeemRules:           "Belanja Mantap",
	//	RewardTarget:          0,
	//	QRCodeId:              "",
	//	IsReqBillNumber:       true,
	//	IsReqTotalTransaction: true,
	//	IsPushNotification:    true,
	//	IsLendCard:            true,
	//	IsGiveCard:            true,
	//	IsWelcomeBonus:        true,
	//}
	//db.Create(&special)

	//db.AutoMigrate(&model.Employee{}, &model.MerchantCategory{}, &model.Merchant{}, &model.TransactionMerchant{},
	//		&model.Outlet{}, &model.Program{}, &model.SpecialProgram{}, &model.Reward{}, &model.MerchantSocialMedia{},
	//		&model.Voucher{}, &model.ProgramCard{}, &model.CardType{} )
	//db.AutoMigrate(&model.Merchant{})
	//merchant := model.Merchant{
	//	Created:               time.Now(),
	//	CreatedBy:             "Admin",
	//	Modified:              time.Now(),
	//	ModifiedBy:            "Admin",
	//	Active:                true,
	//	IsDeleted:             false,
	//	Deleted:               nil,
	//	Deleted_by:            "",
	//	MerchantName:          "Adidas",
	//	MerchantEmail:         "contact@adidas.com",
	//	MerchantPhoneNumber:   "4567890",
	//	MerchantProvince:      "DKI Jakarta",
	//	MerchantCity:          "Jakarta",
	//	MerchantAddress:       "Jalan Sudirman",
	//	MerchantPostalCode:    "12099",
	//	MerchantCategoryId:    1,
	//	MerchantWebsite:       "www.AdidasIndonesia.com",
	//	MerchantMediaSocialId: 1,
	//	MerchantDescription:   "Jualan sepatu mantap garis 3",
	//	MerchantImageProfile:  "abcde",
	//	MerchantGallery:  "abcdefghijk"    ,
	//}
	//db.Create(&merchant)

	//db.AutoMigrate(&model.Outlet{})
	//outlet := model.Outlet{
	//	Created:          time.Now(),
	//	CreatedBy:        "Admin",
	//	Modified:         time.Now(),
	//	ModifiedBy:       "Admin",
	//	Active:           true,
	//	IsDeleted:        false,
	//	Deleted:          nil,
	//	Deleted_by:       "",
	//	OutletName:       "Adidas Kota Kasablanka",
	//	OutletAddress:    "Kota Kasablanka",
	//	OutletPhone:      "(061)44423466",
	//	OutletCity:       "Jakarta",
	//	OutletProvince:   "DKI Jakarta",
	//	OutletPostalCode: "122920",
	//	OutletLongitude:  "",
	//	OutletLatitude:   "",
	//	OutletDay:        time.Time{},
	//	OutletHour:       time.Time{},
	//	MerchantEmail:    "contact@adidas.com",
	//	Timezone:         "WIB",
	//}
	//db.Create(&outlet)

	//category := model.MerchantCategory{
	//	Id:1,
	//	Created:      time.Now(),
	//	CreatedBy:    "Admin",
	//	Modified:     time.Now(),
	//	ModifiedBy:   "Admin",
	//	Active:       true,
	//	IsDeleted:    false,
	//	Deleted:      nil,
	//	Deleted_by:   "",
	//	CategoryName: "Fashion",
	//	ImageUrl:     "",
	//}
	//db.Create(&category)
	//program := model.Program{
	//	Created:               time.Now(),
	//	CreatedBy:             "Admin",
	//	Modified:              time.Now(),
	//	ModifiedBy:            "Admin",
	//	Active:                true,
	//	IsDeleted:             false,
	//	Deleted:               nil,
	//	Deleted_by:            "",
	//	ProgramName:           "Adidas+",
	//	ProgramImage:          "astghas",
	//	ProgramStartDate:      time.Date(2020,time.May, 1,00,00,00,00000,time.UTC),
	//	ProgramEndDate:        time.Date(2020,time.May, 31,23,59,59,99999,time.UTC),
	//	ProgramDescription:    "Beli 1 ga dapat apa-apa",
	//	Card:                  "Member",
	//	OutletID:              "2",
	//	MerchantEmail:         "contact@adidas.com",
	//	CategoryId:            1,
	//	Benefit:               "Ga ada",
	//	TermsAndCondition:     "Belanja banyak",
	//	Tier:                  "Platinum",
	//	RedeemRules:           "Bayar 10 juta",
	//	RewardTarget:          0,
	//	QRCodeId:              "",
	//	ProgramPoint:          0,
	//	MinPayment:            0,
	//	IsReqBillNumber:       true,
	//	IsReqTotalTransaction: true,
	//	IsPushNotification:    true,
	//	IsLendCard:            true,
	//	IsGiveCard:            true,
	//	IsWelcomeBonus:        true,
	//}
	//db.Create(&program)
}
