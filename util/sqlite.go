package util

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Parameter struct {
	Id                  int
	Protocol            string
	ProxyLevel          int
	ProxyIp             string
	SuperiorProxyIp     string
	Superior            int
	EncryptionCondition string
	ProcessId           int
	Local               string
}

func GetParameterExistPid() map[string]*Parameter {
	data := make(map[string]*Parameter)
	db, err := sql.Open("sqlite3", "./db/sqlite/foo.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	rows, err := db.Query("SELECT * FROM parameter where process_id != 0")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		p := new(Parameter)
		rows.Scan(&p.Id, &p.Protocol, &p.ProxyLevel, &p.ProxyIp, &p.SuperiorProxyIp, &p.Superior, &p.EncryptionCondition, &p.ProcessId, &p.Local)
		id := strconv.Itoa(p.Id)
		data[id] = p
	}
	return data
}

func SaveParameterByPid(data url.Values, pid int) int64 {
	db, err := sql.Open("sqlite3", "./db/sqlite/foo.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	if (data["protocol"][0] == "tclient") || (data["protocol"][0] == "tserver") || (data["protocol"][0] == "tbridge") {
		data["encrypt"][0] = "3"
	}
	encryption_condition, err := getEncryptionCondition(data)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt, err := db.Prepare("INSERT INTO parameter(protocol,proxy_level,proxy_ip,superior_proxy_ip,superior,encryption_condition,process_id, local) values(?,?,?,?,?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	result, err := stmt.Exec(data["protocol"][0], data["proxy"][0], data["proxyIp"][0], data["superiorProxy"][0], data["encrypt"][0], encryption_condition, pid, data["local"][0])
	if err != nil {
		log.Fatal(err.Error())
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
	}
	return lastId
}

func PutParameterPidTo0(pid int) int64 {
	db, err := sql.Open("sqlite3", "./db/sqlite/foo.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt, err := db.Prepare("update parameter set process_id = 0 where process_id = ?")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	result, err := stmt.Exec(pid)
	if err != nil {
		log.Fatal(err.Error())
	}
	affect, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err.Error())
	}
	return affect
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
