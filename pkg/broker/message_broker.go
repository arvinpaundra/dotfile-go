package broker

import (
	"log"
	"sync"

	config "github.com/arvinpaundra/dotfile-go/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

type (
	ExchangeKind string

	QueueConfig struct {
		Name       string
		Durable    bool
		AutoDelete bool
		Exclusive  bool
		NoWait     bool
		Args       amqp.Table
	}

	ExchangeConfig struct {
		Name       string
		Kind       ExchangeKind
		Durable    bool
		AutoDelete bool
		Internal   bool
		NoWait     bool
		Args       amqp.Table
	}
)

var (
	conn *amqp.Connection
	ch   *amqp.Channel
	once sync.Once

	ExchangeDirect ExchangeKind = "direct"
	ExchangeFanout ExchangeKind = "fanout"
	ExchangeTopic  ExchangeKind = "topic"
)

func createConnection() {
	conf := brokerConfig{
		user:     config.C.Rabbitmq.User,
		password: config.C.Rabbitmq.Password,
		host:     config.C.Rabbitmq.Host,
		port:     config.C.Rabbitmq.Port,
	}

	rabbitmq := rabbitmqConfig{
		brokerConfig: conf,
		vhost:        config.C.Rabbitmq.Vhost,
	}

	once.Do(func() {
		rabbitmq.connect()
	})

	log.Println("connected to rabbitmq")
}

func GetBrokerInstance() (*amqp.Connection, *amqp.Channel) {
	if conn == nil || ch == nil {
		createConnection()
	}

	return conn, ch
}

func Close(conn *amqp.Connection) error {
	err := conn.Close()
	if err != nil {
		log.Printf("failed to close rabbit connection: %s", err.Error())
		return err
	}

	return nil
}

func DeclareQueue(ch *amqp.Channel, conf QueueConfig) (amqp.Queue, error) {
	queue, err := ch.QueueDeclare(
		conf.Name,
		conf.Durable,
		conf.AutoDelete,
		conf.Exclusive,
		conf.NoWait,
		conf.Args,
	)

	if err != nil {
		return amqp.Queue{}, err
	}

	return queue, nil
}

func DeclareExchange(ch *amqp.Channel, conf ExchangeConfig) (string, error) {
	err := ch.ExchangeDeclare(
		conf.Name,
		string(conf.Kind),
		conf.Durable,
		conf.AutoDelete,
		conf.Internal,
		conf.NoWait,
		conf.Args,
	)

	if err != nil {
		return "", err
	}

	return conf.Name, nil
}

func BindExchangeToQueue(ch *amqp.Channel, queueName, exchangeName, routeKey string) error {
	err := ch.QueueBind(queueName, routeKey, exchangeName, false, nil)
	if err != nil {
		return err
	}

	return nil
}
