package kafka

import (
	"context"
	"log"

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

type ConfigProducer struct {
	ErrorCallBack
	SuccessCallBack
	Topic   string
	Brokers []string
}

type asyncProducer struct {
	producer sarama.AsyncProducer
	topic    string
}

func NewAsyncProducer(cfg ConfigProducer) (*asyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll // exactly once
	config.Producer.Idempotent = true                // exactly once
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	cfg.SuccessCallBack = wrapSuccessCallBack(cfg.SuccessCallBack)
	cfg.ErrorCallBack = wrapErrorCallBack(cfg.ErrorCallBack)

	producer, err := sarama.NewAsyncProducer(cfg.Brokers, config)
	if err != nil {
		return nil, errors.Wrap(err, "create async producer")
	}

	go func() {
		for errIn := range producer.Errors() { // after retries
			logSendErr(errIn.Msg, errIn.Err)
			if cfg.ErrorCallBack == nil {
				return
			}
			cfg.ErrorCallBack(context.Background(), errIn.Msg, errIn.Err)
		}
	}()

	go func() {
		for message := range producer.Successes() {
			logSend(message)
			if cfg.SuccessCallBack == nil {
				return
			}
			cfg.SuccessCallBack(context.Background(), message)
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

func logSendErr(message *sarama.ProducerMessage, err error) {
	key, _ := message.Key.Encode()
	log.Printf("Message send error: timestamp = %v, key = %s, err = %s, topic = %s, partition = %s, offset = %s", message.Timestamp, string(key), err.Error(), message.Topic, message.Partition, message.Offset)
}

func logSend(message *sarama.ProducerMessage) {
	key, _ := message.Key.Encode()
	log.Printf("Message send: timestamp = %v, key = %s, topic = %s, partition = %s, offset = %s", message.Timestamp, string(key), message.Topic, message.Partition, message.Offset)
}

func wrapErrorCallBack(fn ErrorCallBack) ErrorCallBack {
	return func(ctx context.Context, message *sarama.ProducerMessage, err error) {
		if fn == nil {
			log.Println("err happen, no err handler")
		}
		fn(ctx, message, err)
	}
}

func wrapSuccessCallBack(fn SuccessCallBack) SuccessCallBack {
	return func(ctx context.Context, message *sarama.ProducerMessage) {
		if fn == nil {
			log.Println("message send, no success handler")
		}
		fn(ctx, message)
	}
}
