package amqp

import (
	"github.com/streadway/amqp"
)

// Connection ...
type Connection struct {
	URL         string
	AmqpConn    *amqp.Connection
	AmqpChannel *amqp.Channel
}

// Connect ...
func (m Connection) Connect() (*amqp.Connection, *amqp.Channel, error) {
	var (
		err         error
		conn        *amqp.Connection
		amqpChannel *amqp.Channel
	)

	conn, err = amqp.Dial(m.URL)
	if err != nil {
		return conn, amqpChannel, err
	}

	amqpChannel, err = conn.Channel()

	return conn, amqpChannel, err
}

func (m Connection) PushToQueue(payload map[string]interface{}, exchange, types string, routingKey string) (err error) {
	publisher := NewQueue(m.AmqpConn, m.AmqpChannel)

	if exchange != "" {
		_, _, err = publisher.PushQueueReconnectWithExchange(m.URL, payload, exchange, types, routingKey)
		if err != nil {
			return err
		}
	} else {
		_, _, err = publisher.PushQueueReconnect(m.URL, payload, types, routingKey)
		if err != nil {
			return err
		}
	}

	return nil
}
