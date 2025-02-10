package userModel

import (
	"fmt"
	"time"
)

type UserPayload struct {
	KodeUser    string `json:"kode_user"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type CreateUserResponse struct {
	UserPayload
}

type ExpiredPayload struct {
	Exp int64 `json:"exp"`
}

type SessionPayload struct {
	SessionId string `json:"sess_id"`
}

// Payload contains the payload data of the token
type Payload struct {
	ID      string `json:"jti"`
	Subject string `json:"sub"`
	SessionPayload
	ExpiredPayload
}

// Payload contains the payload data of the token verify
type VerifyPayload struct {
	Payload
	//ExpiredIn int `json:"expired_in"`
	SessionPayload
}

// NewPayload creates a new token payload with a specific user and duration
func NewPayload(userExpiredPayload ExpiredPayload, userSessionPayload SessionPayload, userPayload UserPayload, duration time.Duration) (*Payload, error) {
	payload := &Payload{
		ID:             userPayload.KodeUser,
		Subject:        userPayload.Email,
		SessionPayload: userSessionPayload,
		ExpiredPayload: userExpiredPayload,
	}

	if err := payload.Valid(); err != nil {
		return nil, err
	}

	return payload, nil
}

// Valid checks if the payload is valid or not
func (p *Payload) Valid() error {
	if p.SessionId == "" {
		return fmt.Errorf("invalid payload")
	}

	return nil
}
