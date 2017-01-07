package create

import (
	"log"

	"grr/manifest"
)

// Run ...
func Run(dir string, args []string) {
	if manifest.ExistsIn(dir) {
		return
	}

	if err := manifest.CreateIn(dir); err != nil {
		log.Panic(err)
	}
}
