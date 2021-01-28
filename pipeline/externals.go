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

func AppendToFile(filePath string, content string) error {
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

func GetFortune() (string, error) {
	out, err := exec.Command("fortune").Output()
	if err != nil {
		return "", err
	}
	outputStr := string(out[:])
	fmt.Println(outputStr)
	return outputStr, nil
}

func JoinPath(elem ...string) string {
	return path.Join(elem...)
}

// https://stackoverflow.com/questions/18970265/is-there-an-easy-way-to-stub-out-time-now-globally-during-test
// https://talks.golang.org/2012/10things.slide#8
func ExecHugo(argString string, workingDir string) (string, error) {
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

func GetNow() string {
	return time.Now().Format(time.RFC3339)
}

func GetHugoWorkingDir() string {
	return os.Getenv("HUGO_WORK_DIR")
}

func GetWorkingDir() string {
	dir, _ := os.Getwd()
	return dir
}

func DirExists(path string) (bool, error) {
	return afs.DirExists(path)
}

func Println(a ...interface{}) (int, error) {
	return fmt.Println(a...)
}
