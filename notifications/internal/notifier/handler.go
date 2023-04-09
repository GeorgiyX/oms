package notifier

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"route256/notifications/internal/model"
)

func HandleNotification(ctx context.Context, message *sarama.ConsumerMessage) error {
	var event model.StatusChangeKafka
	err := json.Unmarshal(message.Value, &event)
	if err != nil {
		log.Println("error while unmarshal")
		return nil
	}

	log.Printf("new notification: time=%s, orderID = %d, status = %d", time.Now(), event.OrderID, event.Status)
	return nil
}
