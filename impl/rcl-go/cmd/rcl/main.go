package main

import (
	"fmt"
	"log"
	"os"

	"github.com/reikalang/rcl/impl/rcl-go/rcl"
)

// TODO: commands
// generate ? generate the BlaDecoded struct for decoding into a struct with position information
// json ---encode convert rcl format to json
// json --decode convert from json to rcl

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
