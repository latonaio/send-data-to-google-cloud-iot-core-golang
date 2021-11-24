package model

/* google cloud iot core との通信に必要なモデル定義
*  MQTTクライアントの設定に使われる
*  https://cloud.google.com/iot/docs/how-tos/mqtt-bridge
 */
type CommunucationMQTT struct {
	Projects    string
	Registry    string
	Device      string
	Algorithm   string
	Private_key string
	Region      string
	Server      string
	MacAddress  string
	Terminal    string
}

type MqMessageData struct {
	ImagePath    interface{}
	FaceId       interface{}
	ResponseData interface{}
}
