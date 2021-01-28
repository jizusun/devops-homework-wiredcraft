package pipeline

// Post the hugo post
type Post struct {
	fileName   string
	workingDir string
	kind       string
}

func newPost() *Post {
	fileName := GetNow() + ".md"
	targetPath := GetHugoWorkingDir()
	if len(targetPath) == 0 {
		targetPath = GetWorkingDir()
	}
	return &Post{
		fileName:   fileName,
		workingDir: targetPath,
		kind:       "posts",
	}
}

func (p *Post) save() error {
	postPath := JoinPath(p.workingDir, "content/"+p.kind)
	exist, err := DirExists(postPath)
	if !exist {
		Println("New post cannot be created.")
		return err
	}
	var output string
	output, err = ExecHugo("new "+p.kind+"/"+p.fileName, p.workingDir)
	Println(output)
	return err
}

func (p *Post) getFilePath() string {
	filePath := JoinPath(p.workingDir, "/content/", p.kind, p.fileName)
	Println(filePath)
	return filePath
}

func (p *Post) appendFortune() error {
	str, err := GetFortune()
	if err != nil {
		return err
	}
	AppendToFile(p.getFilePath(), str)
	return nil
}
