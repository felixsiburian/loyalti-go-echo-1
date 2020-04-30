package SendGrid

import (
//"encoding/json"
"fmt"
	"github.com/labstack/gommon/log"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
//"github.com/labstack/echo"
"github.com/sendgrid/sendgrid-go"
"github.com/sendgrid/sendgrid-go/helpers/mail"
)


func SendForgetPass(email *model.EmailForgetPass) error {
	fmt.Println("Masuk ke SendGrid forget pass")
	e := email
	fmt.Println("email : ", email)
	fmt.Println("e : ", e)
	//err := json.NewDecoder(c.Request().Body).Decode(&e)
	//if err != nil {
	//	fmt.Println("Error SendGrid : ", err.Error())
	//}

	//var e Email
	//var a = fmt.Sprintf("%s", viper.Get("SENDGRIDKEY"))
	//fmt.Println(a)
	var a = "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"
	from := mail.NewEmail(e.SenderName, e.SenderEmail)
	subject := e.Subject
		fmt.Println("masuk ke perulangan")
		to := mail.NewEmail(e.ReceiverName, e.ReceiverEmail)
		plainTextContent := "abcd"
		htmlContent, err := e.ParseTemplate("ForgetPasswordEmailTemplate.html", e)
		if err != nil {
			log.Fatal(err)
		}
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		client := sendgrid.NewSendClient(a)
		fmt.Println("isi sendgrid ", a)
		response, err := client.Send(message)
		if err != nil {
			fmt.Println("Error response : ", err.Error())
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(response.Body)
			fmt.Println(response.Headers)
		}

		if err == nil {
			fmt.Println("mail sent")
		}

	fmt.Println("Selesai Tanpa Error : ", e)
	return nil
}

