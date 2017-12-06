package procotol

import (
	"net/url"
	"proxyWebApplication/util"
)

func GetServerCommand(data url.Values) (string, error) {
	path, err := util.GetServerPath()
	if err != nil {
		return "", err
	}
	command := path + "proxy server -r " + data["proxyIp"][0] + " -P " + data["superiorProxy"][0]
	command += util.HandelTls(data["crt"][0], data["key"][0])
	return command, nil
}
