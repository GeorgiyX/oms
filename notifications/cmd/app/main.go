package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	kafka "route256/libs/kafka/consumer"
	"route256/notifications/internal/config"
	"route256/notifications/internal/notifier"
)

func main() {
	consumer, err := kafka.NewConsumerGroup(kafka.ConfigConsumer{
		InitialOffset: kafka.OffsetOldest,
		Topic:         config.Instance.NotificationTopic,
		Group:         "notification-consumer-group",
		Brokers:       config.Instance.Brokers,
		Handler:       notifier.HandleNotification,
		SkipOnErr:     false,
	})

	if err != nil {
		log.Fatalf("fail create consumer group: %v\n", err)
	}

	defer consumer.Cancel()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	<-sigterm
	log.Println("terminating: via signal")
}
