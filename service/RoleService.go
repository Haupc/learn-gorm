package service

import (
	"fmt"
	"rap/repository"
)

var baseRoleStatement = "insert into auth.SE_ROLE (`CODE`, `NAME`, zone_id, `STATUS`, PARENT) values ('%s', '%s', %d, %d, %s);\n"

// GenRoleInsertStatement 2 params
// zoneID : current zone id
// newZoneID : new zone id
func GenRoleInsertStatement(zoneID int, newZoneID int) {
	statement += "\n"
	roles := repository.GetRoles(zoneID)
	for _, role := range roles {
		var s string
		if s = "NULL"; role.Parent.Valid {
			s = fmt.Sprintf("'%d'", role.Parent.Int32)
		}
		statement += fmt.Sprintf(baseRoleStatement, role.Code, role.Name, newZoneID, role.Status, s)
	}
}
