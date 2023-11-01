package chat

import (
	"api/app/chat/domain"
	"api/app/messageBroker"
	"api/utils/logger"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type ConnectHandler struct {
	upgrader      websocket.Upgrader
	messageBroker *messageBroker.MessageBroker
}

func (ch ConnectHandler) Connect(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received connection request")
	conn, err := ch.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// reqParamsInit, err := ch.getReqParams(r)
	if err != nil {
		if err == errConnClosed {
			return
		}
		writeErr(conn, err)
		return
	}

	conn.SetCloseHandler(func(code int, text string) error {
		logger.Info("Connection closed")
		return nil
	})

	defer conn.Close()
	// messageChan := make(chan domain.Message)
	{
		// set up the nats subscription
		// chatSubject := "chat." + reqParamsInit.ChannelId
		// close, err := ch.natsClient.Subscribe(chatSubject)
	}
	for {
		_, reader, err := conn.NextReader()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				return
			}
			writeErr(conn, err)
		}

		logger.Info("Received message...")
		var bigMessage struct {
			Type string          `json:"type"`
			Data json.RawMessage `json:"data,omitempty"`
		}

		err = json.NewDecoder(reader).Decode(&bigMessage)
		if err != nil {
			writeErr(conn, err)
			return
		}

		switch bigMessage.Type {
		case "chatMsg":
			var msg domain.Message

			err := json.Unmarshal(bigMessage.Data, &msg)
			if err != nil {
				writeErr(conn, err)
				return
			}

			logger.Info("Received message", zap.String("text", msg.Text))

			err = conn.WriteJSON(json.RawMessage(`{"message": "Hello from the server"}`))
			if err != nil {
				writeErr(conn, err)
				return
			}
		}
	}
}

var errConnClosed = errors.New("Websocket connection closed")

func (ch *ConnectHandler) getReqParams(r *http.Request) (*domain.ReqParamsInit, error) {
	var req *domain.ReqParamsInit

	req = &domain.ReqParamsInit{
		Username:  r.URL.Query().Get("username"),
		ChannelId: r.URL.Query().Get("channelId"),
	}

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	return req, nil
}

func writeErr(conn *websocket.Conn, err error) {
	conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
	logger.Error(err.Error())
}

func NewChatHandler(mb *messageBroker.MessageBroker) (*ConnectHandler, error) {
	return &ConnectHandler{
		messageBroker: mb,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024, // 1kb
			WriteBufferSize: 1024, // 1kb
			CheckOrigin: func(r *http.Request) bool {
				// return r.Header.Get("Origin") == "http://localhost:5173" // web app
				return true
			},
		},
	}, nil
}
