package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main(){
	nc,err := nats.Connect(nats.DefaultURL)

	if err != nil {
		fmt.Println(err)
	}
	count := 0
	for {
		data := fmt.Sprintf("hello, I am Naresh This my Request count %v",count)
		reply,err := nc.Request("intros",[]byte(data),500*time.Millisecond)
		time.Sleep(1*time.Second)
		if err != nil {
			fmt.Printf("error sending message count = %v, err: %v\n",count,err)
			continue
		}
		count ++
		log.Printf("send message %v, got reply %v",count,string(reply.Data))
	}
}