package pipeline

import "github.com/jizusun/wiredcraft-hugo/externals"

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

func (p *Post) save(dep externals.DependenciesInterface) error {
	postPath := dep.JoinPath(p.workingDir, "content/"+p.kind)
	exist, err := dep.DirExists(postPath)
	if !exist {
		dep.Println("New post cannot be created.")
		return err
	}
	var output string
	output, err = dep.ExecHugo("new "+p.kind+"/"+p.fileName, p.workingDir)
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
