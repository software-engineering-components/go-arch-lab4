package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/software-engineering-components/go-arch-lab4/commands"
	"github.com/software-engineering-components/go-arch-lab4/engine"
)

func main() {
	path := "example.txt"
	file := flag.String("f", path, "Example file")
	flag.Parse()

	input, err := os.Open(*file)
	defer input.Close()
	if err != nil {
		log.Fatalf("error occured while opening file %s", err)
		return
	}

	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		commandLine := scanner.Text()
		cmd := commands.Parse(commandLine)

		eventLoop.Post(cmd)
	}

	eventLoop.AwaitFinish()
}
