package integration

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	devYmlMyCmd = `
commands:
  mycmd:
    run: touch somefile
`
	devYmlMyCmdShort = `
commands:
  mycmd: touch somefile
`
)

func Test_Cmd_Custom(t *testing.T) {
	c := CreateContextAndInit(t)

	project := CreateProject(c, "project", devYmlMyCmd)
	c.Cd(project.Path)

	lines := c.Run("bud mycmd")
	OutputEqual(t, lines, "🐼  running touch somefile")

	files := c.Ls(".")
	require.ElementsMatch(t, files, []string{"dev.yml", "somefile"})
}

func Test_Cmd_Custom_Short_Syntax(t *testing.T) {
	c := CreateContextAndInit(t)

	project := CreateProject(c, "project", devYmlMyCmdShort)
	c.Cd(project.Path)

	lines := c.Run("bud mycmd")
	OutputEqual(t, lines, "🐼  running touch somefile")
}

func Test_Cmd_Custom_Envs_Are_Applied(t *testing.T) {
	c := CreateContextAndInit(t)

	project := CreateProject(c, "project",
		`env:`,
		`  MYVAR: poipoi`,
		`commands:`,
		`  mycmd: "echo __${MYVAR}__ > result"`,
	)
	c.Cd(project.Path)

	c.Run("bud mycmd")
	c.AssertContains("result", "__poipoi__")
}

func Test_Cmd_Custom_Always_Run_In_Project_Root(t *testing.T) {
	c := CreateContextAndInit(t)

	project := CreateProject(c, "project", devYmlMyCmd)
	c.Cd(project.Path)
	c.Run("mkdir foobar")
	c.Cd("foobar")

	lines := c.Run("bud mycmd")
	OutputEqual(t, lines, "🐼  running touch somefile")

	files := c.Ls("..")
	require.ElementsMatch(t, files, []string{"dev.yml", "foobar", "somefile"})
}
