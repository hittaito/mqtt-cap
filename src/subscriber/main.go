package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	ch := make(chan mqtt.Message)
	var f mqtt.MessageHandler = func(c mqtt.Client, msg mqtt.Message) {
		ch <- msg
	}
	opts := mqtt.NewClientOptions()
	opts.SetCleanSession(true)
	opts.SetWill("go-mqtt/will", "will msg", 0, false)
	opts.AddBroker("tcp://172.30.0.2:1883")
	opts.SetKeepAlive(10)
	c := mqtt.NewClient(opts)

	token := c.Connect()
	if token.Wait() && token.Error() != nil {
		log.Fatalf("Mqtt error : %s", token.Error())
	}

	subsc := c.Subscribe("go-mqtt/test", 1, f)
	if subsc.Wait() && subsc.Error() != nil {
		log.Fatal(subsc.Error())
	}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	for {
		select {
		case m := <-ch:
			fmt.Printf("Topic: %v, payload: %v\n", m.Topic(), string(m.Payload()))
		case <-sigCh:
			fmt.Printf("Interrupt")
			c.Disconnect(1000)
			return
		}
	}
}
