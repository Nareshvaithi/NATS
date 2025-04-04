package main

import (
	// "fmt"
	"fmt"
	"log"
	"time"
	"github.com/nats-io/nats.go"
)

func main(){
	nc,err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v",err)
	}

	defer nc.Close()

	nc.Subscribe("intros",func(msg *nats.Msg) {
		fmt.Println("I got a message: %s\n",string(msg.Data))
	})
	nc.Subscribe("naresh",func(msg *nats.Msg) {
		fmt.Println("I am subscribe to naresh: %s\n",string(msg.Data))
	})
	time.Sleep(1*time.Hour)
}