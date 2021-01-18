package models

import (
	"database/sql"
)

type Role struct {
	ID          int
	Name        string        `gorm:"column:NAME"`
	ZoneID      int           `gorm:"column:ZONE_ID"`
	Code        string        `gorm:"column:CODE"`
	Status      int           `gorm:"column:STATUS"`
	Permissions []Permission  `gorm:"many2many:ROLE_PERMISSION"`
	Parent      sql.NullInt32 `gorm:"column:PARENT"`
}

func (Role) TableName() string {
	return "SE_ROLE"
}
