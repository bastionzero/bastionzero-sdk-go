package service

import (
	"fmt"
	"time"
)

// Timestamp represents a time that when marshalled into JSON converts to
// RFC3339 with a UTC timezone. This is necessary because the BastionZero API
// expects times to be in UTC for some APIs (e.g. TimeExpires in policy). All
// exported methods of time.Time can be called on Timestamp.
type Timestamp struct {
	time.Time
}

func (t Timestamp) String() string {
	return t.Time.String()
}

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.Time.UTC().Format(time.RFC3339))
	return []byte(stamp), nil
}

// Equal reports whether t and u are equal based on time.Equal
func (t Timestamp) Equal(u Timestamp) bool {
	return t.Time.Equal(u.Time)
}
