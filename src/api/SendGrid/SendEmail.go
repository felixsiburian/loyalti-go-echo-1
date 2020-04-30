package SendGrid

import (
	//"encoding/json"
	"fmt"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	//"github.com/labstack/echo"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)


func SendMail(email *model.Email) error {
	fmt.Println("Masuk ke SendGrid")
	e := email
	fmt.Println("email : ", email)
	fmt.Println("e : ", e)

	var a = "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"
	from := mail.NewEmail(e.SenderName, e.SenderEmail)
	subject := e.Subject
	for i := range e.Receiver {
		fmt.Println("masuk ke perulangan")
		to := mail.NewEmail(e.Receiver[i].Name, e.Receiver[i].Email)
		plainTextContent := "abcd"
		htmlContent, err := e.ParseTemplate("SendEmailTemplate.html", e)
		if err != nil {
			fmt.Println("Error HTML : ", err.Error())
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
	}
	fmt.Println("Selesai Tanpa Error : ", e)
	return nil
}
