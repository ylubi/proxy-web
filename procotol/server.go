package procotol

import (
	"encoding/json"
	"proxy-web/util"
)

type Server struct {
}

func NewServer() Protocol {
	return &Server{}
}

func (s *Server) GetCommand(data *util.Parameter) (string, error) {
	path, err := util.NewConfig().GetServerPath()
	compress := util.CompressCommand(data.Compress)
	if err != nil {
		return "", err
	}
	command := path + "proxy server -r " + data.ProxyIp + " -P " + data.SuperiorProxyIp + compress
	var encrypt map[string]string
	err = json.Unmarshal([]byte(data.EncryptionCondition), &encrypt)
	if err != nil {
		return "", err
	}
	command += util.HandelTls(encrypt["crt"], encrypt["key"])
	return command, nil
}
