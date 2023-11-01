package utils

import (
	"api/app/chat/domain"
	"database/sql"
	"time"

	"github.com/vmihailenco/msgpack"
)

func NullStringToPtr(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	} else {
		return nil
	}
}

func NullTimeToPtr(s sql.NullTime) *time.Time {
	if s.Valid {
		return &s.Time
	} else {
		return nil
	}
}

func DecodeMessage(messageBinaryData []byte) (*domain.Message, error) {
	var message *domain.Message
	err := msgpack.Unmarshal(messageBinaryData, &message)
	if err != nil {
		return nil, err
	}
	return message, nil
}
