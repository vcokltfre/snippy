package client

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/atotto/clipboard"
	"github.com/urfave/cli/v2"
)

var Download = &cli.Command{
	Name:      "download",
	Usage:     "Download a snippet",
	ArgsUsage: "<id>",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "copy",
			Aliases: []string{"c"},
			Usage:   "Copy the snippet to the clipboard",
		},
	},
	Action: func(c *cli.Context) error {
		conf := LoadConfig()
		path := conf.Base + "/snippets/" + c.Args().First() + "/raw"

		resp, err := http.Get(path)
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			return fmt.Errorf("failed to download snippet: %s", resp.Status)
		}

		var content bytes.Buffer
		_, err = content.ReadFrom(resp.Body)
		if err != nil {
			return err
		}

		if !c.Bool("copy") {
			fmt.Println(content.String())
			return nil
		}

		err = clipboard.WriteAll(content.String())
		if err != nil {
			return err
		}

		fmt.Println("Copied snippet to clipboard.")

		return nil
	},
}
