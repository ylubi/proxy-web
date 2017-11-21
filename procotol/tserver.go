package procotol

import (
	"net/url"
	"proxy/util"
)

func GetTserverCommand(data url.Values) (string, error) {
	path, err := util.GetServerPath()
	if err != nil {
		return "", err
	}
	command := path + "proxy tserver -r " + data["proxyIp"][0] + " -P " + data["superiorProxy"][0]
	command += util.HandelTls(data["crt"][0], data["key"][0])
	return command, nil
}
