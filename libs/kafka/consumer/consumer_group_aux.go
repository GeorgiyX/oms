package kafka

import (
	"go.uber.org/zap"
	"route256/libs/logger"

	"github.com/Shopify/sarama"
)

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
			if err != nil && !c.skipOnErr {
				return err // stop consumer group session, without commit
			}

			session.MarkMessage(message, "")

		case <-session.Context().Done():
			return nil // stop session, go re-balance now
		}
	}
}

func logMessageClaim(message *sarama.ConsumerMessage) {
	logger.Info("Message claimed", zap.Time("time_stamp", message.Timestamp), zap.ByteString("key", message.Key), zap.String("topic", message.Topic))
}
