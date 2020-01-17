package main

import (
	"github.com/urfave/cli"
	"az-util/cmd/server"
	"az-util/cmd/tool"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "az-util"
	app.Commands = []cli.Command{
		server.Server,
		tool.InitDB,
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
