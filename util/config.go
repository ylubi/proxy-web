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

func GetUsernameAndPassword() (string, string, error) {
	config, err := goconfig.LoadConfigFile("./config/config.ini")
	if err != nil {
		return "", "", err
	}
	username, err := config.GetValue("proxy_server", "username")
	if err != nil {
		return "", "", err
	}
	password, err := config.GetValue("proxy_server", "password")
	if err != nil {
		return "", "", err
	}
	return username, password, nil
}
