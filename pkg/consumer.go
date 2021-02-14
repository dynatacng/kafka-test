package kappa

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type consumer interface {
	consume(ctx context.Context)
}

type Consumer struct {
	ReaderConfig kafka.ReaderConfig
	Logger       *logrus.Entry
}

func (c *Consumer) Consume(ctx context.Context) {

	kr := kafka.NewReader(c.ReaderConfig)
	for {
		msg, err := kr.ReadMessage(ctx)
		if err != nil {
			c.Logger.WithField("err", err).Error("Failed to read message")
			panic("failed to read message " + err.Error())
		}

		c.Logger.WithField("message: ", string(msg.Value)).Info("Received message")
	}
}

// NewReaderConfig ...
func NewReaderConfig(topic string, brokers []string) kafka.ReaderConfig {
	return kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: "kafka-test-group",
	}
}
