package model

import (
	"bytes"
	"fmt"
	"html/template"
	"time"
)

type Song struct {
	ID       string `json:"id,omitempty"`
	Album    string `json:"album"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
	Type     string `json:"type"`
}

type AccountMerchant struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ResponseProcess struct {
	Sub        string `json:"sub"`
	Upn        string `json:"upn"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	UserId     string `json:"user_id"`
}

type Response struct {
	Access_Token string `json:"access_token"`
	Expires_In   int    `json:"expires_in"`
	Token_Type   string `json:"token_type"`
}

type User struct {
	Username string
	Password string
}

type NewMerchantTest struct {
	Id            int    `json:"id"`
	MerchantEmail string `json:"merchant_email"`
}

type NewMerchantCommand struct {
	Id                    int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Created               time.Time `json:"created"`
	CreatedBy             string    `json:"created_by"`
	Modified              time.Time `json:"modified"`
	ModifiedBy            string    `json:"modified_by"`
	Active                bool      `json:"active"`
	IsDeleted             bool      `json:"is_deleted"`
	Deleted               time.Time `json:"deleted"`
	Deleted_by            string    `json:"deleted_by"`
	MerchantName          string    `json:"merchant_name"`
	MerchantEmail         string    `json:"merchant_email"`
	MerchantPhoneNumber   string    `json:"merchant_phone_number"`
	MerchantProvince      string    `json:"merchant_province"`
	MerchantCity          string    `json:"merchant_city"`
	MerchantAddress       string    `json:"merchant_address"`
	MerchantPostalCode    string    `json:"merchant_postal_code"`
	MerchantCategoryId    int       `json:"merchant_category_id"`
	MerchantWebsite       string    `json:"merchant_website"`
	MerchantMediaSocialId int       `json:"merchant_media_social_id"`
	MerchantDescription   string    `json:"merchant_description"`
	MerchantImageProfile  string    `json:"merchant_image_profile"`
	MerchantGallery       string    `json:"merchant_gallery"`
	//MerchantPassword      string     `json:"merchant_password"`

	CategoryName string `json:"category_name"`
}

type Merchant struct {
	//GeneralModels
	Id                    int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Created               time.Time `json:"created"`
	CreatedBy             string    `json:"created_by"`
	Modified              time.Time `json:"modified"`
	ModifiedBy            string    `json:"modified_by"`
	Active                bool      `json:"active"`
	IsDeleted             bool      `json:"is_deleted"`
	Deleted               time.Time `json:"deleted"`
	Deleted_by            string    `json:"deleted_by"`
	MerchantName          string    `json:"merchant_name"`
	MerchantEmail         string    `json:"merchant_email"`
	MerchantPhoneNumber   string    `json:"merchant_phone_number"`
	MerchantProvince      string    `json:"merchant_province"`
	MerchantCity          string    `json:"merchant_city"`
	MerchantAddress       string    `json:"merchant_address"`
	MerchantPostalCode    string    `json:"merchant_postal_code"`
	MerchantCategoryId    int       `json:"merchant_category_id"`
	MerchantWebsite       string    `json:"merchant_website"`
	MerchantMediaSocialId int       `json:"merchant_media_social_id"`
	MerchantDescription   string    `json:"merchant_description"`
	MerchantImageProfile  string    `json:"merchant_image_profile"`
	MerchantGallery       string    `json:"merchant_gallery"`

	CategoryName string `json:"category_name"`
}

type Gallery struct {
	Id             int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Merchant_email string `json:"merchant_email"`
	Link           string `json:"link"`
}

type MerchantCategory struct {
	Id           int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Created      time.Time `json:"created"`
	CreatedBy    string    `json:"created_by"`
	Modified     time.Time `json:"modified"`
	ModifiedBy   string    `json:"modified_by"`
	Active       bool      `json:"active"`
	IsDeleted    bool      `json:"is_deleted"`
	Deleted      time.Time `json:"deleted"`
	Deleted_by   string    `json:"deleted_by"`
	CategoryName string    `json:"category_name"`
	ImageUrl     string    `json:"image_url"`
}

type MerchantSocialMedia struct {
	Id                  int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Created             time.Time `json:"created"`
	CreatedBy           string    `json:"created_by"`
	Modified            time.Time `json:"modified"`
	ModifiedBy          string    `json:"modified_by"`
	Active              bool      `json:"active"`
	IsDeleted           bool      `json:"is_deleted"`
	Deleted             time.Time `json:"deleted"`
	Deleted_by          string    `json:"deleted_by"`
	SocialMediaName     string    `json:"social_media_name"`
	SocialMediaImageUrl string    `json:"social_media_image_url"`
}

type CardType struct {
	Id           int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Created      time.Time `json:"created"`
	CreatedBy    string    `json:"created_by"`
	Modified     time.Time `json:"modified"`
	ModifiedBy   string    `json:"modified_by"`
	Active       bool      `json:"active"`
	IsDeleted    bool      `json:"is_deleted"`
	Deleted      time.Time `json:"deleted"`
	Deleted_by   string    `json:"deleted_by"`
	CardTypeName string    `json:"card_type_name"`
}

type Outlet2 struct {
	Id               int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Created          time.Time `json:"created"`
	CreatedBy        string    `json:"created_by"`
	Modified         time.Time `json:"modified"`
	ModifiedBy       string    `json:"modified_by"`
	Active           bool      `json:"active"`
	IsDeleted        bool      `json:"is_deleted"`
	Deleted          time.Time `json:"deleted"`
	Deleted_by       string    `json:"deleted_by"`
	OutletName       string    `json:"outlet_name"`
	OutletAddress    string    `json:"outlet_address"`
	OutletPhone      string    `json:"outlet_phone"`
	OutletCity       string    `json:"outlet_city"`
	OutletProvince   string    `json:"outlet_province"`
	OutletPostalCode string    `json:"outlet_postal_code"`
	OutletLongitude  string    `json:"outlet_longitude"`
	OutletLatitude   string    `json:"outlet_latitude"`
	OutletDay        string    `json:"outlet_day"`
	OutletHour       string    `json:"outlet_hour"`
	MerchantEmail    string    `json:"merchant_email"`
	Timezone         string    `json:"timezone"`

	MerchantName string `json:"merchant_name"`
}

type Program struct {
	Id                    int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"`
	Created               time.Time `json:"created"`
	CreatedBy             string    `json:"created_by"`
	Modified              time.Time `json:"modified"`
	ModifiedBy            string    `json:"modified_by"`
	Active                bool      `json:"active"`
	IsDeleted             bool      `json:"is_deleted"`
	Deleted               time.Time `json:"deleted"`
	Deleted_by            string    `json:"deleted_by"`
	ProgramName           string    `json:"program_name"`
	ProgramImage          string    `json:"program_image"`
	ProgramStartDate      string    `json:"program_start_date"`
	ProgramEndDate        string    `json:"program_end_date"`
	ProgramDescription    string    `json:"program_description"`
	Card                  string    `json:"card"`
	OutletID              int       `json:"outlet_id"`
	MerchantEmail         string    `json:"merchant_email"`
	CategoryId            int       `json:"category_id"`
	Benefit               string    `json:"benefit"`
	TermsAndCondition     string    `json:"terms_and_condition"`
	Tier                  string    `json:"tier"`
	RedeemRules           string    `json:"redeem_rules"`
	RewardTarget          int       `json:"reward_target"`
	QRCodeId              string    `json:"qr_code_id"`
	ProgramPoint          int       `json:"program_point"`
	MinPayment            int       `json:"min_payment"`
	IsReqBillNumber       bool      `json:"is_req_bill_number"`
	IsReqTotalTransaction bool      `json:"is_req_total_transaction"`
	IsPushNotification    bool      `json:"is_push_notification"`
	IsLendCard            bool      `json:"is_lend_card"`
	IsGiveCard            bool      `json:"is_give_card"`
	IsWelcomeBonus        bool      `json:"is_welcome_bonus"`

	MerchantName string `json:"merchant_name"`
	OutletName   string `json:"outlet_name"`
	CategoryName string `json:"category_name"`
}

type SpecialProgram struct {
	Id                    int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Created               time.Time `json:"created"`
	CreatedBy             string    `json:"created_by"`
	Modified              time.Time `json:"modified"`
	ModifiedBy            string    `json:"modified_by"`
	Active                bool      `json:"active"`
	IsDeleted             bool      `json:"is_deleted"`
	Deleted               time.Time `json:"deleted"`
	Deleted_by            string    `json:"deleted_by"`
	ProgramName           string    `json:"program_name"`
	ProgramImage          string    `json:"program_image"`
	ProgramStartDate      string    `json:"program_start_date"`
	ProgramEndDate        string    `json:"program_end_date"`
	ProgramDescription    string    `json:"program_description"`
	Card                  string    `json:"card"`
	OutletID              int       `json:"outlet_id"`
	MerchantEmail         string    `json:"merchant_email"`
	CategoryId            int       `json:"category_id"`
	Benefit               string    `json:"benefit"`
	TermsAndCondition     string    `json:"terms_and_condition"`
	Tier                  string    `json:"tier"`
	RedeemRules           string    `json:"redeem_rules"`
	RewardTarget          float64   `json:"reward_target"`
	QRCodeId              string    `json:"qr_code_id"`
	IsReqBillNumber       bool      `json:"is_req_bill_number"`
	IsReqTotalTransaction bool      `json:"is_req_total_transaction"`
	IsPushNotification    bool      `json:"is_push_notification"`
	IsLendCard            bool      `json:"is_lend_card"`
	IsGiveCard            bool      `json:"is_give_card"`
	IsWelcomeBonus        bool      `json:"is_welcome_bonus"`

	MerchantName string `json:"merchant_name"`
	OutletName   string `json:"outlet_name"`
	CategoryName string `json:"category_name"`
}

type Product struct {
	Id                string    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"`
	Created           time.Time `json:"created"`
	CreatedBy         string    `json:"created_by"`
	Modified          time.Time `json:"modified"`
	ModifiedBy        string    `json:"modified_by"`
	Active            bool      `json:"active"`
	IsDeleted         bool      `json:"is_deleted"`
	Deleted           time.Time `json:"deleted"`
	Deleted_by        string    `json:"deleted_by"`
	ProductName       string    `json:"product_name"`
	ProductDesc       string    `json:"product_desc"`
	MerchantEmail     int       `json:"merchant_email"`
	ProductCategoryId int       `json:"product_category_id"`
}

type ProductCategory struct {
	Id                  string    `gorm:"PRIMARY_KEY;NOT NULL"; json:"id"`
	Created             time.Time `json:"created"`
	CreatedBy           string    `json:"created_by"`
	Modified            time.Time `json:"modified"`
	ModifiedBy          string    `json:"modified_by"`
	Active              bool      `json:"active"`
	IsDeleted           bool      `json:"is_deleted"`
	Deleted             time.Time `json:"deleted"`
	Deleted_by          string    `json:"deleted_by"`
	ProductCategoryDesc string    `json:"product_category_desc"`
}

type MerchantStatus struct {
	Id                 string    `gorm:"PRIMARY_KEY;NOT NULL"; json:"id"`
	Created            time.Time `json:"created"`
	CreatedBy          string    `json:"created_by"`
	Modified           time.Time `json:"modified"`
	ModifiedBy         string    `json:"modified_by"`
	Active             bool      `json:"active"`
	IsDeleted          bool      `json:"is_deleted"`
	Deleted            time.Time `json:"deleted"`
	Deleted_by         string    `json:"deleted_by"`
	MerchantStatusDesc string    `json:"merchant_status_desc"`
}

type Employee struct {
	Id            int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"`
	Created       time.Time `json:"created"`
	CreatedBy     string    `json:"created_by"`
	Modified      time.Time `json:"modified"`
	ModifiedBy    string    `json:"modified_by"`
	Active        bool      `json:"active"`
	IsDeleted     bool      `json:"is_deleted"`
	Deleted       time.Time `json:"deleted"`
	Deleted_by    string    `json:"deleted_by"`
	EmployeeName  string    `json:"employee_name"`
	EmployeeEmail string    `json:"employee_email"`
	EmployeePin   string    `json:"employee_pin"`
	EmployeeRole  string    `json:"employee_role"`
	OutletId      int       `json:"outlet_id"`

	OutletName string `json:"outlet_name"`
}

type TotalPoint struct {
	Total int `json:"total_point"`
}

type TotalChop struct {
	Total int `json:"total_chop"`
}

type TransactionMerchant struct {
	Id               int       `gorm:"PRIMARY_KEY;NOT NUll"; json:"id"`
	Created          time.Time `json:"created"`
	CreatedBy        string    `json:"created_by"`
	Modified         time.Time `json:"modified"`
	ModifiedBy       string    `json:"modified_by"`
	Active           bool      `json:"active"`
	IsDeleted        bool      `json:"is_deleted"`
	Deleted          time.Time `json:"deleted"`
	Deleted_by       string    `json:"deleted_by"`
	MerchantEmail    string    `json:"merchant_email"`
	OutletId         int       `json:"outlet_id"`
	TotalTransaction int       `json:"total_transaction"`
	PointTransaction int       `json:"point_transaction"`
	BillNumber       string    `json:"bill_number"`

	MerchantName string `json:"merchant_name"`
	OutletName   string `json:"outlet_name"`
}

type ProgramCard struct {
	Id                int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY"; json:"id"`
	Created           time.Time `json:"created"`
	CreatedBy         string    `json:"created_by"`
	Modified          time.Time `json:"modified"`
	ModifiedBy        string    `json:"modified_by"`
	Active            bool      `json:"active"`
	IsDeleted         bool      `json:"is_deleted"`
	Deleted           time.Time `json:"deleted"`
	DeletedBy         string    `json:"deleted_by"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	FontColor         string    `json:"font_color"`
	TemplateColor     string    `json:"template_color"`
	IconImage         string    `json:"icon_image"`
	TermsAndCondition string    `json:"terms_and_condition"`
	Benefit           string    `json:"benefit"`
	ValidUntil        string    `json:"valid_until"`
	CurrentPoint      int       `json:"current_point"`
	IsValid           bool      `json:"is_valid"`
	ProgramId         int       `json:"program_id"`
	CardType          string    `json:"card_type"`
	IconImageStamp    string    `json:"icon_image_stamp"`
	MerchantEmail     string    `json:"merchant_email"`
	Tier              string    `json:"tier"`
	TemplatePattern   string    `json:"template_pattern"`

	ProgramName  string `json:"program_name"`
	MerchantName string `json:"merchant_name"`
}

type Test struct {
	Name  string `json:"name"`
	Skill string `json:"skill"`
}

type Emaill struct {
	Email string `json:"email"`
}

type Member struct {
	Silver   string
	Gold     string
	Platinum string
}

type Voucher struct {
	Id                       int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"`
	Created                  time.Time `json:"created"`
	CreatedBy                string    `json:"created_by"`
	Modified                 time.Time `json:"modified"`
	ModifiedBy               string    `json:"modified_by"`
	Active                   bool      `json:"active"`
	IsDeleted                bool      `json:"is_deleted"`
	Deleted                  time.Time `json:"deleted"`
	DeletedBy                string    `json:"deleted_by"`
	VoucherName              string    `json:"voucher_name"`
	StartDate                string    `json:"start_date"`
	EndDate                  string    `json:"end_date"`
	VoucherDescription       string    `json:"voucher_description"`
	VoucherTermsAndCondition string    `json:"voucher_terms_and_condition"`
	IsPushNotification       bool      `json:"is_push_notification"`
	IsGiveVoucher            bool      `json:"is_give_voucher"`
	VoucherPeriod            string    `json:"voucher_period"`
	RewardTermsAndCondition  string    `json:"reward_terms_and_condition"`
	BackgroundVoucherPattern string    `json:"background_voucher"`
	BackgroundVoucherColour  string    `json:"background_voucher_colour"`
	MerchantEmail            string    `json:"merchant_email"`
	OutletId                 int       `json:"outlet_id"`
	ProgramId                int       `json:"program_id"`

	MerchantName string `json:"merchant_name"`
	OutletName   string `json:"outlet_name"`
	ProgramName  string `json:"program_name"`
}

type Province struct {
	Id           int    `gorm:"PRIMARY_KEY;NOT NULL"; json:"id"`
	ProvinceName string `json:"province_name"`
}

type City struct {
	IdProvince int    `json:"id_province"`
	Id         int    `gorm:"PRIMARY_KEY;NOT NULL"; json:"id"`
	CityName   string `json:"city_name"`
}

type Reward struct {
	Id                int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"`
	Created           time.Time `json:"created"`
	CreatedBy         string    `json:"created_by"`
	Modified          time.Time `json:"modified"`
	ModifiedBy        string    `json:"modified_by"`
	Active            bool      `json:"active"`
	IsDeleted         bool      `json:"is_deleted"`
	Deleted           time.Time `json:"deleted"`
	DeletedBy         string    `json:"deleted_by"`
	RedeemPoints      int       `json:"redeem_points"`
	RewardName        string    `json:"reward_name"`
	RedeemRules       string    `json:"redeem_rules"`
	TermsAndCondition string    `json:"terms_and_condition"`
	ProgramId         int       `json:"program_id"`
	MerchantEmail     string    `json:"merchant_email"`
	Outletid          int       `json:"outletid"`

	MerchantName string `json:"merchant_name"`
	ProgramName  string `json:"program_name"`
	OutletName   string `json:"outlet_name"`
}

type Email struct {
	SenderEmail string           `json:"sender_email"`
	SenderName  string           `json:"sender_name"`
	Receiver    []ReceiverStruct `json:"receiver"`
	Subject     string           `json:"subject"`
	Body        string           `json:"body"`
	TextContent string           `json:"text_content"`
}

type ReceiverStruct struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (r *Email) ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(fmt.Sprintf("src/api/SendGrid/%s", templateFileName))
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return "error", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println("error exec : ", err.Error())
		return "nil", err
	}
	r.Body = buf.String()
	return buf.String(), nil
}

type EmailEmployee struct {
	SenderEmail   string `json:"sender_email"`
	SenderName    string `json:"sender_name"`
	EmployeeName  string `json:"employee_name"`
	EmployeeEmail string `json:"employee_email"`
	Subject       string `json:"subject"`
	Body          string `json:"body"`
	TextContent   string `json:"text_content"`
	EmployeePin   string `json:"employee_pin"`
}

func (r *EmailEmployee) ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(fmt.Sprintf("src/api/SendGrid/%s", templateFileName))
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return "error", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println("error exec : ", err.Error())
		return "nil", err
	}
	r.Body = buf.String()
	return buf.String(), nil
}

type EmailForgetPass struct {
	SenderEmail   string `json:"sender_email"`
	SenderName    string `json:"sender_name"`
	ReceiverName  string `json:"receiver_name"`
	ReceiverEmail string `json:"receiver_email"`
	Subject       string `json:"subject"`
	Body          string `json:"body"`
	TextContent   string `json:"text_content"`
}

func (r *EmailForgetPass) ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(fmt.Sprintf("src/api/SendGrid/%s", templateFileName))
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return "error", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println("error exec : ", err.Error())
		return "nil", err
	}
	r.Body = buf.String()
	return buf.String(), nil
}

type From struct {
	Email string `json:"email"`
}

type To struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type DynamicTemplateData struct {
	SenderName string `json:"sender_name"`
	Body       string `json:"body"`
	Pin        string `json:"pin"`
}

type Personalization struct {
	To                  []To                `json:"to"`
	DynamicTemplateData DynamicTemplateData `json:"dynamic_template_data"`
	Subject             string              `json:"subject"`
}

type Emails struct {
	From            From              `json:"from"`
	Personalization []Personalization `json:"personalizations"`
	TemplateId      string            `json:"template_id"`
}
