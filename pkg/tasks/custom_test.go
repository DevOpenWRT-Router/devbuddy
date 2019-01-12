package tasks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCustom(t *testing.T) {
	task := ensureLoadTestTask(t, `
custom:
  met?: test-command
  meet: custom-command
`)

	require.Equal(t, "Task Custom (custom-command) has 1 actions", task.Describe())
}

func TestCustomName(t *testing.T) {
	task := ensureLoadTestTask(t, `
custom:
  name: NAMENAME
  met?: test-command
  meet: custom-command
`)

	require.Equal(t, "Task Custom (NAMENAME) has 1 actions", task.Describe())
}

func TestCustomWithBoolean(t *testing.T) {
	_, err := loadTestTask(t, `
custom:
  met?: false
  meet: custom-command
`)

	require.Error(t, err, "buildFromDefinition() should have failed")
	require.Contains(t, err.Error(), "not a string")
}
