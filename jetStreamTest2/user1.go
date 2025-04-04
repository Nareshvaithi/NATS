package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)


func main() {
	const userName string = "user1"
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()


	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "CHAT",
		Subjects: []string{"chat.*"},
	})

	if err != nil {
		log.Println("Stream might already exist:", err)
	}
	
	sendTopic := "chat.user2" 
	receiveTopic := "chat.user1"
	
	sub, err := js.PullSubscribe(receiveTopic, userName+"Consumer")
	if err != nil {

		log.Fatal("Subscribe error:", err)
	}

	go func() {
		for {
			msgs, err := sub.Fetch(10, nats.MaxWait(2*time.Second)) 
			if err != nil {
				continue
			}
			for _, msg := range msgs {
				fmt.Printf("\nMessage from %s: %s", receiveTopic, string(msg.Data))
				msg.Ack() 
			}
		}
	}()

	
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\nYou: ")
		msg, _ := reader.ReadString('\n')

		_, err := js.Publish(sendTopic, []byte(msg))
		if err != nil {
			log.Println("Publish error:", err)
		}
	}
}