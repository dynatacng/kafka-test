package main

import (
	"context"

	kappa "github.com/kafka-test/pkg"
	"github.com/sirupsen/logrus"
)

const (
	topic  = "go-kafka-topic"
	b1addr = "localhost:9093"
	b2addr = "localhost:9094"
	b3addr = "localhost:9095"
	b4addr = "localhost:9096"
	b5addr = "localhost:9097"
)

func main() {
	ctx := context.Background()

	brokers := []string{b1addr, b2addr, b3addr, b4addr, b5addr}
	topic := "kafka-topic-demo"

	writeConfig := kappa.NewWriterConfig(topic, brokers)
	producer := &kappa.Producer{
		WriterConfig: writeConfig,
		Logger:       logrus.WithField("module", "Producer"),
	}

	readConfig := kappa.NewReaderConfig(topic, brokers)
	consumer := &kappa.Consumer{
		ReaderConfig: readConfig,
		Logger:       logrus.WithField("module", "Consumer"),
	}
	go producer.Produce(ctx)
	consumer.Consume(ctx)

}
