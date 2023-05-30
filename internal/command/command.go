package command

import (
	"bytes"
	"github.com/alleswebdev/go-command-executor/internal/config"
	"github.com/pkg/errors"
	"io"
	"os"
	"os/exec"
)

type Command struct {
	name    string
	start   string
	stop    string
	restart string
}

func GetCommandsMapFromConfig(cfg config.Config) map[string]Command {
	result := make(map[string]Command)

	for _, command := range cfg.Commands {
		result[command.Name] = Command{
			name:    command.Name,
			start:   command.Start,
			stop:    command.Stop,
			restart: command.Restart,
		}
	}

	return result
}

func (c Command) GetName() string {
	return c.name
}

func (c Command) Start() (string, error) {
	return execAndWait(c.start)
}

func (c Command) Stop() (string, error) {
	return execAndWait(c.stop)
}

func (c Command) Restart() (string, error) {
	return execAndWait(c.restart)
}

func execAndWait(command string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", command)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Start()
	if err != nil {
		return "", errors.Wrap(err, "cmd.Start")
	}

	err = cmd.Wait()
	if err != nil {
		return "", errors.Wrap(err, "cmd.Wait")
	}

	return stdoutBuf.String(), nil
}
