package client

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli/v2"
)

var Delete = &cli.Command{
	Name:      "delete",
	Usage:     "Delete a snippet",
	ArgsUsage: "<id>",
	Action: func(c *cli.Context) error {
		conf := LoadConfig()
		path := conf.Base + "/snippets/" + c.Args().First()

		req, err := http.NewRequest("DELETE", path, nil)
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", conf.Auth)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			return fmt.Errorf("failed to delete snippet: %s", resp.Status)
		}

		fmt.Println("Deleted snippet successfully.")

		return nil
	},
}
