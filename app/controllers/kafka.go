package controllers

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/Shopify/sarama"
	"fullcontact-gin/app/core"
	"fullcontact-gin/app/models"
)

var dataCollector sarama.SyncProducer

type KafkaController struct {
}

func (c *KafkaController) Publish(message []byte) {
	if dataCollector == nil {
		log.Println("Init c.dataCollector", dataCollector)
		brokerList := strings.Split(core.Cfg.Kafka.Brokers, ",")
		dataCollector = newDataCollector(brokerList)
	}

	partition, offset, err := dataCollector.SendMessage(&sarama.ProducerMessage{
		Topic: core.Cfg.Kafka.DefaultTopic,
		Value: sarama.ByteEncoder(message),
	})

	if err != nil {
		panic(err)
		fmt.Printf("Your data is stored with unique identifier important/%d/%d", partition, offset)
	}
}

func (c *KafkaController) PublishContact(contact models.Person) {
	kafka := KafkaController{}
	itemStr, err := json.Marshal(contact)
	if err != nil {
		panic(err)
	}
	kafka.Publish(itemStr)
}

func newDataCollector(brokerList []string) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	tlsConfig := getDefaultTlsConfig()
	if tlsConfig != nil {
		config.Net.TLS.Config = tlsConfig
		config.Net.TLS.Enable = true
	}

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	return producer
}

func getDefaultTlsConfig() (t *tls.Config) {
	return t
}
