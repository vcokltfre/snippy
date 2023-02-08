package server

import (
	"os"

	"github.com/labstack/echo"
	"github.com/urfave/cli/v2"
	"github.com/vcokltfre/snippy/server/database"
)

type snippetRequest struct {
	Language string `json:"language"`
	Content  string `json:"content"`
}

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

	db, err := database.Connect()
	if err != nil {
		return err
	}

	e.HideBanner = true
	e.HidePort = true

	e.GET("/", func(c echo.Context) error {
		snippets, err := db.GetSnippets()
		if err != nil {
			return err
		}

		return c.HTML(200, RenderIndex(snippets))
	})

	e.GET("/snippets/:id", func(c echo.Context) error {
		snippet, err := db.GetSnippet(c.Param("id"))
		if err != nil {
			return c.NoContent(404)
		}

		return c.HTML(200, RenderSnippet(snippet.Language, snippet.Content, c.Param("id")))
	})

	e.GET("/snippets/:id/raw", func(c echo.Context) error {
		snippet, err := db.GetSnippet(c.Param("id"))
		if err != nil {
			return c.NoContent(404)
		}

		return c.String(200, snippet.Content)
	})

	e.POST("/snippets/:id", func(c echo.Context) error {
		sr := snippetRequest{}
		if err := c.Bind(&sr); err != nil {
			return err
		}

		err := db.UpsertSnippet(database.Snippet{
			ID:       c.Param("id"),
			Language: sr.Language,
			Content:  sr.Content,
		})
		if err != nil {
			return err
		}

		return c.NoContent(200)
	}, requireAuth)

	e.DELETE("/snippets/:id", func(c echo.Context) error {
		err := db.DeleteSnippet(c.Param("id"))
		if err != nil {
			return err
		}

		return c.NoContent(200)
	}, requireAuth)

	return e.Start(bind)
}
