package types

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/scottweitzner/crane/common"
	"github.com/scottweitzner/crane/internal"
	"gopkg.in/yaml.v3"
)

type DockerfileSourceKind string

const (
	kindLocal DockerfileSourceKind = "Local"
	kindGit   DockerfileSourceKind = "Git"
)

type Manifest struct {
	Source DockerfileSource       `yaml:"source"`
	Values map[string]interface{} `yaml:"values"`
	Output Output                 `yaml:"output"`
}

type Output struct {
	Path      string `yaml:"path"`
	Extension string `yaml:"extension"`
}

type DockerfileSource struct {
	Kind        DockerfileSourceKind `yaml:"kind"`
	GitSource   GitSource            `yaml:"git"`
	LocalSource LocalSource          `yaml:"local"`
}

type LocalSource struct {
	Path string `yaml:"path"`
}

type GitSource struct {
	URL     string `yaml:"url"`
	Version string `yaml:"version"`
	Path    string `yaml:"path"`
}

func ParseManifest(path string) (*Manifest, error) {
	configBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	manifest := &Manifest{
		Source: DockerfileSource{
			Kind: kindLocal,
			LocalSource: LocalSource{
				Path: "Dockerfile",
			},
			GitSource: GitSource{
				Path: "Dockerfile",
			},
		},
		Output: Output{
			Path:      "./crane",
			Extension: ".Dockerfile",
		},
	}

	if err = yaml.Unmarshal(configBytes, manifest); err != nil {
		return nil, err
	}

	// validations for source
	switch manifest.Source.Kind {
	case kindLocal:
		if manifest.Source.LocalSource == (LocalSource{}) {
			return nil, fmt.Errorf("could not find %s source definition", kindLocal)
		}
	case kindGit:
		if manifest.Source.GitSource == (GitSource{}) {
			return nil, fmt.Errorf("could not find %s source definition", kindGit)
		}
	default:
		return nil, fmt.Errorf("unknown source kind %s", manifest.Source.Kind)
	}

	return manifest, nil
}

func (manifest *Manifest) FormSourcePath() (string, error) {
	switch manifest.Source.Kind {
	case kindLocal:
		return manifest.Source.LocalSource.Path, nil
	case kindGit:
		if err := internal.CloneAndSwitchVersion(manifest.Source.GitSource.URL, manifest.Source.GitSource.Version); err != nil {
			return "", err
		}
		path := strings.TrimPrefix(manifest.Source.GitSource.Path, "/")
		return fmt.Sprintf("%s/%s", common.GitClonePath, path), nil
	}
}

func (manifest *Manifest) FormOutputPath() string {
	return fmt.Sprintf("%s%s", manifest.Output.Path, manifest.Output.Extension)
}
