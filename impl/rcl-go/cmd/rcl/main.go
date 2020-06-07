package main

import (
	"fmt"
	"log"
	"os"

	"github.com/reikalang/rcl/impl/rcl-go/rcl"
)

func main() {
	fmt.Println("a new rcl is on the way")
	// TODO: make it a repl? might add extra go.mod so rcl itself won't include extra dependency
	arg := os.Args[1]
	p := rcl.NewParser(arg)
	n, err := p.Parse()
	if err != nil {
		log.Fatal(err)
	}
	n.Pos()
}
