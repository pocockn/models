package sns

import (
	"encoding/base64"
	"encoding/json"
	"time"
)

// Notification events are sent for messages that are published to the SNS
// topic.
type Notification struct {
	MessageID      string    `json:"MessageId"`
	TopicARN       string    `json:"TopicArn"`
	Subject        string    `json:"Subject"`
	Message        string    `json:"Message"`
	Timestamp      time.Time `json:"Timestamp"`
	UnsubscribeURL string    `json:"UnsubscribeURL"`

	// MessageAttributes contain any attributes added to the message when
	// publishing it to SNS. This is most commonly used when transmitting binary
	// date (using raw message delivery).
	MessageAttributes map[string]MessageAttribute `json:"MessageAttributes"`
}

// NewNotification creates a new notification struct.
func NewNotification(data json.RawMessage) (Notification, error) {
	var notification Notification
	err := json.Unmarshal(data, &notification)
	return notification, err
}

type MessageAttribute struct {
	Type  string `json:"Type"`
	Value string `json:"Value"`
}

func (m MessageAttribute) StringValue() string {
	return m.Value
}

func (m MessageAttribute) BinaryValue() ([]byte, error) {
	return base64.StdEncoding.DecodeString(m.Value)
}
