package main

import (
	"log"
	"github.com/bitly/go-nsq"
	"fmt"
)

func main() {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	for i := 0; i < 1000; i++ {
		err := w.Publish("write_test", []byte(fmt.Sprintf("Message: %v", i)))
		if err != nil {
			log.Panic("Could not connect")
		}
	}

	w.Stop()
}