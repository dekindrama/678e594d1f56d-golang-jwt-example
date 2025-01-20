package models

import (
	"time"
)

// User struct
type User struct {
	UserId    string    `json:"user_id" bson:"user_id"`
	Name      string    `json:"name" bson:"name"`
	Username  string    `json:"username" bson:"username"`
	Password  string    `json:"password" bson:"password"`
	CreatedAt time.Time `json:"created_at,omitempty,string" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty,string" bson:"updated_at"`
}
