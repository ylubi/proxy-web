package procotol

import (
	"proxy-web/util"
)

func GetUdpCommand(data *util.Parameter) (string, error) {
	var command string
	encryptCommand, encryptParamater, err := util.HandleEncrypt(data)
	if err != nil {
		return "", err
	}
	path, err := util.GetServerPath()
	if err != nil {
		return "", err
	}
	command = path + "proxy udp -p" + data.ProxyIp + encryptCommand + encryptParamater
	return command, nil
}
