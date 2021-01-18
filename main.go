package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"rap/dbConfig"
	"rap/service"
	"rap/utils"
)

func String(model interface{}) string {
	raw, _ := json.Marshal(model)
	return string(raw)
}

func main() {
	dbConfig.InitConnection()
	service.GenRoleInsertStatement(35, 36)
	service.GenPermissionInsertStatement(35, 36)
	service.GenAPIInsertStatement(35, 36)
	service.GenRolePermissionStatement()
	service.GenAPIPermissionStatement()
	statement := service.GetStatement()
	if utils.WriteFile {
		err := ioutil.WriteFile("insert.sql", []byte(statement), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(service.GetStatement())
}
