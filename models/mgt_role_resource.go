package models

type MgtRoleResource struct {
	RoleId *MgtRole     `orm:"column(role_id);rel(fk)"`
	ResId  *MgtResource `orm:"column(res_id);rel(fk)"`
}
