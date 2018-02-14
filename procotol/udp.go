package procotol

import (
	"proxy-web/util"
)

type Udp struct {
}

func NewUdp() Protocol {
	return &Udp{}
}

func (u *Udp) GetCommand(data *util.Parameter) (string, error) {
	var command string
	encryptCommand, encryptParamater, err := util.HandleEncrypt(data)
	if err != nil {
		return "", err
	}
	path, err := util.NewConfig().GetServerPath()
	if err != nil {
		return "", err
	}
	command = path + "proxy udp -p" + data.ProxyIp + encryptCommand + encryptParamater
	return command, nil
}
