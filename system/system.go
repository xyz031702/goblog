package system

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

//Configs contains application configurations for all gin modes
type Configs struct {
	Debug   Config
	Release Config
	Test    Config
}

//Config contains application configuration for active gin mode
type Config struct {
	Database      DatabaseConfig
	SessionSecret string `json:"session_secret"`
}

//DatabaseConfig contains database connection info
type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string //database name
	User     string
	Password string
}

//current loaded config
var config *Config

//LoadConfig unmarshals config for current GIN_MODE
func LoadConfig() {
	configs := &Configs{}

	file := ""
	switch gin.Mode() {
	case gin.DebugMode:
		file = "dev"
	case gin.ReleaseMode:
		file = "release"
	default:
		panic(fmt.Sprintf("Unknown gin mode %s", gin.Mode()))
	}

	data, err := ioutil.ReadFile("config/config_" + file + ".json")
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		panic(err)
	}

	if err := json.Unmarshal(data, configs); err != nil {
		panic(err)
	}
}

//GetConfig returns actual config
func GetConfig() *Config {
	return config
}
