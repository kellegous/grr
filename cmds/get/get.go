package get

import (
	"log"
	"os"
	"path/filepath"

	"github.com/kellegous/grr/git"
	"github.com/kellegous/grr/internal"
	"github.com/kellegous/grr/manifest"
)

func findReposFor(m *manifest.Manifest) ([]*manifest.Repo, error) {
	var repos []*manifest.Repo
	if err := filepath.Walk(
		m.DepsPath(),
		func(path string, nfo os.FileInfo, err error) error {
			if !nfo.IsDir() || !git.HasRepo(path) {
				return nil
			}

			repo, err := git.Read(path)
			if err != nil {
				return err
			}

			repos = append(repos, repo)

			return filepath.SkipDir
		}); err != nil {
		return nil, err
	}

	return repos, nil
}

// Run ...
func Run(dir string, args []string) {
	var m manifest.Manifest

	if err := m.LoadFrom(dir); err != nil {
		log.Panic(err)
	}

	for _, arg := range args {
		if err := internal.Go(&m, "get", "-u", arg); err != nil {
			log.Panic(err)
		}

		m.AddImport(arg)
	}

	repos, err := findReposFor(&m)
	if err != nil {
		log.Panic(err)
	}

	m.Repos = repos
	if err := m.Save(); err != nil {
		log.Panic(err)
	}
}
