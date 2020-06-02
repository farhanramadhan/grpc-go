package messages

import "time"

type Message struct {
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
}
