package senddatagooglecloudiot

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"send-data-to-google-cloud-iot-core-golang/model"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MQTT parameters
const (
	topicType = "events" // or "state"
	qos       = 1        // QoS 2 isn't supported in GCP
	retain    = false
	username  = "unused" // always this value in GCP
)

func SendDataGoogleIot(mqMessageData map[string]interface{}, option model.CommunucationMQTT) {
	var sendToGoogleIoT model.MqMessageData

	// 初期値
	sendToGoogleIoT.FaceId = nil
	sendToGoogleIoT.ImagePath = nil
	sendToGoogleIoT.ResponseData = nil
	// generate MQTT client
	client_id := fmt.Sprintf(
		"projects/%s/locations/%s/registries/%s/devices/%s",
		option.Projects, option.Region, option.Registry, option.Device)

	fmt.Println("Client ID:", client_id)

	// load private key
	keyData, err := ioutil.ReadFile(option.Private_key)
	if err != nil {
		panic(err)
	}

	var key interface{}
	switch option.Algorithm {
	case "RS256":
		key, err = jwt.ParseRSAPrivateKeyFromPEM(keyData)
	case "ES256":
		key, err = jwt.ParseECPrivateKeyFromPEM(keyData)
	default:
		panic(fmt.Errorf("unknown algorithm: %s", option.Algorithm))
	}
	if err != nil {
		panic(err)
	}

	// generate JWT as the MQTT password
	t := time.Now()
	token := jwt.NewWithClaims(jwt.GetSigningMethod(option.Algorithm), &jwt.StandardClaims{
		IssuedAt:  t.Unix(),
		ExpiresAt: t.Add(time.Minute * 20).Unix(),
		Audience:  option.Projects,
	})
	pass, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}

	// configure MQTT client
	opts := mqtt.NewClientOptions().
		AddBroker(option.Server).
		SetClientID(client_id).
		SetUsername(username).
		SetTLSConfig(&tls.Config{MinVersion: tls.VersionTLS12}).
		SetPassword(pass).
		SetProtocolVersion(4) // Use MQTT 3.1.1.

	conn := mqtt.NewClient(opts)

	// connect to GCP Cloud IoT Core
	fmt.Println("Connecting...")
	tok := conn.Connect()
	if err := tok.Error(); err != nil {
		panic(err)
	}
	if !tok.WaitTimeout(time.Second * 5) {
		panic(fmt.Errorf("connection Timeout"))
	}
	if err := tok.Error(); err != nil {
		panic(err)
	}

	// generate topic
	topic := fmt.Sprintf("/devices/%s/%s", option.Device, topicType)

	// メッセージキューから特定の key を受け取った場合のみ該当する value を入れる

	for index, value := range mqMessageData {
		if index == "faceId" {
			sendToGoogleIoT.FaceId = value
		}
		if index == "imagePath" {
			sendToGoogleIoT.ImagePath = value
		}
		if index == "responseData" {
			sendToGoogleIoT.ResponseData = value
		}
	}

	var data = map[string]interface{}{
		"terminalName": option.Terminal,
		"macAddress":   option.MacAddress,
		"imagePath":    sendToGoogleIoT.ImagePath,
		"faceId":       sendToGoogleIoT.FaceId,
		"responseData": sendToGoogleIoT.ResponseData,
	}

	json, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Publishing: '%s'\n", json)
	conn.Publish(topic, qos, retain, json)
	time.Sleep(time.Millisecond * 500)

	// need mqtt reconnect each 20 minutes for long use

	// disconnect
	fmt.Println("Disconnecting...")
	conn.Disconnect(1000)
}
