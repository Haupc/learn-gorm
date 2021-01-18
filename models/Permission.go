package models

//   Profiles []Profile `gorm:"many2many:user_profiles;foreignKey:Refer;joinForeignKey:UserReferID;References:UserRefer;JoinReferences:UserRefer"`
// Which creates join table: user_profiles
//   foreign key: user_refer_id, reference: users.refer
//   foreign key: profile_refer, reference: profiles.user_refer

type Permission struct {
	ID     int
	Name   string `gorm:"column:NAME"`
	ZoneID int    `gorm:"column:ZONE_ID"`
	Status int    `gorm:"column:STATUS"`
	Roles  []Role `gorm:"many2many:ROLE_PERMISSION"`
	APIS   []API  `gorm:"many2many:API_PERMISSION"`
}

func (Permission) TableName() string {
	return "PERMISSION"
}
