package token

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID          uuid.UUID `json:"id"`
	UserID      int64     `json:"user_id"`
	AccountType int64     `json:"account_type"`
	IssuedAt    time.Time `json:"issued_at"`
	ExpiredAt   time.Time `json:"expired_at"`
}

func NewPayload(userID int64, accountType int64, daration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:          tokenID,
		UserID:      userID,
		AccountType: accountType,
		IssuedAt:    time.Now(),
		ExpiredAt:   time.Now().Add(daration),
	}

	return payload, nil
}
