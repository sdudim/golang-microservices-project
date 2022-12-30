package event

import (
	"log"

	ampq "github.com/rabbitmq/amqp091-go"
)

type Emitter struct {
	connection ampq.Connection
}

func (e *Emitter) setup() error {
	channel, err := e.connection.Channel()
	if err != nil {
		panic(err)
	}

	defer channel.Close()
	return declareExchange(channel)
}

func (e *Emitter) Push(event string, severity string) error {
	channel, err := e.connection.Channel()
	if err != nil {
		panic(err)
	}

	defer channel.Close()

	log.Println("Pushing to channetl")
	err = channel.Publish(
		"logs_topic",
		severity,
		false,
		false,
		ampq.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)

	if err != nil {
		panic(err)
	}

	return nil

}

func NewEventEmitter(conn ampq.Connection) (Emitter, error) {
	emitter := Emitter{
		connection: conn,
	}

	err := emitter.setup()
	if err != nil {
		return Emitter{}, err
	}

	return emitter, nil
}
