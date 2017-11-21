package util

import (
	"encoding/json"
)

type Result struct {
	Code   int
	Output interface{}
	Pid    int
}

func (r *Result) ReturnJson() (string, error) {
	res, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	//停止服务
	return string(res), nil
}
