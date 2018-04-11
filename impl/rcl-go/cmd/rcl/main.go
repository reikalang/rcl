package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	icli "github.com/at15/go.ice/ice/cli"
	dlog "github.com/dyweb/gommon/log"

	"github.com/reikalang/rcl/impl/rcl-go/rcl"
)

const (
	myname = "rcl" // reika configuration language
)

var log = dlog.NewApplicationLogger()

var (
	version   string
	commit    string
	buildTime string
	buildUser string
	goVersion = runtime.Version()
)

var buildInfo = icli.BuildInfo{Version: version, Commit: commit, BuildTime: buildTime, BuildUser: buildUser, GoVersion: goVersion}
var cli *icli.Root

func main() {
	cli = icli.New(
		icli.Name(myname),
		icli.Description("Reika Configuration Language"),
		icli.Version(buildInfo),
		icli.LogRegistry(log),
	)
	root := cli.Command()
	root.AddCommand(echoCmd)
	if err := root.Execute(); err != nil {
		Err(err)
	}
}

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "parse and pretty print",
	Long:  "Parse the file and print with indent",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("must provide at least one file")
			return
		}
		for _, fp := range args {
			// just warn about invalid extension, and keep going
			if !strings.HasSuffix(fp, rcl.DotExtension) {
				log.Warnf("%s has invalid extension", fp)
			}
			input, err := antlr.NewFileStream(fp)
			if err != nil {
				log.Warnf("antlr: %s", err)
				continue
			}
			lexer := rcl.NewRCLLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := rcl.NewRCLParser(stream)
			p.BuildParseTrees = true
			// TODO: error listener, it just log to stdout by default ...
			tree := p.Rcl()
			// TODO: visitor should have some methods for exposing error ...
			visitor := rcl.NewEchoVisitor()
			tree.Accept(visitor)
		}
	},
}

func init() {
	log.AddChild(rcl.Registry)
}

func Err(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
