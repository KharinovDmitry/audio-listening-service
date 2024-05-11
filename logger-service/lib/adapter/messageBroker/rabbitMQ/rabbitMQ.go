package rabbitMQ

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"logger-service/lib/adapter/messageBroker"
)

type rabbitMQAdapter struct {
	conn      *amqp.Connection
	rbChannel *amqp.Channel
	logsQueue *amqp.Queue
}

func NewRabbitMQAdapter() messageBroker.MessageBroker {
	return &rabbitMQAdapter{}
}

func (r *rabbitMQAdapter) Connect(url string) error {
	conn, err := amqp.Dial(url)
	if err != nil {
		return err
	}
	r.conn = conn

	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}
	r.rbChannel = ch

	logsQueue, err := r.rbChannel.QueueDeclare(
		messageBroker.LogsQueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	r.logsQueue = &logsQueue

	return nil
}

func (r *rabbitMQAdapter) Close() error {
	if err := r.rbChannel.Close(); err != nil {
		return err
	}
	if err := r.conn.Close(); err != nil {
		return err
	}
	return nil
}

func (r *rabbitMQAdapter) Publish(message messageBroker.LogMessage) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = r.rbChannel.Publish(
		"",
		messageBroker.LogsQueueName,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         data,
		})
	if err != nil {
		return err
	}

	return nil
}

func (r *rabbitMQAdapter) Subscribe(queue string) (<-chan messageBroker.LogMessage, error) {
	q, err := r.rbChannel.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	messages, err := r.rbChannel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	resMessageChannel := make(chan messageBroker.LogMessage)
	go func() {
		for message := range messages {
			var msg messageBroker.LogMessage
			err = json.Unmarshal(message.Body, &msg)
			if err != nil {
				log.Println(err.Error())
			}
			resMessageChannel <- msg
		}
	}()
	return resMessageChannel, nil
}
