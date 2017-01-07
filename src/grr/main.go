package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"path/filepath"

	"grr/cmds/create"
	"grr/cmds/get"
)

const (
	manifestFile = "grr.json"
)

var cmds = map[string]func(string, []string){
	"init": create.Run,
	"get":  get.Run,
}

// FindManifest ...
func FindManifest() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		path := filepath.Join(dir, manifestFile)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}

		par := filepath.Dir(dir)
		if par == dir {
			return "", errors.New("manifest not found")
		}

		dir = par
	}
}

func cmdInit() {
	// if manifest exists, done.
	// write new manifest
}

func cmdGet() {
	// build GOPATH, execute go get
	// walk the deps/src directory and build manifest
	// write manifest
}

func cmdRestore() {
	// read manifest
	// for each repo, make sure the repo is at the right rev and cloned
}

func cmdInstall() {
	// read manifest
	// for each import, build GOPATH, execute go install
}

func getWorkingDir(dir string) (string, error) {
	if dir != "" {
		return dir, nil
	}

	return os.Getwd()
}

func main() {
	flagDir := flag.String("dir", "", "")
	flag.Parse()

	wd, err := getWorkingDir(*flagDir)
	if err != nil {
		log.Panic(err)
	}

	args := flag.Args()
	cmd := "help"
	if len(args) > 0 {
		cmd = args[0]
		args = args[1:]
	}

	h := cmds[cmd]
	if h == nil {
		log.Panic("no command")
	}

	h(wd, args)

	// init
	// get
	// restore
	// install
}
