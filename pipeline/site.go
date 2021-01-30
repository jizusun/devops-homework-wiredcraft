package pipeline

import (
	"errors"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/blang/semver/v4"
	"github.com/jizusun/wiredcraft-hugo/externals"
)

// Site the hugo site
type Site struct {
	version    string
	envName    string
	workingDir string
}

func loadSite(envName string, dep externals.DependenciesInterface) (*Site, error) {
	workingDir := dep.GetHugoWorkingDir()
	if len(workingDir) == 0 {
		workingDir = dep.GetWorkingDir()
	}
	version, err := getCurrentVersion(dep, workingDir)
	if err != nil {
		return nil, err
	}

	return &Site{version: version, envName: envName, workingDir: workingDir}, nil
}

func getCurrentVersion(dep externals.DependenciesInterface, workingDir string) (string, error) {
	str, err := dep.ReadFileContent(path.Join(workingDir, "config/_default/params.toml"))
	if err != nil {
		return "", err
	}
	var params hugoConfigParamsToml
	toml.Decode(str, &params)
	version := params.Version
	if len(version) == 0 {
		return "", errors.New("No version found")
	}
	return params.Version, nil
}

type hugoConfigParamsToml struct {
	Version string `toml:"version"`
}

func (s *Site) getIncrementedVersion(currentVersion string) string {
	currentSemVer, _ := semver.Make(currentVersion)
	if s.envName == "dev" {
		currentSemVer.IncrementPatch()
	} else {
		currentSemVer.IncrementMinor()
	}
	return currentSemVer.String()

}

func (s *Site) incrementVersion(dep externals.DependenciesInterface) {
	newVersion := s.getIncrementedVersion(s.version)
	s.version = newVersion
}

func (s *Site) compile() {

}

func (s *Site) release() {

}
