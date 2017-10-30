package main

import (
	"log"
	"sync"

	"github.com/bitly/go-nsq"
)

type handler struct {}
func (*handler) HandleMessage(message *nsq.Message) error {
	log.Printf("Got a message: %v", string(message.Body))
	wg.Done()
	return nil
}

var wg = &sync.WaitGroup{}

func main() {

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("nsq-demo-topic", "nsq-demo-channel", config)
	q.AddConcurrentHandlers(&handler{}, 5)
	err := q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}

	for {
		wg.Add(1)
		wg.Wait()
	}


}
