package externals

// DependenciesInterface just for mocking
type DependenciesInterface interface {
	AppendToFile(filePath string, content string) error
	GetFortune() (string, error)
	JoinPath(elem ...string) string
	ExecHugo(argString string, workingDir string) (string, error)
	GetNow() string
	GetHugoWorkingDir() string
	GetWorkingDir() string
	DirExists(path string) (bool, error)
	Println(a ...interface{})
	ReadFileContent(filename string) ([]byte, error)
	WriteFile(filename string, data []byte) error
	AddCommitAndPush(message string, workingDir string) (string, error)
	GitTagAndPush(version string, workingDir string) error
}
