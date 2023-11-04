package agent

import (
	"api/app/chat/domain"
	"api/app/messageBroker"
	"api/utils/logger"
	"encoding/json"
	"io"

	"github.com/gorilla/websocket"
)

type msgT int

// Agent is a websocket connection self-contained
// Which handlers receiving, sending and validation of messages
type Agent struct {
	conn          *websocket.Conn
	messageBroker *messageBroker.MessageBroker
	Uuid          string
	channelUUID   string
	closeSub      func()
	closed        bool
}

func New(conn *websocket.Conn, messageBroker *messageBroker.MessageBroker) *Agent {
	return &Agent{
		conn:          conn,
		messageBroker: messageBroker,
	}
}

type msg struct {
	Type  msgT        `json:"type"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (a *Agent) HandleConnection(reqParamsInit *domain.ReqParamsInit) {
	a.conn.SetCloseHandler(func(code int, text string) error {
		a.closed = true
		return nil
	})
	a.channelUUID = reqParamsInit.ChannelId
	a.Uuid = reqParamsInit.UUID

	messageChan := make(chan *domain.Message)
	{
		// set up the nats subscription
		var close func()
		chatSubject := "chat." + reqParamsInit.ChannelId
		close, err := a.messageBroker.Subscribe(chatSubject, reqParamsInit.UUID, messageChan)
		if err != nil {
			writeErr(a.conn, "Error subscribing to the chat subject")
			return
		}
		a.closeSub = close
	}
	// handle the websocket connection and messages
	a.loop(messageChan)
}

func (a *Agent) loop(m chan *domain.Message) {
	// one channel for receiving messages from the nats server
	// a.k.a messages from other users
	go func() {
		for {
			if a.closed {
				return
			}
			_, reader, err := a.conn.NextReader()
			// a.conn.ReadMessage()
			if err != nil {
				logger.Error(err.Error())
				writeErr(a.conn, err.Error())
				continue
			}

			a.handleClientMessage(reader)
		}
	}()

	// and one channel for the messages that the user sends directly to the server
	go func() {
		// close the nats subscription
		defer a.closeSub()
		// close the websocket connection
		defer a.conn.Close()
		for {
			select {
			case message := <-m:
				err := a.conn.WriteJSON(message)
				if err != nil {
					writeErr(a.conn, "Error sending the message")
					return
				}
			}
		}
	}()
}

func (a *Agent) handleClientMessage(reader io.Reader) {
	// handle message sent by the client
	// send message to the nats server to save it in chat history and send it to the other
	// subscribed users
	logger.Info("Received message...")
	var bigMessage struct {
		Type string          `json:"type"`
		Data json.RawMessage `json:"data,omitempty"`
	}

	err := json.NewDecoder(reader).Decode(&bigMessage)
	if err != nil {
		writeErr(a.conn, "Error decoding the message")
		return
	}

	switch bigMessage.Type {
	case "chatMsg":
		var msg domain.Message
		err := json.Unmarshal(bigMessage.Data, &msg)
		if err != nil {
			writeErr(a.conn, "Error unmarshalling the message")
			return
		}
		msg.FromUUID = a.Uuid
		a.handleChatMessage(&msg)
	}
}

func (a *Agent) handleChatMessage(message *domain.Message) {
	// do validation on the message here
	subject := "chat." + a.channelUUID
	a.messageBroker.Send(subject, message)
}

func writeErr(conn *websocket.Conn, err string) {
	conn.WriteJSON(msg{Error: err, Type: 5})
}

func writeFatal(conn *websocket.Conn, err string) {
	conn.WriteJSON(msg{Error: err, Type: 5})
	conn.Close()
}
