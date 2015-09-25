package models

type MgtUserRole struct {
	UserId *MgtUser `orm:"column(user_id);rel(fk)"`
	RoleId *MgtRole `orm:"column(role_id);rel(fk)"`
}
