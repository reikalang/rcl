package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"

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
	echoCmd := &cobra.Command{
		Use:   "echo",
		Short: "parse and pretty print",
		Long:  "Parse the file and print with indent",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("TODO: real echo")
			// TODO: grab one file, print with indent
		},
	}
	root.AddCommand(echoCmd)
	if err := root.Execute(); err != nil {
		Err(err)
	}
}

func init() {
	log.AddChild(rcl.Registry)
}

func Err(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
