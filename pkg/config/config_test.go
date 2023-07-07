package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	t.Run("Load config with error", func(t *testing.T) {
		got, err := LoadConfig("")
		require.Error(t, err)
		require.Nil(t, got)
	})

	t.Run("Load config without error", func(t *testing.T) {
		got, err := LoadConfig("../../configs/config.json")
		require.NoError(t, err)
		require.NotNil(t, got)
	})
}
