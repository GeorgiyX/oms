package kafka

import (
	"context"
	"log"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

type MessageHandler func(ctx context.Context, message *sarama.ConsumerMessage) error

var _ ConsumerGroup = (*consumerGroup)(nil)

type ConsumerGroup interface {
	Cancel() error
}

type consumerGroup struct {
	client     sarama.ConsumerGroup
	clientOnce sync.Once
	cancel     context.CancelFunc
	cancelOnce sync.Once
	wg         *sync.WaitGroup
}

type ConfigConsumer struct {
	InitialOffset int64
	Topic         string
	Group         string
	Brokers       []string
	Handler       MessageHandler
	SkipOnErr     bool
}

func NewConsumerGroup(cfg ConfigConsumer) (*consumerGroup, error) {
	config := sarama.NewConfig()
	config.Version = sarama.MaxVersion
	config.Consumer.Offsets.Initial = cfg.InitialOffset
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

	c := consumer{
		ready:     make(chan bool),
		skipOnErr: cfg.SkipOnErr,
		handler:   wrapNil(cfg.Handler),
	}

	ctx, cancel := context.WithCancel(context.Background())

	client, err := sarama.NewConsumerGroup(cfg.Brokers, cfg.Group, config)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "creating consumer group client")
	}

	cg := &consumerGroup{
		client:     client,
		clientOnce: sync.Once{},
		cancel:     cancel,
		cancelOnce: sync.Once{},
		wg:         &sync.WaitGroup{},
	}

	cg.wg.Add(1)
	go func() {
		defer cg.wg.Done()
		for {
			errIn := client.Consume(ctx, cfg.Brokers, &c) // start consume session
			if errIn != nil {
				log.Print("err while call client.Consume")
			}

			if ctx.Err() != nil { // check if context was cancelled
				return
			}
			c.ready = make(chan bool)
		}
	}()

	<-c.ready // await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")

	return cg, nil
}

func (c *consumerGroup) Cancel() (err error) {
	c.cancelOnce.Do(func() {
		c.cancel()
	})
	c.wg.Wait()
	c.clientOnce.Do(func() {
		err = c.client.Close()
	})
	return err
}

func wrapNil(fn MessageHandler) MessageHandler {
	return func(ctx context.Context, message *sarama.ConsumerMessage) error {
		if fn == nil {
			log.Println("new message, no handler")
			return nil
		}
		return fn(ctx, message)
	}
}
