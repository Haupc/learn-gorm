package service

import (
	"fmt"
	"rap/repository"
)

var basePermissionStatement = "insert into PERMISSION (`NAME` , zone_id, `status`) values ('%s', %d, %d);\n"

// GenPermissionInsertStatement 2 params
// zoneID : current zone id
// newZoneID : new zone id
func GenPermissionInsertStatement(zoneID int, newZoneID int) {
	statement += "\n"
	permissions := repository.GetPermissons(zoneID)
	for _, permission := range permissions {
		statement += fmt.Sprintf(basePermissionStatement, permission.Name, newZoneID, permission.Status)
	}
}
