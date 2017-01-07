package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"path/filepath"

	"grr/cmds/create"
	"grr/cmds/get"
	"grr/cmds/help"
	"grr/cmds/install"
)

const (
	manifestFile = "grr.json"
)

var cmds = map[string]func(string, []string){
	"init":    create.Run,
	"get":     get.Run,
	"install": install.Run,
	"help":    help.Run,
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
}
