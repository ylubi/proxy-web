package procotol

import (
	"encoding/json"
	"proxy-web/util"
)

func GetBridgeCommand(data *util.Parameter) (string, error) {
	path, err := util.GetServerPath()
	if err != nil {
		return "", err
	}
	command := path + "proxy bridge -p " + data.ProxyIp
	var encrypt map[string]string
	err = json.Unmarshal([]byte(data.EncryptionCondition), &encrypt)
	command += util.HandelTls(encrypt["crt"], encrypt["key"])
	return command, nil
}
