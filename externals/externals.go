package externals

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/spf13/afero"
)

// Dependencies dep
type Dependencies struct{}

var fs = afero.NewOsFs()
var afs = &afero.Afero{Fs: fs}

// AppendToFile append content to the specified file
func (dep Dependencies) AppendToFile(filePath string, content string) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(content); err != nil {
		return err
	}
	return nil
}

// GetFortune get content from the command fortune
func (dep Dependencies) GetFortune() (string, error) {
	out, err := exec.Command("fortune").Output()
	if err != nil {
		return "", err
	}
	outputStr := string(out[:])
	fmt.Println(outputStr)
	return outputStr, nil
}

// JoinPath wrapper for path.Join
func (dep Dependencies) JoinPath(elem ...string) string {
	return path.Join(elem...)
}

// ExecHugo execute the hugo command
// https://stackoverflow.com/questions/18970265/is-there-an-easy-way-to-stub-out-time-now-globally-during-test
// https://talks.golang.org/2012/10things.slide#8
func (dep Dependencies) ExecHugo(argString string, workingDir string) (string, error) {
	args := strings.Split(argString, " ")
	cmd := exec.Command("hugo", args...)
	cmd.Dir = workingDir
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	output := string(out[:])
	return output, nil
}

// GetNow wrapper for time.Now
func (dep Dependencies) GetNow() string {
	return time.Now().Format(time.RFC3339)
}

// GetHugoWorkingDir get HUGO_WORK_DIR
func (dep Dependencies) GetHugoWorkingDir() string {
	return os.Getenv("HUGO_WORK_DIR")
}

// GetWorkingDir wrapper for os.Getwd
func (dep Dependencies) GetWorkingDir() string {
	dir, _ := os.Getwd()
	return dir
}

// DirExists wrapper for afs.DirExists
func (dep Dependencies) DirExists(path string) (bool, error) {
	return afs.DirExists(path)
}

// Println wrapper for fmt.Println
func (dep Dependencies) Println(a ...interface{}) {
	fmt.Println(a...)
}

// ReadFileContent wrapper for ioutil.Readfile
func (dep Dependencies) ReadFileContent(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// WriteFile wrapper for ioutil.WriteFile
func (dep Dependencies) WriteFile(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0644)
}

func execCommandInDir(workingDir string, name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	cmd.Dir = workingDir
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	output := string(out[:])
	return output, nil
}

// GitTagAndPush tag the version and push
func (dep Dependencies) GitTagAndPush(version string, workingDir string) error {
	_, err := execCommandInDir(workingDir, "git", "tag", version)
	if err != nil {
		return err
	}
	fmt.Println("git tag " + version)
	_, err = execCommandInDir(workingDir, "git", "push", "origin", version)
	if err != nil {
		return err
	}
	fmt.Println("git push origin " + version)
	return nil
}

// AddCommitAndPush git add -A, git commit, git push
func (dep Dependencies) AddCommitAndPush(message string, workingDir string) (string, error) {
	fmt.Println("Working dir: " + workingDir)
	output, err := execCommandInDir(workingDir, "git", "add", ".")
	if err != nil {
		return "", err
	}
	fmt.Println(output)

	output, err = execCommandInDir(workingDir, "git", "commit", "-m", message)
	if err != nil {
		return "", err
	}
	fmt.Println(output)

	output, err = execCommandInDir(workingDir, "git", "pull", "--rebase")
	if err != nil {
		return "", err
	}

	output, err = execCommandInDir(workingDir, "git", "push")
	if err != nil {
		return "", err
	}
	fmt.Println("git push finished.")

	commitID, err := execCommandInDir(workingDir, "git", "rev-parse", "HEAD")
	return commitID, err
}
