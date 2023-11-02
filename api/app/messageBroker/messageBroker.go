package messageBroker

import (
	"api/app/chat/domain"
	"api/pkg/nats"
	"api/utils"
	"api/utils/logger"

	"go.uber.org/zap"
)

// MessageBroker is a wrapper around the nats client which adds
// additional functionality and abstraction
type MessageBroker struct {
	natsClient *nats.Client
}

func (mb *MessageBroker) Subscribe(subject string, uuid string, c chan *domain.Message) (func(), error) {
	closer, err := mb.natsClient.Subscribe(subject, func(message []byte) {
		// parse the message data so that when we send it through the channel
		// we are sending a domain.Message object
		messageData, err := utils.DecodeMessage(message)
		if err != nil {
			messageData = &domain.Message{
				Text:     "message unavailable, decoding error",
				FromUUID: "broker",
			}
		}
		logger.Info("Received message in NATS broker:", zap.Any("message", messageData))
		// check that the newly parsed message has a fromUUID different than the
		// uuid we are subscribing with ( don't send messages back to the sender )
		if messageData.FromUUID != uuid {
			c <- messageData
		} else {
			// do something with the message
		}
	})

	if err != nil {
		logger.Error("Error subscribing to the NATS broker", zap.Error(err))
		return nil, err
	}
	logger.Info("Subscribed to the NATS broker")

	return func() {
		closer.Unsubscribe()
	}, nil
}

func (mb *MessageBroker) Send(subject string, message *domain.Message) error {
	encodedMsg, err := message.Encode()
	if err != nil {
		return err
	}
	return mb.natsClient.Publish(subject, encodedMsg)
}

func New(natsClient *nats.Client) *MessageBroker {
	return &MessageBroker{
		natsClient: natsClient,
	}
}
