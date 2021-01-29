package pipeline

import (
	"testing"

	"github.com/jizusun/wiredcraft-hugo/mocks"
	"github.com/stretchr/testify/suite"
)

type SiteTestSuite struct {
	suite.Suite
	site *Site
	dep  *mocks.DependenciesInterface
}

func (suite *SiteTestSuite) SetupTest() {
	suite.dep = new(mocks.DependenciesInterface)
	suite.dep.On("GetHugoWorkingDir").Return("")
	suite.dep.On("GetWorkingDir").Return("/home/jizu/hugo-website")
	suite.site = newSite("dev", suite.dep)
}
func (suite *SiteTestSuite) Test_newSite() {
	expected := &Site{
		envName:    "dev",
		workingDir: "/home/jizu/hugo-website",
	}
	suite.Equal(expected, suite.site)
}

func (suite *SiteTestSuite) Test_getCurrentVersion() {
	tomlData := `
[params]
	version = "0.0.2"`
	suite.dep.On("GetHugoConfigToml", "/home/jizu/hugo-website").Return(tomlData, nil)
	actual, _ := suite.site.getCurrentVersion(suite.dep)
	expected := "0.0.2"
	suite.Equal(actual, expected)
}

func (suite *SiteTestSuite) Test_getCurrentVersion_EmptyToml() {
	tomlData := ""
	suite.dep.On("GetHugoConfigToml", "/home/jizu/hugo-website").Return(tomlData, nil)
	actual, _ := suite.site.getCurrentVersion(suite.dep)
	expected := ""
	suite.Equal(actual, expected)
}

func (suite *SiteTestSuite) Test_incrementVersion() {
	actual := suite.site.incrementVersion(suite.dep)
	expected := "0.0.1"
	suite.Equal(actual, expected)
}

func TestSiteSuite(t *testing.T) {
	suite.Run(t, new(SiteTestSuite))
}
