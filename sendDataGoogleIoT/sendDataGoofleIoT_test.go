package senddatagooglecloudiot

import (
	"send-data-to-google-cloud-iot-core-golang/model"
)

var option model.CommunucationMQTT

func ExamplePrintAllAttribute() {
	printTestSetAllAttribute()
	// Output:
	// Client ID: xxxx
	// Connecting...
	// Publishing: xxxx
	// Disconnecting...

}

func ExamplePrintSetOnlyAttribute() {
	printTestSetOnlyImagePathAttribute()
	// Output:
	// Client ID: xxxx
	// Connecting...
	// Publishing: xxxx
	// Disconnecting...
	// Client ID: xxxx
	// Connecting...
	// Publishing: xxxx
	// Disconnecting...
	// Client ID: xxxx
	// Connecting...
	// Publishing: xxxx
	// Disconnecting...
}

func ExamplePrintSetIrregular() {
	printTestSetIrregular()
	// Output:
	// Client ID: xxxx
	// Connecting...
	// Publishing: xxxx
	// Disconnecting...
}

// 期待する属性を全て割り当てた場合のテスト
func printTestSetAllAttribute() {
	option.Projects = "xxxx"
	option.Registry = "sample"
	option.Device = "sample"
	option.Algorithm = "RS256"
	option.Private_key = "../rsa_private.pem"
	option.Region = "asia-east1"
	option.Server = "ssl://mqtt.googleapis.com:8883"
	option.MacAddress = "xxxx"
	option.Terminal = "xxxx"

	var testStruct = map[string]interface{}{
		"faceId":       "test",
		"imagePath":    "test",
		"responseData": "test",
	}

	SendDataGoogleIot(testStruct, option)
}

// 期待する属性を一部割り当てた場合のテスト
func printTestSetOnlyImagePathAttribute() {
	option.Projects = "xxxx"
	option.Registry = "sample"
	option.Device = "sample"
	option.Algorithm = "RS256"
	option.Private_key = "../rsa_private.pem"
	option.Region = "asia-east1"
	option.Server = "ssl://mqtt.googleapis.com:8883"
	option.MacAddress = "xxxx"
	option.Terminal = "xxxx"

	var testStructOnlyFaceId = map[string]interface{}{
		"faceId": "test",
	}

	SendDataGoogleIot(testStructOnlyFaceId, option)

	var testStructOnlyImagePath = map[string]interface{}{
		"imagePath": "test",
	}
	SendDataGoogleIot(testStructOnlyImagePath, option)

	var testStructOnlyResponseData = map[string]interface{}{
		"responseData": "test",
	}
	SendDataGoogleIot(testStructOnlyResponseData, option)

}

// 期待する属性を割り当てず、関係ない属性を割り当てた場合のテスト
func printTestSetIrregular() {
	option.Projects = "xxxx"
	option.Registry = "sample"
	option.Device = "sample"
	option.Algorithm = "RS256"
	option.Private_key = "../rsa_private.pem"
	option.Region = "asia-east1"
	option.Server = "ssl://mqtt.googleapis.com:8883"
	option.MacAddress = "xxxx"
	option.Terminal = "xxxx"

	var testStructIrregular = map[string]interface{}{
		"hogehoge": "test",
	}

	SendDataGoogleIot(testStructIrregular, option)

}
