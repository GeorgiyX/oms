package kafka

import (
	"context"
	"log"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

type MessageHandler func(ctx context.Context, message *sarama.ConsumerMessage) error

type consumer struct {
	ready     chan bool
	skipOnErr bool
	handler   MessageHandler
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (c *consumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready) // Mark the consumer as ready
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (c *consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (c *consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				return nil
			}
			logMessageClaim(message)

			err := c.handler(session.Context(), message)
			if err != nil && !c.skipOnErr { //
				return err // stop consumer group session, without commit
			}

			session.MarkMessage(message, "")

		case <-session.Context().Done():
			return nil // stop session, go re-balance now
		}
	}
}

func logMessageClaim(message *sarama.ConsumerMessage) {
	log.Printf("Message claimed: timestamp = %v, key = %s, topic = %s", message.Timestamp, string(message.Key), message.Topic)
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
}

func NewConsumerGroup(cfg ConfigConsumer) (*consumerGroup, error) {
	config := sarama.NewConfig()
	config.Version = sarama.MaxVersion
	config.Consumer.Offsets.Initial = cfg.InitialOffset
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

	c := consumer{
		ready:     make(chan bool),
		skipOnErr: false,
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
