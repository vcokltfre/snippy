package client

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type config struct {
	Base string `json:"base"`
	Auth string `json:"auth"`
}

func SaveConfig(conf *config) {
	data, err := json.Marshal(conf)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(path.Join(os.Getenv("HOME"), ".snippy"), data, 0644)
	if err != nil {
		panic(err)
	}
}

func LoadConfig() *config {
	_, err := os.Stat(path.Join(os.Getenv("HOME"), ".snippy"))
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Please run `snippy init`")
			os.Exit(1)
		}
		panic(err)
	}

	data, err := os.ReadFile(path.Join(os.Getenv("HOME"), ".snippy"))
	if err != nil {
		panic(err)
	}
	var conf config
	err = json.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}
	return &conf
}
