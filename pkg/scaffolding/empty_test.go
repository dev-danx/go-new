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
