package SendGrid

import (
//"encoding/json"
"fmt"
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
	for i := range e.Receiver {
		fmt.Println("masuk ke perulangan")
		to := mail.NewEmail(e.Receiver[i].Name, e.Receiver[i].Email)
		plainTextContent := "abcd"
		htmlContent := "Don't worry, we got your back!. " + "You can reset your password by clicking on this " +
			"<a href=https://loyalti-express-dev-ui.azurewebsites.net/resetpassword>Reset Password</a>"
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

