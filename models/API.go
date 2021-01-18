package models

type API struct {
	ID          int
	API         string       `gorm:"column:API"`
	Method      string       `gorm:"column:METHOD"`
	Status      int          `gorm:"column:STATUS"`
	ZoneID      int          `gorm:"column:ZONE_ID"`
	Permissions []Permission `gorm:"many2many:API_PERMISSION"`
}

func (API) TableName() string {
	return "API"
}
