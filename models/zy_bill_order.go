package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZyBillOrder struct {
	Id            int       `orm:"column(id);auto"`
	UserId        int64     `orm:"column(user_id)"`
	OrderNo       string    `orm:"column(order_no);size(50)"`
	OrderType     int8      `orm:"column(order_type)"`
	OrderAmount   float64   `orm:"column(order_amount);digits(10);decimals(2)"`
	Coupon        string    `orm:"column(coupon);size(100);null"`
	ChannelType   int8      `orm:"column(channel_type)"`
	ChannelNote   string    `orm:"column(channel_note);size(500);null"`
	RelatePackage string    `orm:"column(relate_package);size(100);null"`
	OrderStatus   int8      `orm:"column(order_status)"`
	Remark        string    `orm:"column(remark);size(250);null"`
	Version       int       `orm:"column(version)"`
	UpdatedAt     time.Time `orm:"column(updated_at);type(datetime);null"`
	CreatedAt     time.Time `orm:"column(created_at);type(datetime);auto_now_add"`
}

func (t *ZyBillOrder) TableName() string {
	return "zy_bill_order"
}

func init() {
	orm.RegisterModel(new(ZyBillOrder))
}

// AddZyBillOrder insert a new ZyBillOrder into database and returns
// last inserted Id on success.
func AddZyBillOrder(m *ZyBillOrder) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZyBillOrderById retrieves ZyBillOrder by Id. Returns error if
// Id doesn't exist
func GetZyBillOrderById(id int) (v *ZyBillOrder, err error) {
	o := orm.NewOrm()
	v = &ZyBillOrder{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZyBillOrder retrieves all ZyBillOrder matches certain condition. Returns empty list if
// no records exist
func GetAllZyBillOrder(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZyBillOrder))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ZyBillOrder
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateZyBillOrder updates ZyBillOrder by Id and returns error if
// the record to be updated doesn't exist
func UpdateZyBillOrderById(m *ZyBillOrder) (err error) {
	o := orm.NewOrm()
	v := ZyBillOrder{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZyBillOrder deletes ZyBillOrder by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZyBillOrder(id int) (err error) {
	o := orm.NewOrm()
	v := ZyBillOrder{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZyBillOrder{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
