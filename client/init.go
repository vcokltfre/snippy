package client

import "github.com/urfave/cli/v2"

var Init = &cli.Command{
	Name:      "init",
	Usage:     "Initialise the client",
	ArgsUsage: "<base_url> <auth>",
	Action: func(c *cli.Context) error {
		baseUrl := c.Args().First()
		auth := c.Args().Get(1)

		conf := config{
			Base: baseUrl,
			Auth: auth,
		}
		SaveConfig(&conf)

		return nil
	},
}
