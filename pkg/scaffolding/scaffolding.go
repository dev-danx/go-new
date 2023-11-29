package scaffolding

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type ProjectScaffolding interface {
	Name() string
	CreateNew(projectName string)
	readFiles() map[string]string
}

func ProjectList() []ProjectScaffolding {
	result := make([]ProjectScaffolding, 0)
	result = append(result, newEmpty())
	return result
}

func makeFolder(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func makeFile(name, content string) {
	file, err := os.Create(name)

	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(content)
}

func runShellCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)

	// The `Output` method executes the command and
	// collects the output, returning its value
	_, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}
}

func initProject(name string, files map[string]string) {
	makeFolder("cmd")
	makeFolder("pkg")
	makeFolder("internal")
	mainFile := "cmd/main.go"
	makeFile(mainFile, files[mainFile])
	dockerFile := "Dockerfile"
	makefile := "makefile"
	makeFile(dockerFile, files[dockerFile])
	makeFile(makefile, files[makefile])
	runShellCommand("go", "mod", "init", name)
}
