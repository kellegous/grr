package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/kellegous/grr/manifest"
)

func envFor(m *manifest.Manifest) []string {
	env := os.Environ()
	val := fmt.Sprintf("GOPATH=%s", strings.Join(m.GoPath(), ":"))
	for i, v := range env {
		if strings.HasPrefix(v, "GOPATH=") {
			env[i] = val
			return env
		}
	}

	return append(env, val)
}

// Go ...
func Go(m *manifest.Manifest, args ...string) error {
	c := exec.Command("go", args...)
	c.Env = envFor(m)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	return c.Run()
}
