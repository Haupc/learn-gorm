package utils

import (
	"gorm.io/gorm/schema"
)

func GetNamingStrategy() NamingStrategy {
	namingStraegy := NamingStrategy{}
	namingStraegy.SingularTable = true
	return namingStraegy
}

type NamingStrategy struct {
	schema.NamingStrategy
}

// JoinTableName convert string to join table name
func (ns NamingStrategy) JoinTableName(str string) string {
	return str
}
