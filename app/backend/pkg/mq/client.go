package mq

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

var (
	emptyAMQPMessage = amqp.Publishing{}
)

type Message interface {
	Ack() error
	Decode(v interface{}) error
}

type AMQPMessage struct {
	msg amqp.Delivery
}

func NewAMQPMessage(msg amqp.Delivery) AMQPMessage {
	return AMQPMessage{
		msg: msg,
	}
}

func (m AMQPMessage) Ack() error {
	return m.msg.Ack(false)
}

func (m AMQPMessage) Decode(v interface{}) error {
	return json.Unmarshal(m.msg.Body, v)
}

// Client generic message queue client.
type Client interface {
	Close() error
	Send(msg interface{}, exchange, queue string) error
	Subscribe(queue, client string) (chan Message, error)
}

// AMQPClient client implemenation for the AMQP protocol.
type AMQPClient struct {
	url     string
	conn    *amqp.Connection
	channel *amqp.Channel
}

// NewClient creates a new MQ client.
func NewClient(conf Config) (Client, error) {
	conn, err := amqp.Dial(conf.URI())
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	client := &AMQPClient{
		url:     conf.URI(),
		conn:    conn,
		channel: ch,
	}
	return client, nil
}

func (c *AMQPClient) Subscribe(queue, client string) (chan Message, error) {
	deliveryChan, err := c.channel.Consume(queue, client, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	messageChannel := make(chan Message)
	go func() {
		for delivery := range deliveryChan {
			messageChannel <- NewAMQPMessage(delivery)
		}
		close(messageChannel)
	}()

	return messageChannel, nil
}

// Send serializes and sends a message to the specified queue.
func (c *AMQPClient) Send(msg interface{}, exchange, queue string) error {
	amqpMessage, err := newAMQPMessage(msg)
	if err != nil {
		return err
	}

	return c.channel.Publish(exchange, queue, false, false, amqpMessage)
}

func newAMQPMessage(msg interface{}) (amqp.Publishing, error) {
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return emptyAMQPMessage, err
	}

	amqpMessage := amqp.Publishing{
		ContentType: "application/json",
		Body:        msgBody,
	}
	return amqpMessage, nil
}

// Close closes the underlying connection to the mq server.
func (c *AMQPClient) Close() error {
	err := c.channel.Close()
	if err != nil {
		log.Printf("Failed to close channel: %s", err.Error())
	}
	return c.conn.Close()
}
