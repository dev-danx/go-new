package scaffolding

import (
	"embed"
	"fmt"
	"github.com/dev-danx/go-new/pkg/contentReader"
)

type emptyProject struct {
}

func (p *emptyProject) readFiles() map[string]string {
	files = make(map[string]string)
	files["cmd/main.go"] = contentReader.ReadAsString(res, "fileTemplates/empty/main.go")
	files["Dockerfile"] = contentReader.ReadAsString(res, "buildFiles/Dockerfile")
	files["makefile"] = contentReader.ReadAsString(res, "buildFiles/makefile")
	return files
}

var (
	//go:embed fileTemplates/empty/*
	//go:embed buildFiles/*
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
	p.readFiles()
	initProject(name, files)
	fmt.Println("Project is Created, good luck on our project!")
}
