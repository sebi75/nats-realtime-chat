package nats

import (
	"fmt"

	nats "github.com/nats-io/nats.go"
)

type Client struct {
	conn *nats.Conn
}

func New(url string) (*Client, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to the nats client: %v", err)
	}

	return &Client{conn: conn}, nil
}

func (c *Client) Subscribe(subject string, callback func(message []byte)) (*nats.Subscription, error) {
	// do other message formatting here
	// before sending it to the nats server if neccessary
	return c.conn.Subscribe(
		subject,
		func(msg *nats.Msg) {
			callback(msg.Data)
		},
	)
}

func (c *Client) Publish(subject string, message []byte) error {
	return c.conn.Publish(subject, message)
}
