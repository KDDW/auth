package models

import (
	"fmt"
	"time"

	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID        int64     `bun:",pk,autoincrement" json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	RealmID   int64     `json:"realm_id"`
	Realm     *Realm    `bun:"rel:belongs-to,join:realm_id=id" json:"realm"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
}

func (u *User) HashPassword() {

	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		fmt.Println("Cannot hash password")
		panic(err)
	}
	u.Password = string(bytes)
}
