package token

import (
	"github/ThoPham02/research_management/api/constant"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Payload struct {
	ID          uuid.UUID `json:"id"`
	UserID      int32     `json:"userId"`
	TypeAccount int32     `json:"typeAccount"`
	IssuedAt    time.Time `json:"issuedAt"`
	ExpiredAt   time.Time `json:"expiredAt"`
}

func NewPayload(userID int32, typeAccount int32, daration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:          tokenID,
		UserID:      userID,
		TypeAccount: typeAccount,
		IssuedAt:    time.Now(),
		ExpiredAt:   time.Now().Add(daration),
	}

	return payload, nil
}

func GetPayload(c *gin.Context) *Payload {
	payload := c.Value(constant.PayloadKey).(*Payload)

	return payload
}
