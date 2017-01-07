package git

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/kellegous/grr/manifest"
)

// GitDir ...
const GitDir = ".git"

// HasRepo ...
func HasRepo(dir string) bool {
	if _, err := os.Stat(filepath.Join(dir, GitDir)); err == nil {
		return true
	}
	return false
}

func run(cmd string, args ...string) (string, error) {
	var buf bytes.Buffer
	c := exec.Command(cmd, args...)
	c.Stdout = &buf
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return "", err
	}
	return strings.TrimSpace(buf.String()), nil
}

func readURL(dir string) (string, error) {
	return run(
		"git",
		fmt.Sprintf("--git-dir=%s", filepath.Join(dir, GitDir)),
		"remote",
		"get-url",
		"origin")
}

func readRev(dir string) (string, error) {
	return run(
		"git",
		fmt.Sprintf("--git-dir=%s", filepath.Join(dir, GitDir)),
		"rev-parse",
		"HEAD")
}

// Read ...
func Read(dir string) (*manifest.Repo, error) {
	url, err := readURL(dir)
	if err != nil {
		return nil, err
	}

	rev, err := readRev(dir)
	if err != nil {
		return nil, err
	}

	return &manifest.Repo{
		URL: url,
		Rev: rev,
		Vcs: "git",
		Dir: dir,
	}, nil
}
