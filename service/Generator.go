package service

import (
	"fmt"
	"rap/repository"
	"rap/utils"
	"strings"
)

// Statement : generated statment
var statement = ""
var rpStatement = "insert into ROLE_PERMISSION rp (ROLE_ID, PERMISSION_ID) values (select r.ID from SE_ROLE r where r.`CODE` = '%s'), (select p.ID from PERMISSION p where p.`NAME` = '%s');\n"
var apStatement = "insert into API_PERMISSION ap (API_ID, PERMISSION_ID) values (select a.ID from API where (a.API = '%s' and a.METHOD = '%s'),  (select p.ID from PERMISSION p where p.`NAME` = '%s');\n"

// GetStatement return statement
func GetStatement() string {
	if utils.InsertIgnore {
		return strings.ReplaceAll(statement, "insert", "insert ignore")
	}
	return statement
}

func GenRolePermissionStatement() {
	statement += "\n"
	for _, p := range repository.Permissions {
		for _, r := range p.Roles {
			statement += fmt.Sprintf(rpStatement, r.Code, p.Name)
		}
	}
}

func GenAPIPermissionStatement() {
	statement += "\n"
	for _, p := range repository.Permissions {
		for _, a := range p.APIS {
			statement += fmt.Sprintf(apStatement, a.API, a.Method, p.Name)
		}
	}
}
