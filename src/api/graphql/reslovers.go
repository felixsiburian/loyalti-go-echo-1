package graphQL

import (
	"fmt"
	"github.com/graphql-go/graphql"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/repository"
)

type Song model.Song

func MerchantResolver(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println("masuk ke merchant resolver")
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	email, mail := p.Args["email"].(string)
	sort, tap := p.Args["sort"].(int)
	if ok && sip && tap && mail {
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		var mails *string = &email
		merchant := repository.GetMerchant(pages, sizes, sorts, mails)
		fmt.Println(merchant)
		return merchant, nil
	} else if ok && sip && tap {
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		merchant := repository.GetMerchant(pages, sizes, sorts, nil)
		fmt.Println(merchant)
		return merchant, nil
	} else if ok && sip && mail {
		var pages *int = &page
		var sizes *int = &size
		var mails *string = &email
		merchant := repository.GetMerchant(pages, sizes, nil, mails)
		fmt.Println(merchant)
		return merchant, nil
	} else if ok && sip {
		var pages *int = &page
		var sizes *int = &size
		merchant := repository.GetMerchant(pages, sizes, nil, nil)
		fmt.Println(merchant)
		return merchant, nil
	} else if mail {
		var mails *string = &email
		merchant := repository.GetMerchant(nil, nil, nil, mails)
		fmt.Println(merchant)
		return merchant, nil
	}

	merchant := repository.GetMerchant(nil, nil, nil, nil)
	fmt.Println(merchant)
	return merchant, nil
}

func MerchantCategoryResolver(p graphql.ResolveParams) (interface{}, error) {
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	sort, top := p.Args["sort"].(int)
	if ok && sip && top {
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		category := repository.GetCategory(pages, sizes, sorts)
		fmt.Println(category)
		return category, nil
	}

	category := repository.GetCategory(nil, nil, nil)

	return category, nil
}

func MerchantCardTypeResolver(p graphql.ResolveParams) (interface{}, error) {
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	sort, top := p.Args["sort"].(int)

	if ok && sip && top {
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		card := repository.GetCardType(pages, sizes, sorts)
		fmt.Println(card)
		return card, nil
	}
	card := repository.GetCardType(nil, nil, nil)
	return card, nil
}

func SocialMediaResolver(p graphql.ResolveParams) (interface{}, error) {
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	sort, top := p.Args["sort"].(int)
	if ok && sip && top {
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		sosmed := repository.GetSocialMedia(pages, sizes, sorts)
		fmt.Println(sosmed)
		return sosmed, nil
	}
	sosmed := repository.GetSocialMedia(nil, nil, nil)
	return sosmed, nil
}

func ProvinceResolver(p graphql.ResolveParams) (interface{}, error) {
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	sort, top := p.Args["sort"].(int)
	if ok && sip && top {
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		province := repository.GetProvince(pages, sizes, sorts)
		fmt.Println(province)
		return province, nil
	}
	province := repository.GetProvince(nil, nil, nil)
	return province, nil
}

func CityResolver(p graphql.ResolveParams) (interface{}, error) {
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	sort, top := p.Args["sort"].(int)
	if ok && sip && top {
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		province := repository.GetCity(pages, sizes, sorts)
		fmt.Println(province)
		return province, nil
	}
	province := repository.GetCity(nil, nil, nil)
	return province, nil
}

func SpecialProgramResolver(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println("masuk ke resolver")
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	sort, deh := p.Args["sort"].(int)
	category, cat := p.Args["category"].(int)
	email, mail := p.Args["email"].(string)

	if ok && sip && deh && cat && mail {
		fmt.Println("masuk ke 1")
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		var cats *int = &category
		var mails *string = &email
		program := repository.GetSpecialProgram(pages, sizes, sorts, cats, mails)
		fmt.Println(program)
		return program, nil
	} else if ok && sip && deh && cat {
		fmt.Println("masuk ke 2")
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		var cats *int = &category
		program := repository.GetSpecialProgram(pages, sizes, sorts, cats, nil)
		fmt.Println(program)
		return program, nil
	} else if ok && sip && deh && mail {
		fmt.Println("masuk ke 3")
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		var mails *string = &email
		program := repository.GetSpecialProgram(pages, sizes, sorts, nil, mails)
		fmt.Println(program)
		return program, nil
	} else if ok && sip && deh {
		fmt.Println("masuk ke 5")
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		program := repository.GetSpecialProgram(pages, sizes, sorts, nil, nil)
		fmt.Println(program)
		return program, nil
	} else if ok && sip && cat {
		fmt.Println("masuk ke 6")
		var pages *int = &page
		var sizes *int = &size
		var cats *int = &category
		program := repository.GetSpecialProgram(pages, sizes, nil, cats, nil)
		fmt.Println(program)
		return program, nil
	} else if ok && sip && mail {
		fmt.Println("masuk ke 7")
		var pages *int = &page
		var sizes *int = &size
		var mails *string = &email
		program := repository.GetSpecialProgram(pages, sizes, nil, nil, mails)
		fmt.Println(program)
		return program, nil
	} else if ok && sip {
		fmt.Println("masuk ke 8 ")
		var pages *int = &page
		var sizes *int = &size
		program := repository.GetSpecialProgram(pages, sizes, nil, nil, nil)
		fmt.Println(program)
		return program, nil
	} else if cat {
		fmt.Println("masuk ke 9")
		var cats *int = &category
		program := repository.GetSpecialProgram(nil, nil, nil, cats, nil)
		fmt.Println(program)
		return program, nil
	} else if mail {
		fmt.Println("masuk ke 10")
		var mails *string = &email
		program := repository.GetSpecialProgram(nil, nil, nil, nil, mails)
		fmt.Println(program)
		return program, nil
	}

	program := repository.GetSpecialProgram(nil, nil, nil, nil, nil)

	return program, nil
}

//program function
func ProgramResolver(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println("masuk ke resolver")
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	sort, deh := p.Args["sort"].(int)
	category, cat := p.Args["category"].(int)
	email, mail := p.Args["email"].(string)

	if ok && sip && deh && cat && mail {
		fmt.Println("masuk ke 1")
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		var cats *int = &category
		var mails *string = &email
		program := repository.GetProgram(pages, sizes, sorts, cats, mails)
		fmt.Println(program)
		return program, nil
	} else if ok && sip && deh && cat {
		fmt.Println("masuk ke 2")
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		var cats *int = &category
		program := repository.GetProgram(pages, sizes, sorts, cats, nil)
		fmt.Println(program)
		return program, nil
	} else if ok && sip && deh && mail {
		fmt.Println("masuk ke 3")
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		var mails *string = &email
		program := repository.GetProgram(pages, sizes, sorts, nil, mails)
		fmt.Println(program)
		return program, nil
	} else if ok && sip && deh {
		fmt.Println("masuk ke 5")
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		program := repository.GetProgram(pages, sizes, sorts, nil, nil)
		fmt.Println(program)
		return program, nil
	} else if ok && sip && cat {
		fmt.Println("masuk ke 6")
		var pages *int = &page
		var sizes *int = &size
		var cats *int = &category
		program := repository.GetProgram(pages, sizes, nil, cats, nil)
		fmt.Println(program)
		return program, nil
	} else if ok && sip && mail {
		fmt.Println("masuk ke 7")
		var pages *int = &page
		var sizes *int = &size
		var mails *string = &email
		program := repository.GetProgram(pages, sizes, nil, nil, mails)
		fmt.Println(program)
		return program, nil
	} else if ok && sip {
		fmt.Println("masuk ke 8 ")
		var pages *int = &page
		var sizes *int = &size
		program := repository.GetProgram(pages, sizes, nil, nil, nil)
		fmt.Println(program)
		return program, nil
	} else if cat {
		fmt.Println("masuk ke 9")
		var cats *int = &category
		program := repository.GetProgram(nil, nil, nil, cats, nil)
		fmt.Println(program)
		return program, nil
	} else if mail {
		fmt.Println("masuk ke 10")
		var mails *string = &email
		program := repository.GetProgram(nil, nil, nil, nil, mails)
		fmt.Println(program)
		return program, nil
	}

	program := repository.GetProgram(nil, nil, nil, nil, nil)

	return program, nil
}

func OutletResolver(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println("masuk ke resolver")
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	id, top := p.Args["id"].(int)
	email, mail := p.Args["email"].(string)
	if ok && sip && top && mail {
		var pages *int = &page
		var sizes *int = &size
		var outlet_id *int = &id
		var emails *string = &email
		outlet := repository.GetOutlet(pages, sizes, outlet_id, emails)
		fmt.Println(outlet)
		return outlet, nil
	} else if ok && sip && mail {
		var paging *int = &page
		var sizing *int = &size
		var emails *string = &email
		outlet := repository.GetOutlet(paging, sizing, nil, emails)
		return outlet, nil
	} else if ok && sip && top {
		var paging *int = &page
		var sizing *int = &size
		var merchant *int = &id
		outlet := repository.GetOutlet(paging, sizing, merchant, nil)
		return outlet, nil
	} else if ok && sip {
		var paging *int = &page
		var sizing *int = &size
		outlet := repository.GetOutlet(paging, sizing, nil, nil)
		return outlet, nil
	} else if mail {
		var emails *string = &email
		outlet := repository.GetOutlet(nil, nil, nil, emails)
		fmt.Println(outlet)
		return outlet, nil
	}
	outlet := repository.GetOutlet(nil, nil, nil, nil)
	return outlet, nil
}

func EmployeeResolver(p graphql.ResolveParams) (interface{}, error) {
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	sort, top := p.Args["id"].(int)
	if ok && sip && top {
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		employee := repository.GetEmployee(pages, sizes, sorts)
		fmt.Println(employee)
		return employee, nil
	} else if ok && sip {
		var paging *int = &page
		var sizing *int = &size
		employee := repository.GetEmployee(paging, sizing, nil)
		return employee, nil
	}
	employee := repository.GetEmployee(nil, nil, nil)
	return employee, nil
}

func TotalPointResolver(p graphql.ResolveParams)(interface{}, error){
	id := p.Args["id"].(int)
	pay := p.Args["pay"].(int)
	pin := p.Args["pin"].(string)
	outletid := p.Args["outletid"].(int)
	cardtype := p.Args["cardtype"].(string)
	var ids int = id
	var pays int = pay
	var pins string = pin
	var outletids int = outletid
	var cardtypes string = cardtype
	total := repository.TotalPoint(ids, pays, pins, outletids, cardtypes)
	return total, nil
}

func TotalChopResolver(p graphql.ResolveParams)(interface{}, error){
	id := p.Args["id"].(int)
	pay := p.Args["pay"].(int)
	pin := p.Args["pin"].(string)
	outletid := p.Args["outletid"].(int)
	cardtype := p.Args["cardtype"].(string)
	var ids int = id
	var pays int = pay
	var pins string = pin
	var outletids int = outletid
	var cardtypes string = cardtype
	total := repository.TotalChop(ids, pays, pins, outletids, cardtypes)
	return total, nil
}

//func MerchantCardMemberResolver(p graphql.ResolveParams)(interface{}, error){
//	fmt.Println("masuk ke resolver member")
//	program_id := p.Args["program_id"].(int)
//	var program_ids int = program_id
//	card := repository.GetCardMember(program_ids)
//	fmt.Println(card)
//	return card, nil
//}

func TransactionResolver(p graphql.ResolveParams) (interface{}, error) {
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	sort, top := p.Args["sort"].(int)
	outletid, tap := p.Args["outletid"].(int)
	if ok && sip && top && tap{
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		var outletids *int = &outletid
		transaction := repository.GetTransaction(pages, sizes, sorts, outletids)
		fmt.Println(transaction)
		return transaction, nil
	}else if ok && sip && top {
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		transaction := repository.GetTransaction(pages, sizes, sorts, nil)
		fmt.Println(transaction)
		return transaction, nil
	}else if ok && sip && tap {
		var pages *int = &page
		var sizes *int = &size
		var outletids *int = &outletid
		transaction := repository.GetTransaction(pages, sizes, nil, outletids)
		fmt.Println(transaction)
		return transaction, nil
	}else if tap{
		var outletids *int = &outletid
		transaction := repository.GetTransaction(nil, nil, nil, outletids)
		fmt.Println(transaction)
		return transaction, nil
	}

	transaction := repository.GetTransaction(nil, nil, nil, nil)

	return transaction, nil
}

func SongResolver(p graphql.ResolveParams) (interface{}, error) {
	users := []Song{
		Song{
			ID:       "1",
			Album:    "ts-fearless",
			Title:    "Fearless",
			Duration: "4:01",
			Type:     "song",
		},
		Song{
			ID:       "2",
			Album:    "ts-fearless",
			Title:    "Fifteen",
			Duration: "4:54",
			Type:     "song",
		},
	}
	return users, nil
	//return nil, nil
}

func CardResolver(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println("masuk ke resolver card")
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	id, top := p.Args["id"].(int)
	card_type, tipe := p.Args["card_type"].(string)

	var outlet []model.ProgramCard

	fmt.Println("keluar dari if member")

	if ok && sip && top && tipe {
		fmt.Println("masuk ke if pertama")
		var pages *int = &page
		var sizes *int = &size
		var top *int = &id
		var types *string = &card_type
		if (card_type == "Member") {
			fmt.Println("masuk ke if member")
			outlet = repository.GetCardMember(id)
		}
		outlet := repository.GetCardMerchant(pages, sizes, top, types)
		return outlet, nil
	} else if ok && sip && top {
		fmt.Println("masuk ke if kedua")
		var pages *int = &page
		var sizes *int = &size
		var card_id *int = &id
		outlet := repository.GetCardMerchant(pages, sizes, card_id, nil)

		fmt.Println(outlet)
		return outlet, nil
	} else if ok && sip {
		fmt.Println("masuk ke if ketiga")
		var paging *int = &page
		var sizing *int = &size
		outlet := repository.GetCardMerchant(paging, sizing, nil, nil)
		return outlet, nil
	} else if ok && sip && top {
		fmt.Println("masuk ke if keempat")
		var paging *int = &page
		var sizing *int = &size
		var program_id *int = &id
		outlet := repository.GetCardMerchant(paging, sizing, program_id, nil)
		return outlet, nil
	}
	fmt.Println("lewat semua")
	outlet = repository.GetCardMerchant(nil, nil, nil, nil)
	fmt.Println(outlet)
	return outlet, nil
}

func VoucherResolver(p graphql.ResolveParams) (interface{}, error) {
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	sort, top := p.Args["sort"].(int)
	merchant_id, tap := p.Args["merchant_email"].(string)
	if ok && sip && top && tap {
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		var merchant_ids *string = &merchant_id
		voucher := repository.GetVoucher(pages, sizes, sorts, merchant_ids)
		fmt.Println(voucher)
		return voucher, nil
	}
	if ok && sip && tap {
		var pages *int = &page
		var sizes *int = &size
		var merchant_ids *string = &merchant_id
		voucher := repository.GetVoucher(pages, sizes, nil, merchant_ids)
		fmt.Println(voucher)
		return voucher, nil
	}
	if ok && sip {
		var pages *int = &page
		var sizes *int = &size
		voucher := repository.GetVoucher(pages, sizes, nil, nil)
		fmt.Println(voucher)
		return voucher, nil
	}
	if ok && sip && top {
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		voucher := repository.GetVoucher(pages, sizes, sorts, nil)
		fmt.Println(voucher)
		return voucher, nil
	}
	if tap {
		var merchant_ids *string = &merchant_id
		voucher := repository.GetVoucher(nil, nil, nil, merchant_ids)
		fmt.Println(voucher)
		return voucher, nil
	}

	voucher := repository.GetVoucher(nil, nil, nil, nil)
	return voucher, nil
}

func RewardResolver(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println("masuk ke resolver")
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	sort, top := p.Args["sort"].(int)
	merchant_email, tap := p.Args["merchant_email"].(string)
	if ok && sip && top && tap {
		fmt.Println("masuk if 1")
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		var merchant_emails *string = &merchant_email
		voucher := repository.GetReward(pages, sizes, sorts, merchant_emails)
		fmt.Println(voucher)
		return voucher, nil
	}
	if tap {
		fmt.Println("masuk if 2")
		var merchant_emails *string = &merchant_email
		voucher := repository.GetReward(nil, nil, nil, merchant_emails)
		fmt.Println(voucher)
		return voucher, nil
	}
	voucher := repository.GetReward(nil, nil, nil, nil)
	return voucher, nil
}

func ReviewResolver(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println("masuk ke review resolver")
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	sort, top := p.Args["sort"].(int)
	program_name, name := p.Args["program_name"].(string)

	if ok && sip && top && name {
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		var names *string = &program_name
		review := repository.GetReview(pages, sizes, sorts, names)
		fmt.Println(review)
		return review, nil
	}

	if ok && sip && top {
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		review := repository.GetReview(pages, sizes, sorts, nil)
		fmt.Println(review)
		return review, nil
	}

	if ok && sip {
		var pages *int = &page
		var sizes *int = &size
		review := repository.GetReview(pages, sizes, nil, nil)
		fmt.Println(review)
		return review, nil
	}

	if name {
		var names *string = &program_name
		review := repository.GetReview(nil, nil, nil, names)
		fmt.Println(review)
		return review, nil
	}


	review := repository.GetReview(nil, nil, nil, nil)
	return review, nil
}
