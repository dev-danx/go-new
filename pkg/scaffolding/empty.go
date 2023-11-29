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
	files["cmd/main.go"] = contentReader.ReadAsString(goFiles, "fileTemplates/empty/main.go")
	files["Dockerfile"] = contentReader.ReadAsString(buildFiles, "buildFiles/Dockerfile")
	files["makefile"] = contentReader.ReadAsString(buildFiles, "buildFiles/makefile")
	return files
}

var (
	//go:embed fileTemplates/empty/*
	goFiles embed.FS
	//go:embed buildFiles/*
	buildFiles embed.FS

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
	files = p.readFiles()
	initProject(name, files)
	fmt.Println("Project is Created, good luck on our project!")
}
