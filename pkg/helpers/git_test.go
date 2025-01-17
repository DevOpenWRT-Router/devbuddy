package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/devbuddy/devbuddy/pkg/test"
)

func TestGitGithubProjectURL(t *testing.T) {
	tmpdir := t.TempDir()
	writer := test.Project(tmpdir)
	writer.CreateGitRepo(t)

	url, err := NewGitRepo(tmpdir).BuildGithubProjectURL()

	require.NoError(t, err, "BuildGithubProjectURL() failed")
	require.Equal(t, "https://github.com/org1/repo1/tree/master", url)
}

func TestGitGithubPullrequestURL(t *testing.T) {
	tmpdir := t.TempDir()
	writer := test.Project(tmpdir)
	writer.CreateGitRepo(t)

	url, err := NewGitRepo(tmpdir).BuildGithubPullrequestURL()

	require.NoError(t, err, "BuildGithubProjectURL() failed")
	require.Equal(t, "https://github.com/org1/repo1/pull/master?expand=1", url)
}
