package util

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
	"net/url"
	"time"

	"github.com/boltdb/bolt"
)

type Parameter struct {
	Id        string
	Name      string // 名称
	Params    string // 参数
	ProcessId int    // 进程 id
	Status    string // 是否开启
	Auto      string // 是否自动开启
	Key       string // .key 文件路径
	Crt       string // .crt 文件路径
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
	parameter := new(Parameter)

	parameter.Name = data.Get("name")
	parameter.Params = data.Get("param")
	parameter.Crt = data.Get("crt")
	parameter.Key = data.Get("key")
	if err != nil {
		return "", "", err
	}
	parameter.Status = "未开启"
	parameter.Auto = data.Get("auto")
	parameter.ProcessId = 0

	var stringId string
	var buf []byte
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("proxy"))
		if data["id"][0] == "" {
			id, _ := b.NextSequence()
			stringId = string(id)
		} else {
			stringId = data.Get("id")
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

func SaveSession(sessionId string) error {
	db, err := bolt.Open("./db/bucket/proxy.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("session"))
		b.Put([]byte("sessionId"), []byte(sessionId))
		t := time.Now()
		timeStamp := t.Unix() + 3600
		b.Put([]byte("time"), IntToBytes(int(timeStamp)))
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func GetSession() (string, int, error) {
	db, err := bolt.Open("./db/bucket/proxy.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	var sessionId string
	var timeStamp int
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("session"))
		sessionId = string(b.Get([]byte("sessionId")))
		timeStamp = BytesToInt(b.Get([]byte("time")))
		return nil
	})
	return sessionId, timeStamp, nil
}

func InitSession() {
	db, err := bolt.Open("./db/bucket/proxy.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("session"))
		if err != nil {
			return err
		}
		b.Put([]byte("sessionId"), []byte(""))
		b.Put([]byte("time"), IntToBytes(0))
		return nil
	})
	if err != nil {
		log.Fatal(err.Error())
	}
}

func IntToBytes(n int) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int(tmp)
}
