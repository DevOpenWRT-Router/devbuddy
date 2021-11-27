package integration

import (
	"testing"

	"github.com/devbuddy/devbuddy/tests/context"
	"github.com/stretchr/testify/require"
)

func Test_Env_Single(t *testing.T) {
	c := CreateContextAndInit(t)

	c.Write("dev.yml", `env: {KEY1: VAL1}`)

	lines := c.Run("bud up", context.ExitCode(0))
	OutputEqual(t, lines, "◼︎ Env")
	require.Equal(t, "VAL1", c.GetEnv("KEY1"))

	// Add a second var

	c.Write("dev.yml", `env: {KEY1: VAL1, KEY2: VAL2}`)

	lines = c.Run("bud up", context.ExitCode(0))
	OutputEqual(t, lines, "◼︎ Env")
	require.Equal(t, "VAL1", c.GetEnv("KEY1"))
	require.Equal(t, "VAL2", c.GetEnv("KEY2"))

	// Clean the env when leaving the project directory

	c.Run("cd /")
	require.Equal(t, "", c.GetEnv("KEY1"))
	require.Equal(t, "", c.GetEnv("KEY2"))
}
