package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZyUserWallet struct {
	Id              int       `orm:"column(user_id);pk"`
	UserName        string    `orm:"column(user_name);size(50)"`
	Balance         float64   `orm:"column(balance);digits(10);decimals(2)"`
	TotalAmount     float64   `orm:"column(total_amount);digits(10);decimals(2)"`
	AmountOnThatDay float64   `orm:"column(amount_on_that_day);digits(10);decimals(2)"`
	PayPasswd       string    `orm:"column(pay_passwd);size(50);null"`
	PaySalt         string    `orm:"column(pay_salt);size(20);null"`
	Status          int8      `orm:"column(status);null"`
	LastPutDate     time.Time `orm:"column(last_put_date);type(datetime);null"`
	LastConsumeDate time.Time `orm:"column(last_consume_date);type(datetime);null"`
	FrozenDate      time.Time `orm:"column(frozen_date);type(datetime);null"`
	Version         int       `orm:"column(version);null"`
	UpdatedAt       time.Time `orm:"column(updated_at);type(timestamp);null;auto_now"`
	CreatedAt       time.Time `orm:"column(created_at);type(datetime);null"`
}

func (t *ZyUserWallet) TableName() string {
	return "zy_user_wallet"
}

func init() {
	orm.RegisterModel(new(ZyUserWallet))
}

// AddZyUserWallet insert a new ZyUserWallet into database and returns
// last inserted Id on success.
func AddZyUserWallet(m *ZyUserWallet) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZyUserWalletById retrieves ZyUserWallet by Id. Returns error if
// Id doesn't exist
func GetZyUserWalletById(id int) (v *ZyUserWallet, err error) {
	o := orm.NewOrm()
	v = &ZyUserWallet{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZyUserWallet retrieves all ZyUserWallet matches certain condition. Returns empty list if
// no records exist
func GetAllZyUserWallet(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZyUserWallet))
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

	var l []ZyUserWallet
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

// UpdateZyUserWallet updates ZyUserWallet by Id and returns error if
// the record to be updated doesn't exist
func UpdateZyUserWalletById(m *ZyUserWallet) (err error) {
	o := orm.NewOrm()
	v := ZyUserWallet{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZyUserWallet deletes ZyUserWallet by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZyUserWallet(id int) (err error) {
	o := orm.NewOrm()
	v := ZyUserWallet{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZyUserWallet{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
