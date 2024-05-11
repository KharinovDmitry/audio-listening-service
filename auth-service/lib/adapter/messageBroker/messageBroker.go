package messageBroker

import "time"

var (
	LogsQueueName = "logs_queue"
)

type LogMessage struct {
	AppName string    `json:"appName"`
	Level   int       `json:"level"`
	Time    time.Time `json:"time"`
	Text    string    `json:"text"`
}

type MessageBroker interface {
	Connect(url string) error
	Publish(message LogMessage) error
	Subscribe(queue string) (<-chan LogMessage, error)
	Close() error
}
