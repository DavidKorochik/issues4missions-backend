package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashAndCheckSecretCode(t *testing.T) {
	secretCode := RandomString(4) // The length of the secret code should be 4 (not longer or shorter)

	hashedSecretCode, err := HashSecretCode(secretCode)
	require.NoError(t, err)
	require.NotEmpty(t, hashedSecretCode)

	err = CheckSecretCode(secretCode, hashedSecretCode)
	require.NoError(t, err)

	wrongSecretCode := RandomString(4)

	hashedWrongSecretCode, err := HashSecretCode(wrongSecretCode)
	require.NoError(t, err)
	require.NotEmpty(t, hashedSecretCode)

	err = CheckSecretCode(wrongSecretCode, hashedSecretCode)
	require.Error(t, err)
	require.NotEqual(t, hashedSecretCode, hashedWrongSecretCode)
}
