package bidder

import (
	"io/ioutil"
	"encoding/json"
	"os"
)

type Config struct {
	Server string `json:"server"`
	Port string  `json:"port"`
}


func ReadConfig(configFile string) (*Config, error) {

	if _, err := os.Stat(configFile); os.IsNotExist(err){
		return nil, err
	}

	f, err := ioutil.ReadFile(configFile)

	if err != nil{
		return nil, err
	}
	var res *Config
	e := json.Unmarshal(f, &res)

	if e != nil{
		return nil, e
	}
	return res, nil
}