package pipeline

import (
	"errors"
	"fmt"
	"path"

	"github.com/blang/semver/v4"
	"github.com/jizusun/wiredcraft-hugo/externals"
	"github.com/pelletier/go-toml"
)

// Site the hugo site
type Site struct {
	envName    string
	workingDir string
	version    string
	oldVersion string
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
	paramsTomlPath := path.Join(workingDir, "config/_default/params.toml")
	bytes, err := dep.ReadFileContent(paramsTomlPath)
	if err != nil {
		return "", err
	}
	var params hugoConfigParamsToml
	toml.Unmarshal(bytes, &params)
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

func (s *Site) incrementVersion(dep externals.DependenciesInterface) error {
	newVersion := s.getIncrementedVersion(s.version)
	params := &hugoConfigParamsToml{
		Version: newVersion,
	}
	b, _ := toml.Marshal(params)
	paramsTomlPath := path.Join(s.workingDir, "config/_default/params.toml")
	err := dep.WriteFile(paramsTomlPath, b)
	if err != nil {
		return err
	}
	s.oldVersion = s.version
	s.version = newVersion
	return nil
}

func (s *Site) compile(dep externals.DependenciesInterface) error {
	output, err := dep.ExecHugo("", s.workingDir)
	fmt.Println(output)
	return err
}

func (s *Site) release(dep externals.DependenciesInterface) error {
	hugoBranchCommitMessage := "Version: " + s.oldVersion + " => " + s.version
	// the hugo branch
	err := dep.AddCommitAndPush(hugoBranchCommitMessage, s.workingDir)
	// the gh-pages branch, the public folde
	publicDir := path.Join(s.workingDir + "/public")
	err = dep.AddCommitAndPush("build site", publicDir)
	return err
}
