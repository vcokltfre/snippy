package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"github.com/vcokltfre/snippy/server"
)

var app = &cli.App{
	Name:  "snippy",
	Usage: "A simple way to store and manage snippets of code",
	Commands: []*cli.Command{
		server.ServerCommand,
	},
}

func init() {
	godotenv.Load()
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			panic(err)
		}
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}
