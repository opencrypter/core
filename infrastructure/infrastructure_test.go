package infrastructure_test

import (
	"github.com/opencrypter/api/infrastructure"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMigrateDb(t *testing.T) {
	t.Run("It should use default values on missing environment variables", func(t *testing.T) {
		os.Clearenv()
		assert.NotPanics(t, infrastructure.MigrateDb)
	})

	t.Run("It should crash on connection fail", func(t *testing.T) {
		os.Setenv("DB_HOST", "invalid")
		assert.Panics(t, infrastructure.MigrateDb)
	})
}
