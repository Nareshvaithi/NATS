package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func main(){
	nc,err := nats.Connect(nats.DefaultURL)

	if err != nil {
		fmt.Println("connection err ",err)
	}

	nc.Subscribe("intros",func(msg *nats.Msg) {
		var count int 
		var tem string
		data := string(msg.Data)
		fmt.Sscanf(data,"%s %v",&tem,&count)
		replyData := fmt.Sprintf("ack message # %v",count)
		msg.Respond([]byte(replyData))
		fmt.Printf("I got a message: %s\n",data)
	})
}