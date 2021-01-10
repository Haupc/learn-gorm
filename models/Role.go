package models

type Role struct {
	ID          int
	Name        string        `gorm:"column:NAME"`
	ZoneID      int           `gorm:"column:ZONE_ID"`
	Code        string        `gorm:"column:CODE"`
	Status      int           `gorm:"column:STATUS"`
	Permissions []*Permission `gorm:"many2many:ROLE_PERMISSION;jointable_foreignkey:role_id;associaton_jointable_foreignkey:permission_id" json:"permissions,omitempty"`
}

func (Role) TableName() string {
	return "SE_ROLE"
}
