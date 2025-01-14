package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct{
  DbUrl string `json:"db_url"`
  CurrentUser string `json:"current_user_name"`
}


func Read() (Config, error){

  conf := Config {}

  homeDir, err := os.UserHomeDir()
  if err != nil {
    return conf, err
  }
  configLocation := fmt.Sprintf("%v/blog-aggregatorconfig.json", homeDir)

  jsonFile, err := os.Open(configLocation)
  if err != nil {
    return conf, err
  }
  defer jsonFile.Close()

  byteValue, err := io.ReadAll(jsonFile)
  if err != nil {
    return conf, err
  }

  err = json.Unmarshal(byteValue, &conf)
  if err != nil {
    return conf, err
  }

  return conf, nil
}
