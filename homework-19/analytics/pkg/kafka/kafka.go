package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/segmentio/kafka-go"
)

type Client struct {
	Reader      *kafka.Reader
	Writer      *kafka.Writer
	subscribers []func(msg string, short string)
}

func New(brokers []string, topic string, groupId string) (*Client, error) {
	if len(brokers) == 0 || brokers[0] == "" || topic == "" || groupId == "" {
		return nil, errors.New("не указаны параметры подключения к Kafka")
	}

	c := &Client{
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  brokers,
			Topic:    topic,
			GroupID:  groupId,
			MinBytes: 10e1,
			MaxBytes: 10e6,
		}),
		Writer: &kafka.Writer{
			Addr:                   kafka.TCP(brokers[0]),
			Topic:                  topic,
			Balancer:               &kafka.LeastBytes{},
			AllowAutoTopicCreation: true,
		},
	}

	return c, nil
}

func (c *Client) Subscribe(callback func(msg string, short string)) {
	c.subscribers = append(c.subscribers, callback)
}

func (c *Client) Consumer() {
	for {
		msg, err := c.Reader.FetchMessage(context.Background())
		if err != nil {
			log.Println(err)
		}

		var val map[string]string
		err = json.Unmarshal(msg.Value, &val)
		if err != nil {
			log.Println(err)
		}

		for _, subscriber := range c.subscribers {
			subscriber(val["orig"], val["short"])
		}

		err = c.Reader.CommitMessages(context.Background(), msg)
		if err != nil {
			log.Println(err)
		}
	}
}
