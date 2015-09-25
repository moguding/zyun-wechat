package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZyUserConsume struct {
	Id           int       `orm:"column(id);auto"`
	UserId       int64     `orm:"column(user_id)"`
	ConsumeType  int8      `orm:"column(consume_type)"`
	Fee          float64   `orm:"column(fee);digits(10);decimals(2)"`
	Amount       float64   `orm:"column(amount);digits(10);decimals(2)"`
	CurrencyRate float64   `orm:"column(currency_rate);digits(10);decimals(4)"`
	Source       string    `orm:"column(source);size(100);null"`
	Remark       string    `orm:"column(remark);size(200);null"`
	UpdatedAt    time.Time `orm:"column(updated_at);type(timestamp);null"`
	CreatedAt    time.Time `orm:"column(created_at);type(datetime);auto_now_add"`
}

func (t *ZyUserConsume) TableName() string {
	return "zy_user_consume"
}

func init() {
	orm.RegisterModel(new(ZyUserConsume))
}

// AddZyUserConsume insert a new ZyUserConsume into database and returns
// last inserted Id on success.
func AddZyUserConsume(m *ZyUserConsume) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZyUserConsumeById retrieves ZyUserConsume by Id. Returns error if
// Id doesn't exist
func GetZyUserConsumeById(id int) (v *ZyUserConsume, err error) {
	o := orm.NewOrm()
	v = &ZyUserConsume{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZyUserConsume retrieves all ZyUserConsume matches certain condition. Returns empty list if
// no records exist
func GetAllZyUserConsume(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZyUserConsume))
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

	var l []ZyUserConsume
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

// UpdateZyUserConsume updates ZyUserConsume by Id and returns error if
// the record to be updated doesn't exist
func UpdateZyUserConsumeById(m *ZyUserConsume) (err error) {
	o := orm.NewOrm()
	v := ZyUserConsume{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZyUserConsume deletes ZyUserConsume by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZyUserConsume(id int) (err error) {
	o := orm.NewOrm()
	v := ZyUserConsume{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZyUserConsume{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
