package models

type Permission struct {
	ID     int
	Name   string  `gorm:"column:NAME"`
	ZoneID int     `gorm:"column:ZONE_ID"`
	Status int     `gorm:"column:STATUS"`
	Role   []*Role `gorm:"many2many:ROLE_PERMISSION;"`
	API    []*API  `gorm:"many2many:API_PERMISSION;"`
}

func (Permission) TableName() string {
	return "PERMISSION"
}
