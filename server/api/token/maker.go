package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken creates a new token
	CreateToken(userID int32, accountType int32, duration time.Duration) (string, error)
	// VerifyToken verifies the token
	VerifyToken(token string) (*Payload, error)
}
