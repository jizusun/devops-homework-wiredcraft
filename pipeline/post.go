package pipeline

// Post the hugo post
type Post struct {
	title    string
	filePath string
}

func newPost() *Post {
	return &Post{}
}

func (p *Post) save() error {
	var err error
	return err
}

func (p *Post) updateContent() {

}
