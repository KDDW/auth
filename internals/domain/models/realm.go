package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Realm struct {
	bun.BaseModel `bun:"table:realms,alias:r"`

	ID        int64     `bun:",pk,autoincrement" json:"id"`
	Code      string    `json:"code"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
}
