package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// DepsDir ...
const DepsDir = "deps"

func envFor(dir string) []string {
	env := os.Environ()
	val := fmt.Sprintf("GOPATH=%s:%s", filepath.Join(dir, DepsDir), dir)
	for i, v := range env {
		if strings.HasPrefix(v, "GOPATH=") {
			env[i] = val
			return env
		}
	}

	return append(env, val)
}

// Go ...
func Go(dir string, args ...string) error {
	c := exec.Command("go", args...)
	c.Env = envFor(dir)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	return c.Run()
}
