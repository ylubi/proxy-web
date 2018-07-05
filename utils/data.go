package utils

import (
	"time"
	"os"
	"encoding/json"
	"io/ioutil"
)

const (
	dataFilePath = "./data/services/"
)

func SaveParams(name, command, auto_start, key_file, crt_file, log string) (serviceIdStr string, err error) {
	serviceId := time.Now().UnixNano() / 1000000
	serviceIdStr = NewConvert().IntToString(serviceId, 10)
	filePath, err := NewConfig().GetServicesFilePath()
	if err != nil {
		return
	}
	fd, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(fd)
	if err != nil {
		return
	}
	fd.Close()
	dataMap := make(map[string]interface{})
	json.Unmarshal(data, &dataMap)
	dataMap[serviceIdStr] = auto_start
	data, _ = json.Marshal(dataMap)
	ioutil.WriteFile(filePath, data, 0644)

	// 判断有没有services文件夹
	s, err := os.Stat(dataFilePath)
	if err != nil || !s.IsDir() {
		os.Mkdir(dataFilePath, os.ModePerm)
	}

	fd, err = os.OpenFile(dataFilePath+serviceIdStr+".json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	params := make(map[string]string)
	params["name"] = name
	params["command"] = command
	params["auto_start"] = auto_start
	params["key_file"] = key_file
	params["crt_file"] = crt_file
	params["id"] = serviceIdStr
	params["status"] = "未开启"
	params["log"] = log
	paramJson, _ := json.Marshal(params)
	fd.Write(paramJson)
	fd.Close()
	return
}

func UpdateParams(serviceId, name, command, auto_start, key_file, crt_file, log string) (err error) {
	filePath, err := NewConfig().GetServicesFilePath()
	if err != nil {
		return
	}
	fd, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(fd)
	if err != nil {
		return
	}
	fd.Close()
	dataMap := make(map[string]interface{})
	json.Unmarshal(data, &dataMap)
	dataMap[serviceId] = auto_start
	data, _ = json.Marshal(dataMap)
	ioutil.WriteFile(filePath, data, 0644)

	// 判断有没有services文件夹
	s, err := os.Stat(dataFilePath)
	if err != nil || !s.IsDir() {
		os.Mkdir(dataFilePath, os.ModePerm)
	}

	fd, err = os.OpenFile(dataFilePath+serviceId+".json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	params := make(map[string]string)
	params["name"] = name
	params["command"] = command
	params["auto_start"] = auto_start
	params["key_file"] = key_file
	params["crt_file"] = crt_file
	params["id"] = serviceId
	params["status"] = "未开启"
	params["log"] = log
	paramJson, _ := json.Marshal(params)
	fd.Write(paramJson)
	fd.Close()
	return
}

func DeleteParam(serviceId string) (err error) {
	filePath, err := NewConfig().GetServicesFilePath()
	if err != nil {
		return
	}
	fd, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer fd.Close()

	allData, err := ioutil.ReadAll(fd)
	if err != nil {
		return
	}
	dataMap := make(map[string]interface{})
	err = json.Unmarshal(allData, &dataMap)
	if err != nil {
		return
	}
	delete(dataMap, serviceId)
	dataByte, _ := json.Marshal(dataMap)
	ioutil.WriteFile(filePath, dataByte, 0644)
	os.Remove(dataFilePath + serviceId + ".json")
	return
}

func InitParams() (datas []map[string]interface{}, err error) {
	filePath, err := NewConfig().GetServicesFilePath()
	fd, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer fd.Close()

	allData, err := ioutil.ReadAll(fd)
	if err != nil {
		return
	}
	dataMap := make(map[string]interface{})
	err = json.Unmarshal(allData, &dataMap)
	if err != nil {
		return
	}

	for serviceId, auto_start := range dataMap {
		data := make(map[string]interface{})
		fd1, err := os.OpenFile(dataFilePath+serviceId+".json", os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			continue
		}
		dataByte, err := ioutil.ReadAll(fd1)
		if err != nil {
			continue
		}
		json.Unmarshal(dataByte, &data)

		if auto_start == "是" {
			data["status"] = "已开启"
			datas = append(datas, data)
		} else {
			data["status"] = "未开启"
		}

		dataByte, _ = json.Marshal(data)
		ioutil.WriteFile(dataFilePath+serviceId+".json", dataByte, 0644)
		fd1.Close()
	}

	return
}

func GetAllParams() (datas []map[string]interface{}, err error) {
	filePath, err := NewConfig().GetServicesFilePath()
	fd, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer fd.Close()

	allData, err := ioutil.ReadAll(fd)
	if err != nil || len(allData) == 0 {
		return
	}
	dataMap := make(map[string]interface{})
	err = json.Unmarshal(allData, &dataMap)
	if err != nil {
		return
	}

	for serviceId, _ := range dataMap {
		data := make(map[string]interface{})
		fd, err := os.Open(dataFilePath + serviceId + ".json")
		if err != nil {
			continue
		}
		dataByte, err := ioutil.ReadAll(fd)
		if err != nil {
			continue
		}
		json.Unmarshal(dataByte, &data)
		datas = append(datas, data)
		fd.Close()
	}

	return
}

func GetParamsById(id string) (data map[string]interface{}, err error) {
	fd, err := os.OpenFile(dataFilePath+id+".json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer fd.Close()

	allData, err := ioutil.ReadAll(fd)
	if err != nil {
		return
	}
	err = json.Unmarshal(allData, &data)

	return
}

func ChangeParameterDataById(serviceId, status string) (err error) {
	fd, err := os.OpenFile(dataFilePath+serviceId+".json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer fd.Close()
	data, err := ioutil.ReadAll(fd)
	if err != nil {
		return
	}
	params := make(map[string]string)
	err = json.Unmarshal(data, &params)
	if err != nil {
		return
	}
	params["status"] = status
	paramJson, _ := json.Marshal(params)
	ioutil.WriteFile(dataFilePath+serviceId+".json", paramJson, 0644)
	return
}

func UpdateProxy(ip, port string) (err error) {
	data := make(map[string]interface{})
	data["ip"] = ip
	data["port"] = port
	dataByte, err := json.Marshal(data)
	if err != nil {
		return
	}

	err = ioutil.WriteFile("./data/proxy.json", dataByte, 0644)
	return
}

func GetProxy()(data map[string]string, err error){
	dataByte, err := ioutil.ReadFile("./data/proxy.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(dataByte, &data)
	return
}
