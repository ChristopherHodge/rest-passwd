package config

import (
  "encoding/json"
  "io/ioutil"
  "rest-passwd/logger"
)

// App configuration
type Config struct {
  Port string        // Listening port
  PasswdFile string  // Path to passwd file
  GroupFile string   // Path to group file
}

var log = logger.Get()

const DEFAULT_PORT = "8080"
const DEFAULT_PASSWD = "/etc/passwd"
const DEFAULT_GROUP = "/etc/group"

// persistent config data
var configData = &Config{}

// use local scope to allow mock in tests
var ReadFile = ioutil.ReadFile

// Read the config file
func Read(file string) (*Config) {
  data, err := ReadFile(file) ; if err != nil {
    log.Warn("opening config: ", err)
    log.Warn("using default config values")
    *configData = GetDefaults()
    return configData
  }

  err = json.Unmarshal(data, &configData) ; if err != nil {
    log.Warn("loading config: ", err)
    log.Warn("using default config values")
    *configData = GetDefaults()
  }

  return configData
}

func Get() (*Config) {
  return configData
}

func GetDefaults() (Config) {
  return Config {DEFAULT_PORT, DEFAULT_PASSWD, DEFAULT_GROUP}
}
