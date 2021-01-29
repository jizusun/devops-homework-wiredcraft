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
	GetHugoConfigToml(workingDir string) (string, error)
}
