package main

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://127.30.0.2:1883")
	c := mqtt.NewClient(opts)

	token := c.Connect()
	if token.Wait() && token.Error() != nil {
		log.Fatalf("Mqtt Error: %s", token.Error())
	}

	for i := 0; i < 1; i++ {
		tex := fmt.Sprintf("This is mes #%d", i)
		t := c.Publish("go-mqtt/test", 0, true, tex)
		t.Wait()
	}
	c.Disconnect(250)
	fmt.Println("Fin")
}
