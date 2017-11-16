package bidder

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Server      string `json:"server"`
	Port        string `json:"port"`
	ShortSeries int    `json:"short_series"`
	LongSeries  int    `json:"long_series"`
}

func ReadConfig(configFile string) (*Config, error) {

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return nil, err
	}

	f, err := ioutil.ReadFile(configFile)

	if err != nil {
		return nil, err
	}

	var res *Config
	err = json.Unmarshal(f, &res)

	if err != nil {
		return nil, err
	}
	return res, nil
}
