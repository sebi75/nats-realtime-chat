package chat

import (
	"api/app/chat/agent"
	"api/app/chat/domain"
	"api/app/messageBroker"
	"api/utils"
	"api/utils/logger"
	"errors"
	"net/http"

	"github.com/gorilla/websocket"
)

type ConnectHandler struct {
	upgrader      websocket.Upgrader
	messageBroker *messageBroker.MessageBroker
	closeFunc     func()
}

func (ch ConnectHandler) Connect(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received connection request")
	conn, err := ch.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	reqParamsInit, err := ch.getReqParams(r)
	if err != nil {
		logger.Error(err.Error())
		utils.ResponseWriter(w, http.StatusBadRequest, err.Error())
		return
	}
	agent := agent.New(conn, ch.messageBroker)
	agent.HandleConnection(reqParamsInit)
}

var errConnClosed = errors.New("Websocket connection closed")

func (ch *ConnectHandler) getReqParams(r *http.Request) (*domain.ReqParamsInit, error) {
	var req *domain.ReqParamsInit

	req = &domain.ReqParamsInit{
		Username:  r.URL.Query().Get("username"),
		ChannelId: r.URL.Query().Get("channelId"),
		UUID:      r.URL.Query().Get("uuid"),
	}

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewChatHandler(mb *messageBroker.MessageBroker) (*ConnectHandler, error) {
	return &ConnectHandler{
		messageBroker: mb,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024, // 1kb
			WriteBufferSize: 1024, // 1kb
			CheckOrigin: func(r *http.Request) bool {
				// return r.Header.Get("Origin") == "http://localhost:3000" // web app
				return true
			},
		},
	}, nil
}
