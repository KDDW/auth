package models

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID        int64     `bun:",pk,autoincrement" json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	RealmID   int64     `json:"realm_id"`
	Realm     *Realm    `bun:"rel:belongs-to,join:realm_id=id" json:"realm"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
}

func (u *User) HashPassword() {

	cost := 14

	if strings.ToLower(os.Getenv("APPENV")) == "test" {
		cost = 1
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), cost)
	if err != nil {
		fmt.Println("Cannot hash password")
		panic(err)
	}
	u.Password = string(bytes)
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
