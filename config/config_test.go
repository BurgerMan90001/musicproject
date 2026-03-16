package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApi(t *testing.T) {
	cfg := ReadConfigFile()
	t.Run("port number", func(t *testing.T) {

		assert.Equal(t, 8081, cfg.API.Port, "port number")
	})
	// t.Run("", func(t *testing.T) {
	// 	assert.Equal(t, "", "")
	// })
}
