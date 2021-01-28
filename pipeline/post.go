package pipeline

// Post the hugo post
type Post struct {
	fileName   string
	workingDir string
	kind       string
}

func newPost() *Post {
	fileName := getNow() + ".md"
	targetPath := getHugoWorkingDir()
	if len(targetPath) == 0 {
		targetPath = getWorkingDir()
	}
	return &Post{
		fileName:   fileName,
		workingDir: targetPath,
		kind:       "posts",
	}
}

func (p *Post) save() error {
	postPath := joinPath(p.workingDir, "content/"+p.kind)
	exist, err := dirExists(postPath)
	if !exist {
		println("New post cannot be created.")
		return err
	}
	var output string
	output, err = execHugo("new "+p.kind+"/"+p.fileName, p.workingDir)
	println(output)
	return err
}

func (p *Post) getFilePath() string {
	filePath := joinPath(p.workingDir, "/content/", p.kind, p.fileName)
	println(filePath)
	return filePath
}

func (p *Post) appendFortune() error {
	str, err := getFortune()
	if err != nil {
		return err
	}
	appendToFile(p.getFilePath(), str)
	return nil
}
