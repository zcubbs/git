package git

import (
	"os/exec"
	"strings"
)

// ExecuteCmdWithOutput executes a command and returns the output
func ExecuteCmdWithOutput(command string, args ...string) (string, error) {
	var stdout, stderr strings.Builder

	cmd := exec.Command(command, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return stderr.String(), err
	}
	return stdout.String(), nil
}
