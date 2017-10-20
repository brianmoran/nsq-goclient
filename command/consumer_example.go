package main

import (
	"log"
	"sync"

	"github.com/bitly/go-nsq"
)

var wg = &sync.WaitGroup{}

func handler(message *nsq.Message) error {
	log.Printf("Got a message: %v", string(message.Body))
	wg.Done()
	return nil
}

func main() {

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("write_test", "ch", config)
	q.AddHandler(nsq.HandlerFunc(handler))
	err := q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}

	for {
		wg.Add(1)
		wg.Wait()
	}


}