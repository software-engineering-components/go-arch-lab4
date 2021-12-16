package commands

import (
	"github.com/software-engineering-components/go-arch-lab4/engine"
)

type reverse struct {
	Arg string
}

func (p *reverse) Execute(loop engine.Handler) {
	var reversed string
	for _, v := range p.Arg {
		reversed = string(v) + reversed
	}
	loop.Post(&print{Arg: reversed})
}
