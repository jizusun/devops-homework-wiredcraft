package pipeline

import (
	"errors"
	"wiredcraft-hugo/externals"
)

// Execute accepts only one parameter and executes depends on if it is `dev` or `staging`
func Execute(args []string) error {
	dep := &externals.Dependencies{}
	argErr := errors.New("Only accept one argument: dev or staging")
	if len(args) != 1 {
		return argErr
	}
	envName := args[0]
	if envName != "dev" && envName != "staging" {
		return argErr
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
