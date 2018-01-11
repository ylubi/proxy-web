package util

import (
	"encoding/json"
	"fmt"
)

func HandleEncrypt(data *Parameter) (string, string, error) {
	//proxy：1、顶级代理 2、次级代理
	var encryptCommand string
	var encryptParamater string
	//encrypt:1、不加密 2、tls加密 3、kcp加密 4、ssh加密（密码） 5、ssh加密（密钥）
	switch data.Superior {
	case 1:
		encryptCommand = " -T tcp -P " + data.SuperiorProxyIp
	case 2:
		encryptCommand = " -T udp -P " + data.SuperiorProxyIp
	case 3:
		encryptCommand = " -T tls -P " + data.SuperiorProxyIp
		encrypt, err := getEncrypt(data.EncryptionCondition)
		if err != nil {
			return "", "", err
		}
		encryptParamater = HandelTls(encrypt["crt"], encrypt["key"])
	case 4:
		encryptCommand = " -T kcp -P " + data.SuperiorProxyIp
		encrypt, err := getEncrypt(data.EncryptionCondition)
		if err != nil {
			return "", "", err
		}
		encryptParamater = HandelKcp(encrypt["password"])
	case 5:
		encryptCommand = " -T ssh -P " + data.SuperiorProxyIp
		encrypt, err := getEncrypt(data.EncryptionCondition)
		if err != nil {
			return "", "", err
		}
		encryptParamater = HandelSshPassword(encrypt["username"], encrypt["password"])
	case 6:
		encryptCommand = " -T ssh -P " + data.SuperiorProxyIp
		encrypt, err := getEncrypt(data.EncryptionCondition)
		if err != nil {
			return "", "", err
		}
		encryptParamater = HandelSshKey(encrypt["username"], encrypt["key"], encrypt["password"])
	default:
		err := fmt.Errorf("%s", "parameter encrypt error")
		return "", "", err
	}

	if data.ProxyLevel == 1 {
		return "", encryptParamater, nil
	}
	return encryptCommand, encryptParamater, nil
}

func getEncrypt(encryptionCondition string) (map[string]string, error) {
	var encrypt map[string]string
	err := json.Unmarshal([]byte(encryptionCondition), &encrypt)
	if err != nil {
		return encrypt, err
	}
	return encrypt, nil
}

func HandelTls(crt, key string) string {
	var command string
	path, err := GetServerPath()
	if err != nil {
		return ""
	}
	if crt == "" {
		command += " -C " + path + ".cert/proxy.crt"
	} else {
		command += " -C " + crt
	}
	if key == "" {
		command += " -K " + path + ".cert/proxy.key"
	} else {
		command += " -K " + key
	}
	return command
}

func HandelKcp(password string) string {
	command := " -B " + password
	return command
}

func HandelSshPassword(user, password string) string {
	command := " -u " + user + " -A " + password
	return command
}

func HandelSshKey(user, key, password string) string {
	command := " -u " + user
	path, err := GetServerPath()
	if err != nil {
		return ""
	}
	if key == "" {
		command += " -S " + path + ".cert/proxy.key"
	} else {
		command += " -S " + key
	}
	if password != "" {
		command += " -s " + password
	}
	return command
}
