package scaffolding

import (
	"embed"
	"fmt"
)

type emptyProject struct {
}

var (
	//go:embed fileTemplates/empty
	//go:embed buildFiles
	res embed.FS

	files map[string]string
)

func newEmpty() ProjectScaffolding {
	return &emptyProject{}
}

func (p *emptyProject) Name() string {
	return "Empty"
}

func (p *emptyProject) CreateNew(name string) {
	fmt.Println("Creating Empty Project")
	initProject(name, files)
	fmt.Println("Project is Created, good luck on our project!")
}
