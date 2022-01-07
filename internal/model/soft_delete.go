package model

import (
	"github.com/lib/pq"
)

type SoftDelete struct {
	IsDeleted bool        `json:"is_deleted" db:"is_deleted"`
	DeletedAt pq.NullTime `json:"deleted_at" db:"deleted_at"`
}
