package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/src-d/enry/v2"
	"github.com/urfave/cli/v2"
)

type uploadSnippet struct {
	Language string `json:"language"`
	Content  string `json:"content"`
}

var Upload = &cli.Command{
	Name:      "upload",
	Usage:     "Upload a snippet",
	ArgsUsage: "<id> <file>",
	Action: func(c *cli.Context) error {
		conf := LoadConfig()
		path := conf.Base + "/snippets/" + c.Args().First()

		data, err := os.ReadFile(c.Args().Get(1))
		if err != nil {
			return err
		}
		lang, _ := enry.GetLanguageByContent(c.Args().Get(1), data)

		jsonData, err := json.Marshal(uploadSnippet{
			Language: lang,
			Content:  string(data),
		})
		if err != nil {
			return err
		}

		req, err := http.NewRequest("POST", path, bytes.NewReader(jsonData))
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
			return fmt.Errorf("failed to upload snippet: %s", resp.Status)
		}

		fmt.Println("Uploaded snippet successfully to " + path)

		return nil
	},
}
