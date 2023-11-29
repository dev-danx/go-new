package scaffolding

import (
	"fmt"
	"testing"
)

func Test_emptyProject_readFiles(t *testing.T) {
	p := &emptyProject{}
	result := p.readFiles()
	fmt.Print(result)
	t.Fail()
}

func Test_emptyProject_CreateNew(t *testing.T) {
	t.Skip()
	p := &emptyProject{}
	p.CreateNew("TESTING")
}
