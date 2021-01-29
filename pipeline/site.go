package pipeline

import (
	"github.com/BurntSushi/toml"
	"github.com/jizusun/wiredcraft-hugo/externals"
)

// Site the hugo site
type Site struct {
	version    string
	envName    string
	workingDir string
}

func newSite(envName string, dep externals.DependenciesInterface) *Site {
	targetPath := dep.GetHugoWorkingDir()
	if len(targetPath) == 0 {
		targetPath = dep.GetWorkingDir()
	}
	return &Site{envName: envName, workingDir: targetPath}
}

func (s *Site) getCurrentVersion(dep externals.DependenciesInterface) (string, error) {
	str, err := dep.GetHugoConfigToml(s.workingDir)
	if err != nil {
		return "", err
	}
	var conf HugoTomlConfig
	toml.Decode(str, &conf)
	return conf.Params.Version, nil
}

type HugoTomlConfig struct {
	Params struct {
		Version string `toml:"version"`
	} `toml:"params"`
}

func (s *Site) incrementVersion(dep externals.DependenciesInterface) string {
	return "0.0.1"
}

func (s *Site) compile() {

}

func (s *Site) release() {

}
