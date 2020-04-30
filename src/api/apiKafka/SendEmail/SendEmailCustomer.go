package SendEmail

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/labstack/echo"
	"github.com/radyatamaa/loyalti-go-echo/src/api/host/Config"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"log"
	"net/http"
)

func PublishSendEmailCustomer(c echo.Context) error {
	//var data model.Merchant
	data := new(model.Email)
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	//err := c.Bind(data)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
		panic(err)
	}
	//data.TemplateId = "d-287e301e6f0b4d29b4f7f1143bc9697c"
	fmt.Println(data)

	kafkaConfig := Config.GetKafkaConfig("", "")

	producer, err := sarama.NewSyncProducer([]string{"52.185.161.109:9092"}, kafkaConfig)

	if err != nil {

		panic(err)

	}

	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	var newTopic = "send-email-topic"

	message, _ := json.Marshal(data)
	//message := `{
	//	"merchant_name":`+data.MerchantName+`,
	//	"merchant_email":"contact@jco.com"
	//}`
	msg := &sarama.ProducerMessage{
		Topic: newTopic,
		Value: sarama.StringEncoder(string(message)),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", newTopic, partition, offset)
	return c.JSON(http.StatusOK, string(message))
}
