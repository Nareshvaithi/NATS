package main

import (
	"log"
	"github.com/nats-io/nats.go"
)

func main(){
	nc,err := nats.Connect(nats.DefaultURL) 

	if err != nil {
		log.Fatal("NATS conneting error: ",err)
	}
	defer nc.Close()
	nc.Publish("myMsg",[]byte(string("Hi i am here")))
	
}

