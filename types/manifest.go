package types

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type DockerfileSourceKind string

const (
	kindLocal DockerfileSourceKind = "Local"
	kindGit   DockerfileSourceKind = "Git"
)

type Manifest struct {
	Source DockerfileSource `yaml:"source"`
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
		},
	}

	if err = yaml.Unmarshal(configBytes, manifest); err != nil {
		return nil, err
	}

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
