package models

type API struct {
	ID          int
	API         string        `gorm:"column:API"`
	Method      string        `gorm:"column:METHOD"`
	Status      int           `gorm:"column:STATUS"`
	ZoneID      int           `gorm:"column:ZONE_ID"`
	Permissions []*Permission `gorm:"many2many:API_PERMISSION;jointable_foreignkey:api_id;associaton_jointable_foreignkey:permission_id"`
}

func (API) TableName() string {
	return "API"
}
