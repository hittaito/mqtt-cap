# MQTT cap test

MQTTのパケットキャプチャがしたかっただけ

```CLI
// brokerスタート
docker-compose up

// 購読開始
make subscribe

// publish
go run ./src/publisher

// パケットキャプチャ
make cap
```

キャプチャはdist以下に保存