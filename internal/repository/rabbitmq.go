package repository

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
)

type BrokerRepository interface {
	Publish(ctx context.Context, exchange, key string, payload any) error
	ConsumeWithQos(ctx context.Context, queue, consumer string, prefetch int) (<-chan amqp.Delivery, error)
}

type rabbitmqRepository struct {
	ch *amqp.Channel
}

func NewRabbitmqRepository(ch *amqp.Channel) BrokerRepository {
	return &rabbitmqRepository{ch: ch}
}

func (r *rabbitmqRepository) Publish(ctx context.Context, exchange, key string, payload any) error {
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	err = r.ch.PublishWithContext(
		ctx,
		exchange,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *rabbitmqRepository) ConsumeWithQos(ctx context.Context, queue, consumer string, prefetch int) (<-chan amqp.Delivery, error) {
	messages, err := r.ch.ConsumeWithContext(
		ctx,
		queue,
		consumer,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	err = r.ch.Qos(prefetch, 0, false)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
