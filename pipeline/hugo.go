package pipeline

import (
	"errors"
	"wiredcraft-hugo/externals"
)

var ArgError = errors.New("Only accept one argument: dev or staging")

func checkArgs(args []string) (string, error) {
	if len(args) != 1 {
		return "", ArgError
	}
	envName := args[0]
	if envName != "dev" && envName != "staging" {
		return "", ArgError
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
	site := newSite(envName)
	site.incrementVersion()
	site.compile()
	site.release() // git commit, tag and push
	return nil
}
