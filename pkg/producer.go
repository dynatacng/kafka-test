package kappa

import (
	"context"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type producer interface {
	produce(ctx context.Context)
}

type Producer struct {
	WriterConfig kafka.WriterConfig
	Logger       *logrus.Entry
}

func (p *Producer) Produce(ctx context.Context) {
	i := 0

	kw := kafka.NewWriter(p.WriterConfig)
	for {
		km := kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("produced message id " + strconv.Itoa(i)),
		}
		err := kw.WriteMessages(ctx, km)
		if err != nil {
			p.Logger.WithField("err", err).Error("Failed to write message")
			panic("failed to write message id " + strconv.Itoa(i) + " " + err.Error())
		}
		p.Logger.WithField("Message ID: ", i).Info("Successfully wrote message")

		i++
		time.Sleep(time.Second)

	}
}

// NewWriterConfig ...
func NewWriterConfig(topic string, brokers []string) kafka.WriterConfig {
	return kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	}
}
