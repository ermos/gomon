package main

import (
	"github.com/ermos/cli"
	"github.com/ermos/gomon/cmd"
)

func main() {
	cli.Init("gomon", "Rebuild your application when file changes in the directory are detected.  📦")
	cli.AddOption(
		"ext",
		"List of file's extensions that allows to reload your application (separate with comma)",
		"list",
		).AddShortName("e")
	cli.AddOption(
		"dir",
		"List of dirs contains files that allows to reload your application (separate with comma)",
		"list",
	).AddShortName("d")
	cli.AddAction(cmd.DevHandler{}, "dev", "...args")
	cli.Run()
}
