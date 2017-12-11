package procotol

import (
	"encoding/json"
	"proxy-web/util"
)

func GetClientCommand(data *util.Parameter) (string, error) {
	path, err := util.GetServerPath()
	if err != nil {
		return "", err
	}
	var encrypt map[string]string
	command := path + "proxy client -P " + data.ProxyIp
	err = json.Unmarshal([]byte(data.EncryptionCondition), &encrypt)
	command += util.HandelTls(encrypt["crt"], encrypt["key"])
	return command, nil
}
