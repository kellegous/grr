package manifest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ManifestFilename ...
const (
	ManifestFilename = "grr.json"
	DepsDir          = "deps"
)

// ErrManifestNotFound ...
var ErrManifestNotFound = errors.New("not found")

// Repo ...
type Repo struct {
	URL string `json:"url"`
	Dir string `json:"dir"`
	Rev string `json:"rev"`
	Vcs string `json:"git"`
}

// Manifest ...
type Manifest struct {
	Imports []string `json:"imports"`
	Repos   []*Repo  `json:"repos"`
	dir     string
}

// AddImport ...
func (m *Manifest) AddImport(name string) {
	for _, imp := range m.Imports {
		if imp == name {
			return
		}
	}

	m.Imports = append(m.Imports, name)
}

// CreateIn ...
func CreateIn(dir string) error {
	return ioutil.WriteFile(
		filepath.Join(dir, ManifestFilename),
		[]byte("{}\n"),
		os.ModePerm)
}

// ExistsIn ...
func ExistsIn(dir string) bool {
	if _, err := os.Stat(filepath.Join(dir, ManifestFilename)); err == nil {
		return true
	}
	return false
}

// LoadFrom ...
func (m *Manifest) LoadFrom(dir string) error {
	filename, err := FindManifestIn(dir)
	if err != nil {
		return err
	}

	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer r.Close()

	if err := json.NewDecoder(r).Decode(m); err != nil {
		return err
	}

	m.dir = dir

	return nil
}

// Save ...
func (m *Manifest) Save() error {
	b, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(
		filepath.Join(m.dir, ManifestFilename),
		b,
		os.ModePerm)
}

// DepsPath ...
func (m *Manifest) DepsPath() string {
	return filepath.Join(m.dir, DepsDir)
}

// FindManifestIn ...
func FindManifestIn(dir string) (string, error) {
	for {
		path := filepath.Join(dir, ManifestFilename)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}

		par := filepath.Dir(dir)
		if par == dir {
			return "", ErrManifestNotFound
		}

		dir = par
	}
}
