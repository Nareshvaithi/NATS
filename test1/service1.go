package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	
	nc.Subscribe("user1.inbox", func(m *nats.Msg) {
		fmt.Println("User 2:", string(m.Data))
	})

	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Print("You: \n")
		msg, _ := reader.ReadString('\n')
		nc.Publish("user2.inbox", []byte(msg))
	}
}
