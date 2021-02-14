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

// Producer ...
type Producer struct {
	WriterConfig kafka.WriterConfig
	Logger       *logrus.Entry
}

// Produce ...
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
		p.Logger.WithField("Message ID: ", i).WithField("time wrote: ", time.Now()).Info("Successfully wrote message")

		i++
		//time.Sleep(500 * time.Millisecond)

	}
}

// NewWriterConfig ...
func NewWriterConfig(topic string, brokers []string) kafka.WriterConfig {
	// logger := log.New(os.Stdout, "kafka producer: ", 0)
	return kafka.WriterConfig{
		// Logger:       logger,
		Brokers:      brokers,
		Topic:        topic,
		BatchSize:    5,               // wait until 5 messages before writing
		BatchTimeout: 5 * time.Second, // after 4 seconds, write all pending messages
		RequiredAcks: -1,
	}
	/*
		Required Acks -- Required Acknowledgements
		-1 = All brokers received the message
		0 =  No brokers acknowledge receiving the message
		1 = Leader broker acknowledges receiving the message
	*/
}
