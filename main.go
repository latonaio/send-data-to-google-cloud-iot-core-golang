package main

import (
	"os"
	"send-data-to-google-cloud-iot-core-golang/model"
	senddatagoogleiot "send-data-to-google-cloud-iot-core-golang/sendDataGoogleIoT"

	"github.com/latonaio/golang-logging-library/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client"
)

func main() {

	// model definition
	var option model.CommunucationMQTT
	option.Projects = os.Getenv("GOOGLE_PROJECT")
	option.Registry = os.Getenv("GOOGLE_REGISTRY")
	option.Device = os.Getenv("GOOGLE_DEVICE")
	option.Algorithm = os.Getenv("ALGORITHM")
	option.Private_key = os.Getenv("PRIVATE_KEY")
	option.Region = os.Getenv("REGION")
	option.Server = os.Getenv("SERVER")
	option.MacAddress = os.Getenv("MAC_ADDRESS")
	option.Terminal = os.Getenv("TERMINAL_NAME")

	var (
		queueFrom  = os.Getenv("TERMINAL_NAME")
		rabbimqURL = os.Getenv("RABBITMQ_URL")
	)

	logger := logger.NewLogger()
	client, err := rabbitmq.NewRabbitmqClient(
		rabbimqURL,
		[]string{queueFrom},
		[]string{},
	)
	if err != nil {
		logger.Error("can't connect to rabbitmq")
		return
	}

	// 受信を終了する
	defer client.Close()

	iter, err := client.Iterator()
	if err != nil {
		logger.Error("can't iteration")
		return
	}
	// 受信を終了する際には Stop を呼ぶ
	defer client.Stop()

	for message := range iter {
		logger.Info("get data from %v, data: %v", message.QueueName(), message.Data())

		senddatagoogleiot.SendDataGoogleIot(message.Data(), option)
		logger.Info("send data to google iot core")
		message.Success()
	}

}
