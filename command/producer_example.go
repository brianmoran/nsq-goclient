package main

import (
	"log"
	"github.com/bitly/go-nsq"
	"fmt"
)

func main() {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	for i := 0; i < 100000; i++ {
		err := w.Publish("nsq-demo-topic", []byte(fmt.Sprintf("Message: %v", i)))
		if err != nil {
			log.Panic("Could not connect")
		}
	}

	w.Stop()
}
