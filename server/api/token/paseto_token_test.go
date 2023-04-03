package token

import (
	"github/ThoPham02/research_management/api/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreatePaseto(t *testing.T) {
	maker, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)

	userID := utils.RandomInt(0, 10)
	accountType := utils.RandomInt(0, 4)
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := time.Now().Add(duration)

	token, err := maker.CreateToken(userID, accountType, duration)
	require.NoError(t, err)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)

	require.Equal(t, userID, payload.UserID)
	require.Equal(t, accountType, payload.AccountType)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)

}
