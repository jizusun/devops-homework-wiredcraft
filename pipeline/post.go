package pipeline

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/spf13/afero"
)

var fs = afero.NewOsFs()
var afs = &afero.Afero{Fs: fs}

// Post the hugo post
type Post struct {
	fileName   string
	workingDir string
	kind       string
}

func newPost() *Post {
	timestamp := time.Now().Format(time.RFC3339)
	fileName := timestamp + ".md"
	targetPath := os.Getenv("HUGO_WORK_DIR")
	if len(targetPath) == 0 {
		targetPath, _ = os.Getwd()
	}
	return &Post{
		fileName:   fileName,
		workingDir: targetPath,
		kind:       "posts",
	}
}

func (p *Post) save() error {
	postPath := path.Join(p.workingDir, "content/"+p.kind)
	exist, err := afs.DirExists(postPath)
	if !exist {
		fmt.Println("New post cannot be created.")
		return err
	}
	var output string
	output, err = p.execHugo("new "+p.kind+"/"+p.fileName, realExec{})
	fmt.Println(output)
	return err
}

type Executor interface {
	Command(name string, arg ...string) *Cmder
}
type realExec struct {
}

func (realExec) Command(name string, arg ...string) *exec.Cmd {
	return exec.Command(name, arg...)
}

type Cmder interface {
	Output() ([]byte, error)
}

// https://stackoverflow.com/questions/18970265/is-there-an-easy-way-to-stub-out-time-now-globally-during-test
// https://talks.golang.org/2012/10things.slide#8
type realCmd struct {
	Dir string
}

func (rc realCmd) Output() ([]byte, error) {
	return rc.Output()
}

func (p *Post) execHugo(argString string, exec realExec) (string, error) {
	args := strings.Split(argString, " ")
	cmd := exec.Command("hugo", args...)
	cmd.Dir = p.workingDir
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	output := string(out[:])
	return output, nil
}

func (p *Post) appendFortune() error {
	out, err := exec.Command("fortune").Output()
	if err != nil {
		return err
	}
	outputStr := string(out[:])
	fmt.Println(outputStr)

	filePath := path.Join(p.workingDir, "/content/", p.kind, p.fileName)
	fmt.Println(filePath)
	var f *os.File
	f, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(out); err != nil {
		return err
	}
	return nil
}
