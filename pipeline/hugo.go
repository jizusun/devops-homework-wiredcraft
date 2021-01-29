package pipeline

import (
	"errors"

	"github.com/jizusun/wiredcraft-hugo/externals"
)

// ErrArguments the error when the argument is incorrect
var ErrArguments = errors.New("Only accept one argument: dev or staging")

func checkArgs(args []string) (string, error) {
	if len(args) != 1 {
		return "", ErrArguments
	}
	envName := args[0]
	if envName != "dev" && envName != "staging" {
		return "", ErrArguments
	}
	return envName, nil
}

// Execute accepts only one parameter and executes depends on if it is `dev` or `staging`
func Execute(args []string) error {
	dep := &externals.Dependencies{}
	envName, err := checkArgs(args)
	if err != nil {
		return err
	}
	if envName == "dev" {
		post := newPost(dep)
		err := post.save(dep)
		if err != nil {
			return err
		}
		err = post.appendFortune(dep)
		if err != nil {
			return err
		}
	}
	site := newSite(envName, dep)
	site.incrementVersion(dep)
	site.compile()
	site.release() // git commit, tag and push
	return nil
}
