package procotol

import (
	"net/url"
	"proxy/util"
)

func GetUdpCommand(data url.Values) (string, error) {
	var command string
	encryptCommand, encryptParamater, err := util.HandleEncrypt(data)
	if err != nil {
		return "", err
	}
	path, err := util.GetServerPath()
	if err != nil {
		return "", err
	}
	command = path + "proxy udp -p" + data["proxyIp"][0] + encryptCommand + encryptParamater
	return command, nil
}
