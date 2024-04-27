package entity

import (
	"github.com/google/uuid"
	"time"
)

type (
	UserDTO struct {
		Id        uuid.UUID
		Name      string
		Email     string
		Password  string
		Role      string
		CreatedAt *time.Time
	}
	User struct {
		id        uuid.UUID
		name      string
		email     string
		password  string
		role      string
		createdAt *time.Time
	}
)

func NewUser(payload *UserDTO) (*User, error) {
	return RebuildUser(payload), nil
}
func RebuildUser(payload *UserDTO) *User {
	return &User{
		id:        payload.Id,
		name:      payload.Name,
		email:     payload.Email,
		password:  payload.Password,
		role:      payload.Role,
		createdAt: nil,
	}
}

func (p *User) Id() uuid.UUID {
	return p.id
}
func (p *User) Name() string {
	return p.name
}
func (p *User) Email() string {
	return p.email
}
func (p *User) Password() string {
	return p.password
}
func (p *User) SetPassword(pass string) {
	p.password = pass
}

func (p *User) Role() string { return p.role }
func (p *User) CreatedAt() *time.Time {
	return p.createdAt
}
