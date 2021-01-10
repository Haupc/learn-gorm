package models

//   Profiles []Profile `gorm:"many2many:user_profiles;foreignKey:Refer;joinForeignKey:UserReferID;References:UserRefer;JoinReferences:UserRefer"`
// Which creates join table: user_profiles
//   foreign key: user_refer_id, reference: users.refer
//   foreign key: profile_refer, reference: profiles.user_refer

type Permission struct {
	ID     int
	Name   string  `gorm:"column:NAME"`
	ZoneID int     `gorm:"column:ZONE_ID"`
	Status int     `gorm:"column:STATUS"`
	Roles  []*Role `gorm:"many2many:ROLE_PERMISSION;jointable_foreignkey:permission_id;associaton_jointable_foreignkey:role_id" json:"roles,omitempty"`
	APIS   []*API  `gorm:"many2many:API_PERMISSION;jointable_foreignkey:permission_id;associaton_jointable_foreignkey:api_id"`
}

func (Permission) TableName() string {
	return "PERMISSION"
}
