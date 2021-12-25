package main

import (
	"fmt"
	"testing"

	"github.com/software-engineering-components/go-arch-lab4/commands"
	"github.com/software-engineering-components/go-arch-lab4/engine"
)

var data = "test benchmark programm"
var cntRes engine.Command

func BenchmarkCount(b *testing.B) {
	const baseLen = 3000

	for i := 0; i < 20; i++ {
		input := data
		l := baseLen * (i + 1)

		for j := 0; j < l; j++ {
			input = input + "somerandomlineinput"
		}

		b.Run(fmt.Sprintf("len=%d", l), func(b *testing.B) {
			cntRes = commands.Parse(input)
		})
	}
}
