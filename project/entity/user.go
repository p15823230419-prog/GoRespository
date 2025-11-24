package entity

import "time"

type User struct {
	Id        uint64
	Username  string
	Nickname  string
	Avatar    string
	Password  string
	Mobile    string
	Email     string
	Status    int8
	CreatedAt time.Time
}
