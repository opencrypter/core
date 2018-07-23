package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMigrateDb(t *testing.T) {
	t.Run("It should crash on connection fail", func(t *testing.T) {
		os.Setenv("DB_HOST", "invalid")

		assert.Panics(t, MigrateDb)
	})
}
