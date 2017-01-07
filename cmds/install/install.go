package install

import (
	"log"

	"github.com/kellegous/grr/internal"
	"github.com/kellegous/grr/manifest"
)

// Run ...
func Run(dir string, args []string) {
	var m manifest.Manifest

	if err := m.LoadFrom(dir); err != nil {
		log.Panic(err)
	}

	// TODO(knorton): This should ensure that all repos are present.
	for _, arg := range m.Imports {
		if err := internal.Go(&m, "install", arg); err != nil {
			log.Panic(err)
		}
	}
}
