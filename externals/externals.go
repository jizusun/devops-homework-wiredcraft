package externals

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/spf13/afero"
)

type Dependencies struct{}

var fs = afero.NewOsFs()
var afs = &afero.Afero{Fs: fs}

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

func (dep Dependencies) GetFortune() (string, error) {
	out, err := exec.Command("fortune").Output()
	if err != nil {
		return "", err
	}
	outputStr := string(out[:])
	fmt.Println(outputStr)
	return outputStr, nil
}

func (dep Dependencies) JoinPath(elem ...string) string {
	return path.Join(elem...)
}

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

func (dep Dependencies) GetNow() string {
	return time.Now().Format(time.RFC3339)
}

func (dep Dependencies) GetHugoWorkingDir() string {
	return os.Getenv("HUGO_WORK_DIR")
}

func (dep Dependencies) GetWorkingDir() string {
	dir, _ := os.Getwd()
	return dir
}

func (dep Dependencies) DirExists(path string) (bool, error) {
	return afs.DirExists(path)
}

func (dep Dependencies) Println(a ...interface{}) {
	fmt.Println(a...)
}
