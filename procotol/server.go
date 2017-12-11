package procotol

import (
	"encoding/json"
	"proxy-web/util"
)

func GetServerCommand(data *util.Parameter) (string, error) {
	path, err := util.GetServerPath()
	if err != nil {
		return "", err
	}
	command := path + "proxy server -r " + data.ProxyIp + " -P " + data.SuperiorProxyIp
	var encrypt map[string]string
	err = json.Unmarshal([]byte(data.EncryptionCondition), &encrypt)
	if err != nil {
		return "", err
	}
	command += util.HandelTls(encrypt["crt"], encrypt["key"])
	return command, nil
}
