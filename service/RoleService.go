package service

import (
	"fmt"
	"rap/repository"
)

var baseRoleStatement = "insert into SE_ROLE (`CODE`, `NAME`, zone_id, `STATUS`) values ('%s', '%s', %d, %d);\n"

// GenRoleInsertStatement 2 params
// zoneID : current zone id
// newZoneID : new zone id
func GenRoleInsertStatement(zoneID int, newZoneID int) {
	statement += "\n"
	roles := repository.GetRoles(zoneID)
	for _, role := range roles {
		statement += fmt.Sprintf(baseRoleStatement, role.Code, role.Name, newZoneID, role.Status)
	}
}
