package akafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func Produce(topic string, servers string, msg []byte) error {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
	})

	if err != nil {
		log.Printf("Failed to create producer: %v", err)
	}
	defer producer.Close()

	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg,
	}, nil)

	if err != nil {
		log.Printf("Failed to produce message: %s", err)
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		log.Printf("Delivery failed: %v", m.TopicPartition.Error)
		return m.TopicPartition.Error
	}

	log.Printf("Delivered message to topic %s [%d] at offset %v\n", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	return nil
}

func Consume(topics []string, servers string, groupId string, msgChan chan *kafka.Message) {
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          groupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Printf("Apache Kafka consuming error: %s", err.Error())
	}
	defer kafkaConsumer.Close()

	err = kafkaConsumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Printf("Apache Kafka topic subscribing error: %s", err.Error())
	}

	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			msgChan <- msg
		}
	}
}
