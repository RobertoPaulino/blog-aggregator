package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const configFileName = ".blog-aggregatorconfig.json"

type Config struct{
  DbUrl string `json:"db_url"`
  CurrentUser string `json:"current_user_name"`
}


func (c *Config) SetUser() {
  
}

func getConfigFilePath() (string, error){

  homeDir, err := os.UserHomeDir()
  if err != nil {
    return "", err
  }
  configLocation := fmt.Sprintf("%v%v", homeDir, configFileName)

  return configLocation, nil
}

func write(conf Config) error {
  // conf -> json | get conf location in home directory | write new json to that location
  configPath, err := getConfigFilePath()
  if err != nil {
    return err 
  }

  jsonData, err := json.Marshal(conf)
  if err != nil {
    return err
  }
  

  
  return nil
}

func Read() (Config, error){

  conf := Config {}

  configLocation, err := getConfigFilePath()
  if err != nil {
    return conf, err
  }

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
