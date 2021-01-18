package service

import (
	"fmt"
	"rap/repository"
)

var baseAPIStatement = "insert into API (API , `method`, zone_id, `status`) values ('%s', '%s', %d, %d);\n"

// GenAPIInsertStatement 2 params
// zoneID : current zone id
// newZoneID : new zone id
func GenAPIInsertStatement(zoneID int, newZoneID int) {
	statement += "\n"
	apis := repository.GetAPIS(zoneID)
	for _, api := range apis {
		statement += fmt.Sprintf(baseAPIStatement, api.API, api.Method, newZoneID, api.Status)
	}
}
