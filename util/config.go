package util

import (
	"github.com/Unknwon/goconfig"
)

func GetServerPath() (string, error) {
	config, err := goconfig.LoadConfigFile("./config/config.ini")
	if err != nil {
		return "", err
	}
	path, err := config.GetValue("proxy_server", "path")
	if err != nil {
		return "", err
	}
	return path, nil
}

func GetServerPort() (string, error) {
	config, err := goconfig.LoadConfigFile("./config/config.ini")
	if err != nil {
		return "", err
	}
	path, err := config.GetValue("proxy_server", "port")
	if err != nil {
		return "", err
	}
	return path, nil
}
