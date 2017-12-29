package procotol

import (
	"fmt"
	"proxy-web/util"
)

func GetSocksCommand(data *util.Parameter) (string, error) {
	var command string
	encryptCommand, encryptParamater, err := util.HandleEncrypt(data)
	alwaysCommand := util.AlwaysCommand(data.Always, data.ProxyLevel)
	if err != nil {
		return "", err
	}
	path, err := util.GetServerPath()
	if err != nil {
		return "", err
	}
	//local 是否本地使用
	switch data.Local {
	case "1":
		command = path + "proxy socks -t tcp -p " + data.ProxyIp + encryptCommand + encryptParamater + alwaysCommand
	case "3":
		command = path + "proxy socks -t tls -p " + data.ProxyIp + encryptCommand + encryptParamater + alwaysCommand
	case "4":
		command = path + "proxy socks -t kcp -p " + data.ProxyIp + encryptCommand + encryptParamater + alwaysCommand
	default:
		err = fmt.Errorf("%s", "paramater local error")
		return "", err
	}
	return command, nil
}
