package repository

import (
	"log"
	"rap/dbConfig"
	"rap/models"
)

// Permissions : singleton
var Permissions []models.Permission

// GetPermissons find all permission by zone id
func GetPermissons(zoneID int) []models.Permission {
	if Permissions == nil {
		if err := dbConfig.DB.Preload("Roles").Preload("APIS").Where("zone_id = ?", zoneID).Find(&Permissions).Error; err != nil {
			log.Fatal(err)
		}
	}
	return Permissions
}
