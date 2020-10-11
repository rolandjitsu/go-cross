package main

import (
	"flag"

	"github.com/rolandjitsu/go-cross/pkg/hello"
)

func main() {
	flag.Parse()

	name := flag.Arg(0)
	if name == "" {
		name = "stranger"
	}

	hello.Greet(name)
}
