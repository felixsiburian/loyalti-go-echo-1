package SendGrid

import (
	//"encoding/json"
	"fmt"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"github.com/spf13/viper"
	//"github.com/labstack/echo"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)


func SendMail(email *model.Email) error {
	fmt.Println("Masuk ke SendGrid")
	e := email
	//err := json.NewDecoder(c.Request().Body).Decode(&e)
	//if err != nil {
	//	fmt.Println("Error SendGrid : ", err.Error())
	//}

	//var e Email
	var a = fmt.Sprintf("%s", viper.Get("SENDGRIDKEY"))
	fmt.Println(a)
	//var a = "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"
	from := mail.NewEmail(e.SenderName, e.SenderEmail)
	subject := e.Subject
	for i := range e.Receiver {
		fmt.Println("masuk ke perulangan")
		to := mail.NewEmail(e.Receiver[i].Name, e.Receiver[i].Email)
		plainTextContent := "abcd"
		htmlContent := e.Body
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
	}
	fmt.Println("Selesai Tanpa Error : ", e)
	return nil
}
