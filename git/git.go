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

func clone(dir, url string) error {
	return exec.Command("git", "clone", url, dir).Run()
}

func fetch(dir string) error {
	return exec.Command("git", fmt.Sprintf("--git-dir=%s", filepath.Join(dir, ".git")), "fetch").Run()
}

func resetTo(dir, rev string) error {
	return exec.Command("git",
		fmt.Sprintf("--git-dir=%s", filepath.Join(dir, ".git")),
		"reset",
		rev,
		"--hard").Run()
}

// Restore ...
func Restore(r *manifest.Repo) error {
	if _, err := os.Stat(r.Dir); err != nil {
		if err := clone(r.Dir, r.URL); err != nil {
			return err
		}
	}

	rc, err := Read(r.Dir)
	if err != nil {
		return err
	}

	if rc.Rev == r.Rev {
		return nil
	}

	if err := fetch(r.Dir); err != nil {
		return err
	}

	return resetTo(r.Dir, r.Rev)
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
