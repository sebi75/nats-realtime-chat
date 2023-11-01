package domain

import (
	"errors"

	"github.com/vmihailenco/msgpack"
)

type Message struct {
	FromUUID string            `json:"fromUuid"`
	Meta     map[string]string `json:"meta"`
	Time     int64             `json:"time"`
	Text     string            `json:"text"`
}

func (m *Message) Encode() ([]byte, error) {
	return msgpack.Marshal(m)
}

type ReqParamsInit struct {
	Username  string `json:"username"`
	ChannelId string `json:"channelId"`
}

func (reqInitParams *ReqParamsInit) Validate() error {
	if reqInitParams.Username == "" {
		return errors.New("Username is required")
	}

	if reqInitParams.ChannelId == "" {
		return errors.New("ChannelId is required")
	}

	return nil
}
