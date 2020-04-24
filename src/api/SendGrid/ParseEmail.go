package SendGrid

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

