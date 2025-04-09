package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main(){
	nc,err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatal("NATS connecting Error: ",err)
	}
	defer nc.Close()
	nc.Subscribe("myMsg",func(msg *nats.Msg){
		fmt.Println(string(msg.Data))
	})
	select {}
}