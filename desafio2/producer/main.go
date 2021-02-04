package consumer

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sohamkamani/golang-kafka-example/application/kafka"
	"github.com/sohamkamani/golang-kafka-example/infrastructure/db"
	"os"
)

func main() {
	deliveryChan := make(chan ckafka.Event)
	database := db.ConnectDB(os.Getenv("env"))
	producer := kafka.NewKafkaProducer()

	kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
	kafkaProcessor.Consume()

}

