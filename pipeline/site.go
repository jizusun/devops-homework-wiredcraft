package pipeline

// Site the hugo site
type Site struct {
	version string
	envName string
}

func newSite(envName string) *Site {
	return &Site{envName: envName}
}

func (s *Site) incrementVersion() {

}

func (s *Site) compile() {

}

func (s *Site) release() {

}
