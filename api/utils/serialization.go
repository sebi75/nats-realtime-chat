package utils

import (
	"database/sql"
	"time"
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
