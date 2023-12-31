package git

import (
	"fmt"
	"strings"
)

func FileHasChanges(repoPath, file, lastCommit, currentCommit string) (bool, error) {
	output, err := ExecuteCmdWithOutput(
		"git", "-C", repoPath, "diff", "--name-only", lastCommit, currentCommit)
	if err != nil {
		if strings.Contains(err.Error(), "exit status 128") {
			return false, nil
		}

		return false, fmt.Errorf("failed to check if file has changes cmd=%s %w",
			fmt.Sprintf("git -C %s diff --name-only %s %s", repoPath, lastCommit, currentCommit),
			err,
		)
	}
	return strings.Contains(output, file), nil
}
