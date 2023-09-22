package main

import (
	"api/env"
	"api/utils/logger"
	"encoding/json"
	"errors"
	"log"

	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type API struct {
	upgrader websocket.Upgrader
}

type reqParamsInit struct {
	Username  string `json:"username"`
	ChannelId string `json:"channelId"`
}

func main() {
	_ = env.GetConfig()
	// cfg.ConfigSanityCheck()
	// _, err := nats.New(cfg.NATS.Url)
	// if err != nil {
	// panic(err)
	// }
	r := mux.NewRouter()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	api := &API{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024, // 1kb
			WriteBufferSize: 1024, // 1kb
			CheckOrigin: func(r *http.Request) bool {
				// return r.Header.Get("Origin") == "http://localhost:5173" // web app
				return true
			},
		},
	}

	r.HandleFunc("/ping", api.ping).Methods(http.MethodGet)
	r.HandleFunc("/connect", api.connect).Methods(http.MethodGet)

	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r))
}

type message struct {
	Text string `json:"text"`
}

func (api *API) connect(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received connection request")
	conn, err := api.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = api.getReqParams(r)
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

func (api *API) getReqParams(r *http.Request) (*reqParamsInit, error) {
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

func (api *API) ping(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received ping request...")
	w.WriteHeader(http.StatusOK)
	// set the content type to json
	w.Header().Set("Content-Type", "application/json")
	// write the json response
	w.Write([]byte(`{"message": "pong"}`))
}

func writeErr(conn *websocket.Conn, err error) {
	conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
	log.Println(err)
}
