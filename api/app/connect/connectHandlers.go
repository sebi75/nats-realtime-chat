package connect

import (
	"api/utils/logger"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type message struct {
	Text string `json:"text"`
}

type reqParamsInit struct {
	Username  string `json:"username"`
	ChannelId string `json:"channelId"`
}

type ConnectHandler struct {
	upgrader websocket.Upgrader
}

func (ch ConnectHandler) Connect(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received connection request")
	conn, err := ch.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = ch.getReqParams(r)
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
			var msg message

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

func (reqInitParams *reqParamsInit) Validate() error {
	if reqInitParams.Username == "" {
		return errors.New("Username is required")
	}

	if reqInitParams.ChannelId == "" {
		return errors.New("ChannelId is required")
	}

	return nil
}

func (ch *ConnectHandler) getReqParams(r *http.Request) (*reqParamsInit, error) {
	var req *reqParamsInit

	req = &reqParamsInit{
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

func NewConnectHandler() (*ConnectHandler, error) {
	return &ConnectHandler{
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
