# send-data-to-google-cloud-iot-core-golang
AION では、主にエッジコンピューティング環境からデータを収集・管理維持するための Data Stack / Data Hub として、Google Cloud IoT Core を、一つの重要な選択肢として選択しています。       
AION では、Google Cloud IoT Core にエッジ環境からデータを送信するときに、Pub/Subメッセージングモデルを採用した MQTT による通信が利用されています。  
本マイクロサービスは、エッジで RabbitMQ のキューから受け取ったメッセージを、エッジから Google Cloud IoT Core に送信するためのマイクロサービスです。送信されたメッセージは Pub/Sub のメッセージから確認することができますが、トピックにスタックしているので、別途メッセージをデータとして活用したい場合は Google Cloud Storage などからサブスクライブしてください。
本マイクロサービスは、エッジ環境側にセットアップされるリソースです。 

## 動作環境

* エッジコンピューティング環境（本マイクロサービスがデプロイされるエッジ端末）  
* OS: Linux OS  
* CPU: ARM/AMD/Intel  
* RabbitMQ もしくは AION 導入済みのエッジ環境  
* Kubernetes 導入済みのエッジ環境  
* Golang Runtime  


## セットアップ
### PUB/SUB　の設定
本サービスはメッセージの送受信にGoogle Pub/Subを利用します。
以下は、Pub/Subの設定例です。

```
1. Pub/Subのコンソール画面に移動する
2. メニューバー上部「トピックを作成」をクリック、トピックを作成する。
3. トピックIDを sample としてトピックを作成する。
```

### IoT Core　の設定
本サービスは端末の管理をGoogle IoT Coreを利用します。
以下は、IoT Core レジストリの設定例です。

```
1. IoT Coreのコンソール画面に移動する
2. メニューバー上部「レジストリを作成」をクリック、レジストリを作成する。
3. レジストリIDを sample と設定
4. リージョンを asia-east1 と設定
5. Cloud Pub/Sub トピックを 「Pub/Subの設定」で作成した project/yourProjectName/topics/sample を設定

```

また、各種端末からメッセージを受けとるため、メニューバー左にある「デバイス」から端末の設定を行います。
以下は、IoT Core デバイスの設定例です。

```
1. IoT Coreのコンソール画面に移動する
2. メニューバー左 「デバイス」をクリック
3. メニューバー上部 「デバイスを作成」をクリック
4. デバイスIDを sample と設定
5. 作成をクリック
6. 作成したデバイスIDをクリック
7. 詳細、構成、認証のタブのうち、認証タブをクリック
8. 公開鍵を追加をクリック
9. 以下のコマンドを実行
    openssl genrsa -out rsa_private.pem 2048
    openssl rsa -in rsa_private.pem -pubout -out rsa_cert.pem
10.rsa_cert.pem　の中身を「公開鍵の値」へ貼り付ける
11.追加をクリック
```

### send-data-to-google-cloud-iot-core-golangを利用する

1. git clone 
```
git clone git@github.com:latonaio/send-data-to-google-cloud-iot-core-golang.git
```

2. goのパッケージをインストール
```
go mod download
go mod tidy
```

3. 秘密鍵を配置する
send-data-to-google-cloud-iot-core-golangのルートディレクトリに先ほど作成した「rsa_private.pem」を置きます。

4. .env を作成する
.envを作成します。REGISTRY名などはセットアップで設定した値を利用しています。
```
	GOOGLE_PROJECT=xxxx　　　　　　　　　　　　// GCPのプロジェクトID
	GOOGLE_REGISTRY=sample                 // IoT Core で設定したレジストリID
	GOOGLE_DEVICE=sample                   // IoT Core で設定したデバイスID
	ALGORITHM=RS256                        // 暗号化方式の指定  
	PRIVATE_KEY=./rsa_private.pem          // 利用する秘密鍵のパス
	REGION=asia-east1                      // IoT Core で設定したリージョン
	SERVER=ssl://mqtt.googleapis.com:8883  // GoogleのMQTTサーバ　[詳細](https://cloud.google.com/iot/docs/how-tos/mqtt-bridge)
	MAC_ADDRESS=xxxx                       // メッセージを送信したいデバイスのMACアドレス
	TERMINAL_NAME=xxxx                     // メッセージを送信したいデバイスの名前
  RABBITMQ_URL=xxxx                      // rabbitmq送信元のホスト名
  QUEUE_ORIGIN=xxxx                      // rabbitmq送信元のキュー名
```

5. テストを実行する
接続確認のテストを行います。
./sendDataGoogleIoT/sendDataGoogleIoT.go 内の.env で設定したものと同様の設定をテストする関数のoption構造体に設定してください。

設定例
```
	option.Projects = "xxxx"
	option.Registry = "sample"
	option.Device = "sample"
	option.Algorithm = "RS256"
	option.Private_key = "../rsa_private.pem"
	option.Region = "asia-east1"
	option.Server = "ssl://mqtt.googleapis.com:8883"
	option.MacAddress = "xxxx"
	option.Terminal = "xxxx"
```

テストを実行する（そのままではテストに失敗します。Projects などを適切に設定してください。）
```
go test ./sendDataGoogleIoT
```

