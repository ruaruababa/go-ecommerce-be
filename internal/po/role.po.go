package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id; type: int; primaryKey; not null; autoIncrement; "`
	RoleName string `gorm:"column:role_name; type: varchar(255);"`
	RoleNote string `gorm:"column:role_note; type: text;"`
}

func (u *Role) TableName() string {
	return "go_roles"
}
