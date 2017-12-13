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
	Id                  string
	Protocol            string
	ProxyLevel          int
	ProxyIp             string
	SuperiorProxyIp     string
	Superior            int
	EncryptionCondition string
	ProcessId           int
	Local               string
	Status              string
	Auto                string
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

func GetParameterById(id string) (*Parameter, error) {
	db, err := bolt.Open("./db/bucket/proxy.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	parameter := new(Parameter)
	err = db.View(func(tx *bolt.Tx) error {
		bt := tx.Bucket([]byte("proxy"))
		if err != nil {
			return err
		}
		data := bt.Get([]byte(id))
		err = json.Unmarshal(data, &parameter)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return parameter, err
	}
	return parameter, nil
}

func SaveParameter(data url.Values) (string, string, error) {
	db, err := bolt.Open("./db/bucket/proxy.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	if (data["protocol"][0] == "client") || (data["protocol"][0] == "server") || (data["protocol"][0] == "bridge") {
		data["superior"][0] = "3"
	}
	parameter := new(Parameter)

	parameter.EncryptionCondition, err = getEncryptionCondition(data)
	if err != nil {
		return "", "", err
	}
	parameter.Protocol = data["protocol"][0]
	parameter.ProxyLevel, err = strconv.Atoi(data["proxyLevel"][0])
	if err != nil {
		return "", "", err
	}
	parameter.ProxyIp = data["proxyIp"][0]
	parameter.SuperiorProxyIp = data["superiorProxyIp"][0]
	parameter.Superior, err = strconv.Atoi(data["superior"][0])
	if err != nil {
		return "", "", err
	}
	parameter.Status = "未开启"
	parameter.Auto = data["auto"][0]
	parameter.ProcessId = 0
	parameter.Local = data["local"][0]
	var stringId string
	var buf []byte
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("proxy"))
		if data["id"][0] == "" {
			id, _ := b.NextSequence()
			stringId = string(id)
		} else {
			stringId = data["id"][0]
		}
		parameter.Id = stringId
		buf, err = json.Marshal(parameter)
		if err != nil {
			return err
		}
		b.Put([]byte(stringId), buf)
		return nil
	})
	if err != nil {
		return "", "", err
	}
	return stringId, string(buf), nil
}

func ChangeParameterDataById(pid int, status, id string) error {
	db, err := bolt.Open("./db/bucket/proxy.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	parameter := new(Parameter)
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("proxy"))
		data := b.Get([]byte(id))
		err := json.Unmarshal(data, &parameter)
		if err != nil {
			return err
		}
		if status != "" {
			parameter.Status = status
		}
		parameter.ProcessId = pid
		data, err = json.Marshal(parameter)
		if err != nil {
			return err
		}
		b.Put([]byte(id), data)
		return nil
	})
	return err
}

func DeleteParameterDataById(id string) error {
	db, err := bolt.Open("./db/bucket/proxy.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("proxy"))
		err := b.Delete([]byte(id))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func getEncryptionCondition(data url.Values) (string, error) {
	jsons := make(map[string]string)
	switch data["superior"][0] {
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
		jsons["password"] = data["password"][0]
		data, err := json.Marshal(jsons)
		if err != nil {
			log.Fatal(err.Error())
		}
		return string(data), nil

	}
	err := fmt.Errorf("%s", "parameter error")
	return "", err
}
