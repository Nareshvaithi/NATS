// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/nats-io/nats.go"
// )

// func main() {
// 	nc, err := nats.Connect(nats.DefaultURL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer nc.Close()

// 	// Enable JetStream
// 	js, err := nc.JetStream()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Create a stream (Only needed once)
// 	_, err = js.AddStream(&nats.StreamConfig{
// 		Name:     "CHAT",
// 		Subjects: []string{"chat.*"},
// 	})
// 	if err != nil {
// 		log.Println("Stream might already exist:", err)
// 	}

// 	sendTopic := "chat.user1"
// 	receiveTopic := "chat.user2"

// 	sub, err := js.PullSubscribe(receiveTopic, "user2Consumer")
// 	if err != nil {
// 		log.Fatal("Subscribe error:", err)
// 	}

// 	go func() {
// 		for {
// 			msgs, err := sub.Fetch(10, nats.MaxWait(2*time.Second))
// 			if err != nil {
// 				continue
// 			}

// 			for _, msg := range msgs {
// 				fmt.Printf("\nMessage from %s: %s", receiveTopic, string(msg.Data))
// 				msg.Ack()
// 			}
// 		}
// 	}()

// 	
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		fmt.Print("\nYou: ")
// 		msg, _ := reader.ReadString('\n')

// 		_, err := js.Publish(sendTopic, []byte(msg))
// 		if err != nil {
// 			log.Println("Publish error:", err)
// 		}
// 	}
// }

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func main(){
	const userName string = "user2"
	nc,err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatal("Error conecting",err)
	}

	defer nc.Close()

	js,err := nc.JetStream()

	if err != nil {
		log.Fatal(err)
	}

	_,err = js.AddStream(&nats.StreamConfig{
		Name: "CHAT",
		Subjects: []string{"chat.*"},
	}) 
	if err != nil {
		log.Fatal(err)
	}
	sendTopic := "chat.user1"
	receivedTopic := "chat.user2"
	sub,err := js.PullSubscribe(receivedTopic ,userName+"Consumer")
	if err != nil {
		log.Fatal(err)
	}
	go func(){
		for {
			msgs,err := sub.Fetch(10, nats.MaxWait(2 * time.Second))
			if err != nil {
				continue
			}

			for _,msg := range msgs {
				fmt.Printf("\n Message Received from %s: %s",receivedTopic, string(msg.Data))
				msg.Ack()
			}
		}
	}()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("YOU \n")
		msg,_ := reader.ReadString('\n')
		_,err := js.Publish(sendTopic,[]byte(msg))
		if err != nil {
			log.Fatal(err)
		}
		

	}
}