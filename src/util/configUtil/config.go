package configUtil

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var Conf = Configuration{}

var ErrorCode = 1

type Configuration struct {
	Listen        string `yaml:"listen"`
	LogLevel      string `yaml:"logLevel"`
	LogFile       string `yaml:"logFile"`
	RouterLogFile string `yaml:"routerLogFile"`
	LogColor      bool   `yaml:"logColor"`
	LogMaxSize    int    `yaml:"logMaxSize"`
	LogMaxBackups int    `yaml:"logMaxBackups"`
	LogMaxAge     int    `yaml:"logMaxAge"`
}

func getConfigFile() (string, error) {
	var cfgFile = "src/koala.yaml"

	if cfgFile == "" {
		var err error
		cfgFile, err = filepath.Abs(os.Args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return cfgFile, err
			//os.Exit(ErrorCode)
		}
		cfgFile += ".yaml"
	}
	return cfgFile, nil
}

func (c *Configuration) InitConfig() (*Configuration, error) {
	cfgFile, err := getConfigFile()
	log.Println("cfgFile:", cfgFile)
	if err != nil {
		log.Println("\tERROR\tFailed to get config file.")
		return nil, err
	}

	yamlFile, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		log.Printf("\tERROR\tyamlFile.Get err   #%v ", err)
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return nil, err
	}
	log.Println("\tINFO\tConfig init successfully.")
	return &Conf, nil
}

func GetLogLevel() string {
	if len(Conf.LogLevel) < 0 {
		Conf.LogLevel = "info"
	}
	return strings.ToLower(Conf.LogLevel)
}

func GetLogFile() string {
	if len(Conf.LogFile) < 0 {
		Conf.LogFile = "./logs/koala.log"
	}
	return Conf.LogFile
}

func GetLogColor() bool {
	return Conf.LogColor
}

func GetLogMaxSize() int {
	if Conf.LogMaxSize <= 0 {
		Conf.LogMaxSize = 100
	}
	return Conf.LogMaxSize
}

func GetLogMaxBackups() int {
	if Conf.LogMaxBackups <= 0 {
		Conf.LogMaxBackups = 50
	}
	return Conf.LogMaxBackups
}

func GetLogMaxAge() int {
	if Conf.LogMaxAge <= 0 {
		Conf.LogMaxAge = 30
	}
	return Conf.LogMaxAge
}

func GetListen() string {
	if len(Conf.Listen) <= 0 {
		Conf.Listen = "8088"
	}
	return Conf.Listen
}

func GetRouterLogFile() string {
	if len(Conf.RouterLogFile) <= 0 {
		Conf.RouterLogFile = "./logs/router.log"
	}
	return Conf.RouterLogFile
}
