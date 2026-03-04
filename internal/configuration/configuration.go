package configuration

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type config struct {
	Webserver struct {
		Host string `json:"Host"`
		Port int    `json:"Port"`
	}
	Host        string
	HostAddress string `json:"HostAddress"`
}

var Configuration config

func (c *config) Init(path string) error {
	configFile, err := os.Open(path)
	if err != nil {
		panic("Could not read config file")
	}
	defer configFile.Close()

	jsonBytes, err := io.ReadAll(configFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonBytes, c)
	if err != nil {
		return err
	}

	if c.Webserver.Port == 80 {
		c.Host = c.Webserver.Host
	} else {
		c.Host = fmt.Sprintf("%s:%d", c.Webserver.Host, c.Webserver.Port)
	}

	return nil
}
