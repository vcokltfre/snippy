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

	e.GET("/", func(c echo.Context) error {
		return c.HTML(200, RenderIndex())
	})

	e.GET("/snippets/:id", func(c echo.Context) error {
		return c.HTML(200, RenderSnippet("go", "func main() {}", c.Param("id")))
	})

	e.GET("/snippets/:id/raw", func(c echo.Context) error {
		return c.String(200, "func main() {}")
	})

	return e.Start(bind)
}
