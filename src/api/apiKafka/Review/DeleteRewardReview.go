package Review

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

func PublishDeleteReview(c echo.Context) error {
	//var data model.Merchant
	data := new(model.Review)
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	//err := c.Bind(data)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
		panic(err)
	}

	fmt.Println(data)

	if data.Rating < 0 {
		return c.String(http.StatusBadRequest, "Your rating is invalid")
	} else if data.Rating > 5 {
		return c.String(http.StatusBadRequest, "Your rating is invalid")
	}

	//if len(data.MerchantEmail) == 0 && len(data.MerchantName) == 0 {
	//	return c.String(http.StatusBadRequest, "Nama dan Email kosong")
	//} else if len(data.MerchantEmail) == 0 {
	//	return c.String(http.StatusBadRequest, "Email tidak boleh kosong")
	//} else if len(data.MerchantName) == 0 {
	//	return c.String(http.StatusBadRequest, "Name tidak boleh kosong")
	//}

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

	var newTopic = "delete-review-topic"

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
