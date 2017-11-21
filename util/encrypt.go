package util

import (
	"fmt"
	"net/url"
)

func HandleEncrypt(data url.Values) (string, string, error) {
	//proxy：1、顶级代理 2、次级代理
	var encryptCommand string
	var encryptParamater string
	//encrypt:1、不加密 2、tls加密 3、kcp加密 4、ssh加密（密码） 5、ssh加密（密钥）
	switch data["encrypt"][0] {
	case "1":
		encryptCommand = " -T tcp -P " + data["superiorProxy"][0]
	case "2":
		encryptCommand = " -T udp -P " + data["superiorProxy"][0]
	case "3":
		encryptCommand = " -T tls -P " + data["superiorProxy"][0]
		encryptParamater = HandelTls(data["crt"][0], data["key"][0])
	case "4":
		encryptCommand = " -T kcp -P " + data["superiorProxy"][0]
		encryptParamater = HandelKcp(data["password"][0])
	case "5":
		encryptCommand = " -T ssh -P " + data["superiorProxy"][0]
		encryptParamater = HandelSshPassword(data["username"][0], data["password"][0])
	case "6":
		encryptCommand = " -T ssh -P " + data["superiorProxy"][0]
		encryptParamater = HandelSshKey(data["username"][0], data["key"][0])
	default:
		err := fmt.Errorf("%s", "parameter encrypt error")
		return "", "", err
	}
	if data["proxy"][0] == "1" {
		return "", encryptParamater, nil
	}
	return encryptCommand, encryptParamater, nil
}

func HandelTls(crt, key string) string {
	var command string
	if crt == "" {
		command += " -C ../proxyService/.cert/proxy.crt"
	} else {
		command += " -C " + crt
	}
	if key == "" {
		command += " -K ../proxyService/.cert/proxy.key"
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

func HandelSshKey(user, key string) string {
	command := " -u " + user
	if key == "" {
		command += " -S ../proxyService/.cert/proxy.key"
	} else {
		command += " -S " + key
	}
	return command
}
