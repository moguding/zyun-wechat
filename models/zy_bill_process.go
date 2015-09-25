package models

import "time"

type ZyBillProcess struct {
	OrderNo    string    `orm:"column(order_no);size(50)"`
	IsComplete int8      `orm:"column(is_complete)"`
	UpdatedAt  time.Time `orm:"column(updated_at);type(datetime);null"`
	CreatedAt  time.Time `orm:"column(created_at);type(datetime);auto_now_add"`
}
