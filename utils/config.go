package utils

import (
	"log"

	"github.com/Unknwon/goconfig"
)

type Config struct {
	File *goconfig.ConfigFile
}

func NewConfig() *Config {
	config, err := goconfig.LoadConfigFile("./config/config.ini")
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Config{
		File: config,
	}
}

func (c *Config) GetServerPath() (string, error) {
	path, err := c.File.GetValue("proxy_server", "path")
	if err != nil {
		return "", err
	}
	return path, nil
}

func (c *Config) GetServerPort() (string, error) {
	path, err := c.File.GetValue("proxy_server", "port")
	if err != nil {
		return "", err
	}
	return path, nil
}

func (c *Config) GetUsernameAndPassword() (string, string, error) {
	username, err := c.File.GetValue("proxy_server", "username")
	if err != nil {
		return "", "", err
	}
	password, err := c.File.GetValue("proxy_server", "password")
	if err != nil {
		return "", "", err
	}
	return username, password, nil
}

func (c *Config) GetServicesFilePath() (string, error) {
	serviceFile, err := c.File.GetValue("proxy_server", "services")
	if err != nil {
		return "", err
	}
	return serviceFile, nil
}

func (c *Config) UpdateAutoStart(autoStart string)(isSuccess bool){
	isSuccess = c.File.SetValue("config", "auto_start", autoStart)
	return
}

func (c *Config) GetAutoStart()(autoStart bool){
	autoStart = c.File.MustBool("config", "auto_start")
	return
}