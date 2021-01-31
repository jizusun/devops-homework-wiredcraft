package pipeline

import (
	"errors"
	"testing"

	"github.com/jizusun/wiredcraft-hugo/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type SiteTestSuite struct {
	suite.Suite
	site *Site
	dep  *mocks.DependenciesInterface
}

func (suite *SiteTestSuite) SetupTest() {
	suite.dep = new(mocks.DependenciesInterface)
	suite.site = &Site{
		envName:    "dev",
		workingDir: "/home/jizu/hugo-website",
		version:    "0.1.5",
	}
}

func (suite *SiteTestSuite) Test_newSite() {
	tomlData := []byte(`version = "0.0.2"`)
	suite.dep.On("GetHugoWorkingDir").Return("")
	suite.dep.On("GetWorkingDir").Return("/home/jizu/hugo-website")
	suite.dep.On("ReadFileContent", "/home/jizu/hugo-website/config/_default/params.toml").Return(tomlData, nil)
	actual, _ := loadSite("dev", suite.dep)
	expected := &Site{
		envName:    "dev",
		workingDir: "/home/jizu/hugo-website",
		version:    "0.0.2",
	}
	suite.Equal(expected, actual)
}

func (suite *SiteTestSuite) Test_newSite_failedToReadFile() {
	errFailedToOpen := errors.New("failed to open the file")
	suite.dep.On("GetHugoWorkingDir").Return("")
	suite.dep.On("GetWorkingDir").Return("/home/jizu/hugo-website")
	suite.dep.On("ReadFileContent", "/home/jizu/hugo-website/config/_default/params.toml").Return(nil, errFailedToOpen)
	site, err := loadSite("dev", suite.dep)
	suite.Equal(err, errFailedToOpen)
	suite.Nil(site)
}
func (suite *SiteTestSuite) Test_newSite_EmptyVersion() {
	tomlData := []byte(``)
	suite.dep.On("GetHugoWorkingDir").Return("")
	suite.dep.On("GetWorkingDir").Return("/home/jizu/hugo-website")
	suite.dep.On("ReadFileContent", "/home/jizu/hugo-website/config/_default/params.toml").Return(tomlData, nil)
	site, err := loadSite("dev", suite.dep)
	suite.NotNil(err)
	suite.Nil(site)
}

func (suite *SiteTestSuite) Test_incrementVersion_Dev() {
	suite.dep.On("WriteFile", "/home/jizu/hugo-website/config/_default/params.toml", mock.Anything).Return(nil)
	suite.site.incrementVersion(suite.dep)
	expected := "0.1.6"
	suite.Equal(suite.site.version, expected)
}

func (suite *SiteTestSuite) Test_incrementVersion_Staging() {
	suite.site.envName = "staging"
	suite.dep.On("WriteFile", mock.Anything, mock.Anything).Return(nil)
	suite.site.incrementVersion(suite.dep)
	expected := "0.2.0"
	suite.Equal(suite.site.version, expected)
}

// func (suite *SiteTestSuite) Test_compile() {
// 	suite.site.compile(suite.dep)
// }

func (suite *SiteTestSuite) Test_release_dev() {
	suite.dep.On("AddCommitAndPush", mock.Anything, mock.Anything).Return("4aff69", nil)
	actual := suite.site.release(suite.dep)
	suite.Nil(actual)
}

func (suite *SiteTestSuite) Test_release_staging() {
	suite.site.envName = "staging"
	suite.dep.On("AddCommitAndPush", mock.Anything, mock.Anything).Return("4aff69", nil)
	suite.dep.On("GitTagAndPush", mock.Anything, mock.Anything).Return(nil)
	actual := suite.site.release(suite.dep)
	suite.Nil(actual)
}
func TestSiteSuite(t *testing.T) {
	suite.Run(t, new(SiteTestSuite))
}
