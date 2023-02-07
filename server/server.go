package server

import (
	"os"

	"github.com/labstack/echo"
	"github.com/urfave/cli/v2"
)

var ServerCommand = &cli.Command{
	Name:      "server",
	Usage:     "Start the server",
	ArgsUsage: "[bind]",
	Action: func(c *cli.Context) error {
		bind := c.Args().First()
		if bind == "" {
			if os.Getenv("SNIPPY_BIND") != "" {
				bind = os.Getenv("SNIPPY_BIND")
			} else {
				bind = ":8080"
			}
		}

		return start(bind)
	},
}

func start(bind string) error {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	return e.Start(bind)
}
