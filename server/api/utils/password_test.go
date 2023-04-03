package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	password := RandomString(8)

	hashPassword, err := HashPassword(password)
	require.NoError(t, err)

	correct := ComparePassword(password, hashPassword)
	require.True(t, correct)

	wrongPassword := RandomString(8)
	wrong := ComparePassword(wrongPassword, hashPassword)
	require.False(t, wrong)
}
