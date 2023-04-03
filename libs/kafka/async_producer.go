package kafka

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

var _ AsyncProducer = (*asyncProducer)(nil)

type AsyncProducer interface {
	Send(key string, message []byte)
	Close() error
}

type ErrorCallBack func(ctx context.Context, message *sarama.ProducerMessage, err error)
type SuccessCallBack func(ctx context.Context, message *sarama.ProducerMessage)
type Closer func() error

type Config struct {
	ErrorCallBack
	SuccessCallBack
	Topic   string
	Brokers []string
}

type asyncProducer struct {
	producer sarama.AsyncProducer
	topic    string
}

func NewAsyncProducer(cfg Config) (*asyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll // exactly once
	config.Producer.Idempotent = true                // exactly once
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewAsyncProducer(cfg.Brokers, config)
	if err != nil {
		return nil, errors.Wrap(err, "create async producer")
	}

	go func() {
		for errIn := range producer.Errors() { // after retries
			if cfg.ErrorCallBack == nil {
				return
			}
			cfg.ErrorCallBack(context.Background(), errIn.Msg, errIn.Err)
		}
	}()

	go func() {
		for msg := range producer.Successes() {
			if cfg.SuccessCallBack == nil {
				return
			}
			cfg.SuccessCallBack(context.Background(), msg)
		}
	}()

	return &asyncProducer{
		producer: producer,
		topic:    cfg.Topic,
	}, nil
}

func (a *asyncProducer) Send(key string, message []byte) {
	a.producer.Input() <- &sarama.ProducerMessage{
		Topic:     a.topic,
		Key:       sarama.StringEncoder(key),
		Value:     sarama.ByteEncoder(message),
		Partition: -1,
	}
}

func (a *asyncProducer) Close() error {
	return a.producer.Close()
}
