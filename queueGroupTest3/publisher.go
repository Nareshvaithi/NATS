package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main(){
	nc,natsErr := nats.Connect(nats.DefaultURL)

	if natsErr != nil {
		fmt.Println("Error Connecting NATS server",natsErr)
	}
	 defer nc.Close()
	
	 for i := 1 ; i <= 10 ; i ++ {
		message := fmt.Sprintf("TASK # %v",i)

		nc.Publish("jobs",[]byte(string(message)))
		fmt.Println("Published: ",message)
		time.Sleep(1 * time.Second)
	 } 
}