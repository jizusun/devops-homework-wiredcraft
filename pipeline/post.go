package pipeline

import (
	"errors"

	"github.com/jizusun/wiredcraft-hugo/externals"
)

// Post the hugo post
type Post struct {
	fileName   string
	workingDir string
	kind       string
}

func newPost(dep externals.DependenciesInterface) *Post {
	fileName := dep.GetNow() + ".md"
	targetPath := dep.GetHugoWorkingDir()
	if len(targetPath) == 0 {
		targetPath = dep.GetWorkingDir()
	}
	return &Post{
		fileName:   fileName,
		workingDir: targetPath,
		kind:       "posts",
	}
}

func (p *Post) contentFolderExist(dep externals.DependenciesInterface) bool {
	postPath := dep.JoinPath(p.workingDir, "content")
	exist, _ := dep.DirExists(postPath)
	return exist
}

func (p *Post) save(dep externals.DependenciesInterface) error {
	isCorrectWorkingDir := p.contentFolderExist(dep)
	if !isCorrectWorkingDir {
		return errors.New("The folder is required: content")
	}
	output, err := dep.ExecHugo("new "+p.kind+"/"+p.fileName, p.workingDir)
	dep.Println(output)
	return err
}

func (p *Post) getFilePath(dep externals.DependenciesInterface) string {
	return dep.JoinPath(p.workingDir, "/content/", p.kind, p.fileName)
}

func (p *Post) appendFortune(dep externals.DependenciesInterface) error {
	str, err := dep.GetFortune()
	if err != nil {
		return err
	}
	filePath := p.getFilePath(dep)
	err = dep.AppendToFile(filePath, str)
	if err != nil {
		return err
	}
	return nil
}
