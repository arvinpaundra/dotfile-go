package broker

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type (
	brokerConfig struct {
		user     string
		password string
		host     string
		port     string
	}

	rabbitmqConfig struct {
		brokerConfig
		vhost string
	}
)

func (conf *rabbitmqConfig) connect() {
	url := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/%s",
		conf.user,
		conf.password,
		conf.host,
		conf.port,
		conf.vhost,
	)

	var err error

	conn, err = amqp.Dial(url)
	if err != nil {
		log.Fatalf("failed connect to rabbitmq: %s", err.Error())
	}

	ch, err = conn.Channel()
	if err != nil {
		log.Fatalf("failed to create channel: %s", err.Error())
	}
}
