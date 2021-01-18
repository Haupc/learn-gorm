package main

import (
	"encoding/json"
	"rap/dbConfig"
	"rap/service"
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
	// fmt.Println(service.GetStatement())
	// fmt.Println(String(repository.Permissions[2]))

}
