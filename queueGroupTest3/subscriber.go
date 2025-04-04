package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func main(){
	nc,natsErr := nats.Connect(nats.DefaultURL)

	if natsErr != nil {
		fmt.Println("Error Connecting NATS ",natsErr)
	}
	defer nc.Close()
	
	_,subErr1 := nc.QueueSubscribe("jobs","workers",func(msg *nats.Msg) {
		fmt.Printf("workers group jobs Subscriber 1 received: %s\n", string(msg.Data))
	})
	if subErr1 != nil {
		fmt.Println("Error subscribing: %v",subErr1)
	}

	_,subErr2 := nc.QueueSubscribe("jobs","workers",func(msg *nats.Msg) {
		fmt.Printf("workers group jobs Subscriber 2 received: %s\n", string(msg.Data))
	})
	if subErr2 != nil {
		fmt.Println("Error subscribing: %v",subErr2)
	}
	select {}
}