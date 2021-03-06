package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"go.uber.org/zap"
)

type CmdExecutor struct {
	log *zap.SugaredLogger
}

func NewExecutor(log *zap.SugaredLogger) *CmdExecutor {
	return &CmdExecutor{
		log: log,
	}
}

func (c *CmdExecutor) ExecuteCommandWithOutput(command string, env []string, arg ...string) (string, error) {
	commandWithPath, err := exec.LookPath(command)
	if err != nil {
		return fmt.Sprintf("unable to find command:%s in path", command), err
	}
	c.log.Infow("running command", "command", commandWithPath, "args", strings.Join(arg, " "))
	cmd := exec.Command(commandWithPath, arg...)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, env...)
	return runCommandWithOutput(cmd, true)
}

func runCommandWithOutput(cmd *exec.Cmd, combinedOutput bool) (string, error) {
	var output []byte
	var err error

	if combinedOutput {
		output, err = cmd.CombinedOutput()
	} else {
		output, err = cmd.Output()
	}

	out := strings.TrimSpace(string(output))

	return out, err
}
