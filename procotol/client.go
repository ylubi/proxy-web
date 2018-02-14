package procotol

import (
	"encoding/json"
	"proxy-web/util"
)

type Client struct {
}

func NewClient() Protocol {
	return &Client{}
}

func (c *Client) GetCommand(data *util.Parameter) (string, error) {
	path, err := util.NewConfig().GetServerPath()
	compress := util.CompressCommand(data.Compress)
	if err != nil {
		return "", err
	}
	var encrypt map[string]string
	command := path + "proxy client -P " + data.ProxyIp + compress
	err = json.Unmarshal([]byte(data.EncryptionCondition), &encrypt)
	command += util.HandelTls(encrypt["crt"], encrypt["key"])
	return command, nil
}
