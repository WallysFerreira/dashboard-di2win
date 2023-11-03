package service

import "time"

type Extract struct {
	id              int
	created_at      time.Time
	pages_processed int
	doc_type        string
	user_id         int
}
