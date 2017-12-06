package procotol

import (
	"net/url"
	"proxyWebApplication/util"
)

func GetBridgeCommand(data url.Values) (string, error) {
	path, err := util.GetServerPath()
	if err != nil {
		return "", err
	}
	command := path + "proxy bridge -p " + data["proxyIp"][0]
	command += util.HandelTls(data["crt"][0], data["key"][0])
	return command, nil
}
