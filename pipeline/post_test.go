package pipeline

import (
	"testing"
	"github.com/jizusun/wiredcraft-hugo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func Test_newPost(t *testing.T) {
	timestamp := "2021-01-27T22:58:52+08:00"
	dir := "/home/jizu/hugo-website"
	dep := new(mocks.DependenciesInterface)
	dep.On("GetNow").Return(timestamp)
	dep.On("GetHugoWorkingDir").Return("")
	dep.On("GetWorkingDir").Return(dir)
	actual := newPost(dep)
	expected := &Post{
		fileName:   timestamp + ".md",
		workingDir: dir,
		kind:       "posts",
	}
	assert.Equal(t, expected, actual)
}

func Test_newPost_HugoWorkingDirFromEnv(t *testing.T) {
	timestamp := "2021-01-27T22:58:52+08:00"
	dir := "/home/jizu/hugo-website"
	dep := new(mocks.DependenciesInterface)
	dep.On("GetNow").Return(timestamp)
	dep.On("GetWorkingDir").Return("")
	dep.On("GetHugoWorkingDir").Return(dir)
	actual := newPost(dep)
	expected := &Post{
		fileName:   timestamp + ".md",
		workingDir: dir,
		kind:       "posts",
	}
	assert.Equal(t, expected, actual)
}

type PostTestSuite struct {
	suite.Suite
	post *Post
	dep  *mocks.DependenciesInterface
}

func (suite *PostTestSuite) SetupTest() {
	suite.post = &Post{
		fileName:   "2021-01-27T22:58:52+08:00.md",
		workingDir: "/home/jizu/hugo-website",
		kind:       "posts",
	}
	suite.dep = new(mocks.DependenciesInterface)
}

func (suite *PostTestSuite) Test_save_success() {
	dir := "/home/jizu/hugo-website/content/posts"
	workdir := "/home/jizu/hugo-website"
	suite.dep.On("JoinPath", workdir, "content/posts").Return(dir)
	suite.dep.On("DirExists", dir).Return(true, nil)
	suite.dep.On("ExecHugo", "new posts/2021-01-27T22:58:52+08:00.md", workdir).Return("", nil)
	suite.dep.On("Println", mock.Anything)
	actual := suite.post.save(suite.dep)
	suite.Nil(actual)
}

func (suite *PostTestSuite) Test_appendFortune_success() {
	filepath := "/home/jizu/hugo-website/content/posts/2021-01-27T22:58:52+08:00.md"
	fortune := "carpe diem"
	suite.dep.On("GetFortune").Return(fortune, nil)
	args := []interface{}{"/home/jizu/hugo-website", "/content/", "posts", "2021-01-27T22:58:52+08:00.md"}
	suite.dep.On("JoinPath", args...).Return(filepath)
	suite.dep.On("AppendToFile", filepath, fortune).Return(nil)
	actual := suite.post.appendFortune(suite.dep)
	suite.Nil(actual)
}

func TestPostTestSuite(t *testing.T) {
	suite.Run(t, new(PostTestSuite))
}
