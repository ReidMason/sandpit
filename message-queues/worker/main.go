package main

import (
	"log/slog"
	"os"
	"time"
	"worker/internal/messageHandler"
	"worker/internal/storage"

	"github.com/charmbracelet/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

const QUEUE_NAME = "ToProcess"

func main() {
	handler := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.DateTime,
	})
	logger := slog.New(handler)

	logger.Info("")
	db, err := storage.NewPostgresStorage("host=localhost port=5432 user=admin password=admin dbname=postgres sslmode=disable", logger)
	if err != nil {
		logger.Error("Failed to create storage", slog.Any("error", err))
		panic(err)
	}
	defer db.Close()

	db.GetDataset(1)

	mHandler := messageHandler.NewMessageHandler(logger, db)

	logger.Info("Worker started")

	startQueueListener(logger, mHandler)
}

func connectToQueue(logger *slog.Logger) *amqp.Connection {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		logger.Error("Failed to connect to RabbitMQ", slog.Any("error", err))
		panic(err)
	}
	return conn
}

func openChannel(conn *amqp.Connection, logger *slog.Logger) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		logger.Error("Failed to open a channel", slog.Any("error", err))
		panic(err)
	}
	return ch
}

func declareQueue(ch *amqp.Channel, logger *slog.Logger) amqp.Queue {
	q, err := ch.QueueDeclare(
		QUEUE_NAME, // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		logger.Error("Failed to declare a queue", slog.Any("error", err))
		panic(err)
	}
	return q
}

func setQos(ch *amqp.Channel, logger *slog.Logger) {
	err := ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		logger.Error("Failed to set QoS", slog.Any("error", err))
		panic(err)
	}
}

func startQueueListener(logger *slog.Logger, mHandler *messageHandler.MessageHandler) {
	conn := connectToQueue(logger)
	defer conn.Close()

	ch := openChannel(conn, logger)
	defer ch.Close()

	q := declareQueue(ch, logger)
	setQos(ch, logger)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		logger.Error("Failed to register a consumer", slog.Any("error", err))
		panic(err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			body := string(d.Body)
			logger.Info("Received a message", slog.String("body", body))
			shouldAcc, err := mHandler.HandleMessage(body)
			if err != nil {
				logger.Error("Failed to handle message", slog.Any("error", err))
			}

			d.Ack(!shouldAcc)
		}
	}()

	logger.Info("Waiting for messages...")
	<-forever
}
