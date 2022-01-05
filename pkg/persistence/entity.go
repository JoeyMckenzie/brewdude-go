package persistence

import "time"

type BaseAuditableEntity struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
}
