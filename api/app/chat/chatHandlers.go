package chat

import (
	"api/app/auth"
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
	authService   *auth.AuthService
	messageBroker *messageBroker.MessageBroker
	closeFunc     func()
}

func (ch ConnectHandler) Connect(w http.ResponseWriter, r *http.Request) {
	reqParamsInit, err := ch.getReqParams(r)
	if err != nil {
		logger.Error(err.Error())
		utils.ResponseWriter(w, http.StatusBadRequest, err.Error())
		return
	}
	_, appErr := ch.authService.Verify(reqParamsInit.Token)
	if appErr != nil {
		utils.ResponseWriter(w, appErr.Code, appErr.Message)
		return
	}
	conn, err := ch.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	agent := agent.New(conn, ch.messageBroker)
	agent.Uuid = reqParamsInit.UUID
	agent.HandleConnection(reqParamsInit)
}

var errConnClosed = errors.New("Websocket connection closed")

func (ch *ConnectHandler) getReqParams(r *http.Request) (*domain.ReqParamsInit, error) {
	var req *domain.ReqParamsInit

	req = &domain.ReqParamsInit{
		Username:  r.URL.Query().Get("username"),
		ChannelId: r.URL.Query().Get("channelId"),
		UUID:      r.URL.Query().Get("uuid"),
		Token:     r.URL.Query().Get("token"),
	}

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewChatHandler(mb *messageBroker.MessageBroker, as *auth.AuthService) (*ConnectHandler, error) {
	return &ConnectHandler{
		messageBroker: mb,
		authService:   as,
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
