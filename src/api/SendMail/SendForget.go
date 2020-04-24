package SendMail

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"log"
)

func SendForget(pass interface{}) error {
	request := sendgrid.GetRequest("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", "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	fmt.Println("isi email : ", pass)
	//os.Exit(1)
	byte , err := json.Marshal(pass)
	if err != nil {
		fmt.Println("Error Marshal : ", err.Error())
		fmt.Println(string(byte))
		log.Fatal(err)
	}
	request.Body = byte

	response,err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}else{
		fmt.Print(response.StatusCode)
		fmt.Print(response.Body)
		fmt.Print(response.Headers)
	}
	return nil
}