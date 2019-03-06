package main

import (
	"log"

	"github.com/bitly/go-nsq"
)

func main() {
	config := nsq.NewConfig()
	address, _ := nsq.NewProducer("127.0.0.1:4150", config)

	err := address.Publish("kenbunshoku", []byte("defend"))
	if err != nil {
		log.Panic("Could not connect")
	}

	address.Stop()
}
