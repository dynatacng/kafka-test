package kappa

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type consumer interface {
	consume(ctx context.Context)
}

// Consumer ...
type Consumer struct {
	ReaderConfig kafka.ReaderConfig
	Logger       *logrus.Entry
}

// Consume ...
func (c *Consumer) Consume(ctx context.Context) {

	kr := kafka.NewReader(c.ReaderConfig)
	for {
		msg, err := kr.ReadMessage(ctx)
		if err != nil {
			c.Logger.WithField("err", err).Error("Failed to read message")
			panic("failed to read message " + err.Error())
		}

		c.Logger.WithField("message: ", string(msg.Value)).WithField("time read: ", time.Now()).Info("Received message")
	}
}

// NewReaderConfig ...
func NewReaderConfig(topic string, brokers []string) kafka.ReaderConfig {
	// logger := log.New(os.Stdout, "kafka consumer: ", 0)
	return kafka.ReaderConfig{
		// Logger:  logger,
		Brokers:     brokers,
		Topic:       topic,
		GroupID:     "kafka-test-group",
		MinBytes:    5,                 // min # of bytes to be read
		MaxBytes:    1e4,               // max # of bytes to be read
		MaxWait:     3 * time.Second,   // wait at most 3 secs before polling for new data
		StartOffset: kafka.FirstOffset, // consume earliest messages available
		// StartOffset: kafka.LastOffset, // consume the newest messages available after startup
	}
}
