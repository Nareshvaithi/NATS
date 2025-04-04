package main

import (
	// "fmt"
	"fmt"
	"log"
	"time"

	// "time"

	"github.com/nats-io/nats.go"
)

func main(){
	nc,err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v",err)
	}

	defer nc.Close()

	for {
		err := nc.Publish("intros",[]byte("Hello World!"))
		if err != nil {
			fmt.Println(err)
		}
		nareshErr := nc.Publish("naresh",[]byte("hey i am Naresh Thanks for sunbscribe"))
		if nareshErr != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
	}

	
}
	