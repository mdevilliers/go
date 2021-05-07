package exec

import (
	"context"
	"os"
	"os/exec"
	"strings"
)

// GetCommandOutput evaluates the given command and returns the trimmed output
func GetCommandOutput(dir string, name string, args ...string) (string, error) {
	e := exec.Command(name, args...)
	if dir != "" {
		e.Dir = dir
	}
	data, err := e.CombinedOutput()
	text := string(data)
	text = strings.TrimSpace(text)
	return text, err
}

// RunCommand evaluates the given command an error if any
func RunCommand(dir string, name string, args ...string) error {
	e := exec.Command(name, args...)
	if dir != "" {
		e.Dir = dir
	}
	e.Stdout = os.Stdout
	e.Stderr = os.Stdin
	return e.Run()
}

// RunCommandSilent evaluates the given command throwing away any output
func RunCommandSilent(dir string, name string, args ...string) error {
	e := exec.Command(name, args...)
	if dir != "" {
		e.Dir = dir
	}
	return e.Run()
}

type Input struct {
	Name string
	Args []string
}

type Output struct {
	Result []byte
	Err    error
}

func StreamCommandsAndResults(ctx context.Context, dir string, in <-chan Input, out chan<- Output) {

}
