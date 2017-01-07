package main

import (
	"flag"
	"log"
	"os"

	"github.com/kellegous/grr/cmds/create"
	"github.com/kellegous/grr/cmds/get"
	"github.com/kellegous/grr/cmds/help"
	"github.com/kellegous/grr/cmds/install"
)

var cmds = map[string]func(string, []string){
	"init":    create.Run,
	"get":     get.Run,
	"install": install.Run,
	"help":    help.Run,
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
