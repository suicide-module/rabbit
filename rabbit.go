package rabbit

import (
	"log"

	"github.com/wagslane/go-rabbitmq"
)

func PublishData(body []byte) {
	publisher, err := rabbitmq.NewPublisher(
		FullHost,
		rabbitmq.Config{},
		rabbitmq.WithPublisherOptionsLogging,
	)
	defer publisher.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = publisher.Publish(
		[]byte(body),
		[]string{"onyore"},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsMandatory,
		rabbitmq.WithPublishOptionsPersistentDelivery,
		rabbitmq.WithPublishOptionsExchange("amq.direct"),
	)
	if err != nil {
		log.Fatal(err)
	}

	returns := publisher.NotifyReturn()
	go func() {
		for r := range returns {
			log.Printf("message returned from server: %s", string(r.Body))
		}
	}()
}
