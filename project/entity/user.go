package entity

import "time"

type User struct {
	Id        uint64
	Username  string
	Nickname  string
	Password  string
	Phone     string
	Email     string
	Status    int8
	Avatar    string
	Role      []Role
	CreatedAt time.Time
	UpdatedAt time.Time
}
