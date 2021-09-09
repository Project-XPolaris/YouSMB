package application

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var Config AppConfig

type AppConfig struct {
	Addr          string `json:"addr"`
	RPC           string `json:"rpc"`
	SmbConfigPath string `json:"smb_config_path"`
	RestartScript string `json:"restart_script"`
	StartScript   string `json:"start_script"`
	StopScript    string `json:"stop_script"`
}

func LoadAppConfig() error {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	raw, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(raw, &Config)
	return err
}
