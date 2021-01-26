package pipeline

// Execute accepts only one parameter and executes depends on if it is `dev` or `staging`
func Execute(envName string) error {
	// TODO: return error if no such an argument was provided
	// TODO: return error if `envName` is unknown: `dev`, `staging`

	if envName == "dev" {
		post := newPost()
		var err error
		err = post.save()
		if err != nil {
			return err
		}
		post.updateContent()
	}

	site := newSite(envName)
	site.incrementVersion()
	site.compile()
	site.release() // git commit, tag and push
	return nil
}
