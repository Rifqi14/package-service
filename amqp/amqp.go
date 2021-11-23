package amqp

import (
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
)

// IQueue ...
type IQueue interface {
	PushQueue(data map[string]interface{}, types string) error
	PushQueueReconnect(url string, data map[string]interface{}, types, routingKey string) (*amqp.Connection, *amqp.Channel, error)
	PushQueueReconnectWithExchange(url string, data map[string]interface{}, exchange, types, routingKey string) (*amqp.Connection, *amqp.Channel, error)
}

const (
	MailExchange = "mail.exchange"
	//mail incoming
	MailIncoming = "mail.incoming"
	//mail deadletter
	MailDeadLetter = "mail.deadLetter"

	EchoExchange   = "echo.exchange"
	EchoIncoming   = "echo.incoming"
	EchoDeadLetter = "echo.deadLetter"
)

// queue ...
type queue struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// NewQueue ...
func NewQueue(conn *amqp.Connection, channel *amqp.Channel) IQueue {
	return &queue{
		Connection: conn,
		Channel:    channel,
	}
}

// PushQueue ...
func (m queue) PushQueue(data map[string]interface{}, types string) error {
	queue, err := m.Channel.QueueDeclare(types, true, false, false, false, nil)
	if err != nil {
		return err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = m.Channel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	return err
}

// PushQueueReconnect ...
func (m queue) PushQueueReconnect(url string, data map[string]interface{}, types, routingKey string) (*amqp.Connection, *amqp.Channel, error) {
	if m.Connection != nil {
		if m.Connection.IsClosed() {
			c := Connection{
				URL: url,
			}
			newConn, newChannel, err := c.Connect()
			if err != nil {
				return nil, nil, err
			}
			m.Connection = newConn
			m.Channel = newChannel
		}
	} else {
		c := Connection{
			URL: url,
		}
		newConn, newChannel, err := c.Connect()
		if err != nil {
			return nil, nil, err
		}
		m.Connection = newConn
		m.Channel = newChannel
	}

	args := amqp.Table{
		"x-dead-letter-exchange":    "",
		"x-dead-letter-routing-key": routingKey,
	}
	queue, err := m.Channel.QueueDeclare(types, true, false, false, false, args)
	if err != nil {
		return nil, nil, err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, nil, nil
	}

	err = m.Channel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})
	if err != nil {
		return nil, nil, errors.New("itu lah")
	}

	return m.Connection, m.Channel, err
}

func (m queue) PushQueueReconnectWithExchange(url string, data map[string]interface{}, exchange, types, routingKey string) (*amqp.Connection, *amqp.Channel, error) {
	if m.Connection != nil {
		if m.Connection.IsClosed() {
			c := Connection{
				URL: url,
			}
			newConn, newChannel, err := c.Connect()
			if err != nil {
				return nil, nil, err
			}
			m.Connection = newConn
			m.Channel = newChannel
		}
	} else {
		c := Connection{
			URL: url,
		}
		newConn, newChannel, err := c.Connect()
		if err != nil {
			return nil, nil, err
		}
		m.Connection = newConn
		m.Channel = newChannel
	}

	args := amqp.Table{
		"x-dead-letter-exchange":    exchange,
		"x-dead-letter-routing-key": routingKey,
	}
	queue, err := m.Channel.QueueDeclare(types, true, false, false, false, args)
	if err != nil {
		return nil, nil, err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, nil, nil
	}

	err = m.Channel.Publish(exchange, queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})
	if err != nil {
		return nil, nil, err
	}

	return m.Connection, m.Channel, err
}
