package procotol

import (
	"encoding/json"
	"proxy-web/util"
)

type Bridge struct {
}

func NewBridge() Protocol {
	return &Bridge{}
}

func (b *Bridge) GetCommand(data *util.Parameter) (string, error) {
	path, err := util.NewConfig().GetServerPath()
	if err != nil {
		return "", err
	}
	command := path + "proxy bridge -p " + data.ProxyIp
	var encrypt map[string]string
	err = json.Unmarshal([]byte(data.EncryptionCondition), &encrypt)
	command += util.HandelTls(encrypt["crt"], encrypt["key"])
	return command, nil
}
