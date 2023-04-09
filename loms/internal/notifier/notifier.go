package notifier

//go:generate mockery --case underscore --name Notifier --with-expecter

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	kafka "route256/libs/kafka/producer"
	"route256/loms/internal/config"
	"route256/loms/internal/model"
)

var _ Notifier = (*notifier)(nil)

type Notifier interface {
	Close() error
	SendNotification(orderID int64, status model.OrderStatus) error
}

type NotificationCallBack interface {
	MarkAsSent(ctx context.Context, orderID int64) error
	MarkAsPending(ctx context.Context, orderID int64) error
}

type notifier struct {
	producer  kafka.AsyncProducer
	callbacks NotificationCallBack
}

func NewNotifier(callbacks NotificationCallBack) (*notifier, error) {
	n := &notifier{
		callbacks: callbacks,
		producer:  nil,
	}

	p, err := kafka.NewAsyncProducer(kafka.ConfigProducer{
		ErrorCallBack:   n.OnError,
		SuccessCallBack: n.OnSuccess,
		Topic:           config.Instance.NotificationTopic,
		Brokers:         config.Instance.Brokers,
	})
	if err != nil {
		return nil, errors.Wrap(err, "create producer")
	}

	n.producer = p

	return n, nil
}

func (n *notifier) Close() error {
	return n.producer.Close()
}

func (n *notifier) OnSuccess(ctx context.Context, message *sarama.ProducerMessage) {
	orderID, err := orderIDFromMessage(message)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = n.callbacks.MarkAsSent(ctx, orderID)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func (n *notifier) OnError(ctx context.Context, message *sarama.ProducerMessage, err error) {
	log.Println(err.Error())
	orderID, err := orderIDFromMessage(message)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = n.callbacks.MarkAsPending(ctx, orderID)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func orderIDFromMessage(message *sarama.ProducerMessage) (int64, error) {
	bytes, err := message.Value.Encode()
	if err != nil {
		return 0, errors.Wrap(err, "error while encode")
	}

	var event model.StatusChangeKafka
	err = json.Unmarshal(bytes, &event)
	if err != nil {
		return 0, errors.Wrap(err, "error while unmarshal")
	}

	return event.OrderID, nil
}

func (n *notifier) SendNotification(orderID int64, status model.OrderStatus) error {
	key := strconv.FormatInt(orderID, 10)
	payload, err := json.Marshal(model.StatusChangeKafka{
		OrderID: orderID,
		Status:  int16(status),
	})
	if err != nil {
		return errors.Wrap(err, "marshal notification message")
	}

	n.producer.Send(key, payload)
	return nil
}
