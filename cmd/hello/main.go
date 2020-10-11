package main

import (
	"flag"

	"github.com/rolandjitsu/go-cross/pkg/hello"
)

func main() {
	msgpack := flag.Bool("msgpack", false, "Use msgpack")
	flag.Parse()

	name := flag.Arg(0)
	if name == "" {
		name = "stranger"
	}

	if *msgpack {
		err := hello.GreetMsgPack(name)
		if err != nil {
			panic(err)
		}
	} else {
		hello.Greet(name)
	}
}
