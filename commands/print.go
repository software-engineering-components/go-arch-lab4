package commands

import (
	"fmt"

	"github.com/software-engineering-components/go-arch-lab4/engine"
)

type print struct {
	Arg string
}

func (p *print) Execute(loop engine.Handler) {
	fmt.Println(p.Arg)
}
