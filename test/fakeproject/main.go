package main

import (
	"github.com/ermos/gomon/test/fakeproject/internal/hello"
	"os"
)

func main() {
	// Modify SayHello and see what happen :)
	hello.SayHello(os.Args)
}
