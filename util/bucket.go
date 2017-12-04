package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

type Parameter struct {
	Protocol            string
	ProxyLevel          int
	ProxyIp             string
	SuperiorProxyIp     string
	Superior            int
	EncryptionCondition string
	ProcessId           int
	Local               string
}

func GetParameter() map[string]*Parameter {
	data := make(map[string]*Parameter)
	db, err := bolt.Open("./db/bucket/proxy.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bt, err := tx.CreateBucketIfNotExists([]byte("proxy"))
		if err != nil {
			return err
		}
		c := bt.Cursor()

		for k, v := c.Last(); k != nil; k, v = c.Prev() {
			parameter := new(Parameter)
			key := string(k[:])
			err := json.Unmarshal(v, &parameter)
			if err != nil {
				return err
			}
			data[key] = parameter
		}
		return nil
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	return data
}

func SaveParameterByPid(data url.Values, pid int) error {
	db, err := bolt.Open("./db/bucket/proxy.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	if (data["protocol"][0] == "tclient") || (data["protocol"][0] == "tserver") || (data["protocol"][0] == "tbridge") {
		data["encrypt"][0] = "3"
	}
	parameter := new(Parameter)
	parameter.EncryptionCondition, err = getEncryptionCondition(data)
	if err != nil {
		return err
	}
	parameter.Protocol = data["protocol"][0]
	parameter.ProxyLevel, err = strconv.Atoi(data["proxy"][0])
	if err != nil {
		return err
	}
	parameter.ProxyIp = data["proxyIp"][0]
	parameter.SuperiorProxyIp = data["superiorProxy"][0]
	parameter.Superior, err = strconv.Atoi(data["encrypt"][0])
	if err != nil {
		return err
	}
	parameter.ProcessId = pid
	parameter.Local = data["local"][0]
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("proxy"))
		bytePid := strconv.Itoa(pid)
		buf, err := json.Marshal(parameter)
		if err != nil {
			return err
		}
		b.Put([]byte(bytePid), buf)
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func DeleteParameterByPid(pid int) error {
	db, err := bolt.Open("./db/bucket/proxy.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	bytePid := strconv.Itoa(pid)
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("proxy"))
		err = b.Delete([]byte(bytePid))
		return err
	})
	return err
}

func getEncryptionCondition(data url.Values) (string, error) {
	jsons := make(map[string]string)
	switch data["encrypt"][0] {
	case "1":
		return "", nil
	case "2":
		return "", nil
	case "3":
		jsons["crt"] = data["crt"][0]
		jsons["key"] = data["key"][0]
		data, err := json.Marshal(jsons)
		if err != nil {
			log.Fatal(err.Error())
		}
		return string(data), nil
	case "4":
		jsons["password"] = data["password"][0]
		data, err := json.Marshal(jsons)
		if err != nil {
			log.Fatal(err.Error())
		}
		return string(data), nil
	case "5":
		jsons["username"] = data["username"][0]
		jsons["password"] = data["password"][0]
		data, err := json.Marshal(jsons)
		if err != nil {
			log.Fatal(err.Error())
		}
		return string(data), nil
	case "6":
		jsons["username"] = data["username"][0]
		jsons["key"] = data["key"][0]
		data, err := json.Marshal(jsons)
		if err != nil {
			log.Fatal(err.Error())
		}
		return string(data), nil

	}
	err := fmt.Errorf("%s", "parameter error")
	return "", err
}
