package custom_helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildAccountID(t *testing.T) {
	t.Run("Build from valid connector ID", func(t *testing.T) {
		connectorIDStr := "eyJQcm92aWRlciI6InN0cmlwZSIsIlJlZmVyZW5jZSI6IjM4NGEyYzNkLWUzZDktNDY1NS04NjkzLWY1MjQxMTllMjIyNyJ9"
		expectedAccountId := "eyJDb25uZWN0b3JJRCI6eyJQcm92aWRlciI6InN0cmlwZSIsIlJlZmVyZW5jZSI6IjM4NGEyYzNkLWUzZDktNDY1NS04NjkzLWY1MjQxMTllMjIyNyJ9LCJSZWZlcmVuY2UiOiJhY2N0XzFKTzM4WEVwczQ3bDJ4Um0ifQ"

		accountIDStr, err := BuildAccountID(connectorIDStr, "acct_1JO38XEps47l2xRm")
		require.NoError(t, err)
		assert.Equal(t, expectedAccountId, accountIDStr)
	})

	t.Run("Fail to build from invalid connectorId", func(t *testing.T) {
		connectorIDStr := "invalid"

		_, err := BuildAccountID(connectorIDStr, "someAccountReference")
		require.Errorf(t, err, "foo")
	})
}
