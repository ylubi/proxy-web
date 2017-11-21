package procotol

import (
	"fmt"
	"net/url"
	"proxyWebApplication/util"
)

func GetTcpCommand(data url.Values) (string, error) {
	var command string
	encryptCommand, encryptParamater, err := util.HandleEncrypt(data)
	if err != nil {
		return "", err
	}
	path, err := util.GetServerPath()
	if err != nil {
		return "", err
	}
	//local 是否本地使用
	switch data["local"][0] {
	case "1":
		command = path + "/proxy tcp -p " + data["proxyIp"][0] + encryptCommand + encryptParamater
	case "3":
		command = path + "proxy tcp -t tls -p " + data["proxyIp"][0] + encryptCommand + encryptParamater
	case "4":
		command = path + "proxy tcp -t kcp -p " + data["proxyIp"][0] + encryptCommand + encryptParamater
	default:
		err = fmt.Errorf("%s", "paramater local error")
		return "", err
	}

	return command, nil
}
