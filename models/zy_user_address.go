package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZyUserAddress struct {
	Id         int       `orm:"column(id);auto"`
	UserId     int64     `orm:"column(user_id)"`
	FullName   string    `orm:"column(full_name);size(50)"`
	IdentityId int64     `orm:"column(identity_id);null"`
	Province   string    `orm:"column(province);size(50);null"`
	City       string    `orm:"column(city);size(50);null"`
	Region     string    `orm:"column(region);size(50);null"`
	Address    string    `orm:"column(address);size(200);null"`
	Zip        string    `orm:"column(zip);size(20);null"`
	PhoneNum   string    `orm:"column(phone_num);size(50);null"`
	IsDefault  int8      `orm:"column(is_default);null"`
	IsEnabeld  int8      `orm:"column(is_enabeld);null"`
	UpdatedAt  time.Time `orm:"column(updated_at);type(datetime);null"`
	CreatedAt  time.Time `orm:"column(created_at);type(datetime);auto_now_add"`
}

func (t *ZyUserAddress) TableName() string {
	return "zy_user_address"
}

func init() {
	orm.RegisterModel(new(ZyUserAddress))
}

// AddZyUserAddress insert a new ZyUserAddress into database and returns
// last inserted Id on success.
func AddZyUserAddress(m *ZyUserAddress) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZyUserAddressById retrieves ZyUserAddress by Id. Returns error if
// Id doesn't exist
func GetZyUserAddressById(id int) (v *ZyUserAddress, err error) {
	o := orm.NewOrm()
	v = &ZyUserAddress{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZyUserAddress retrieves all ZyUserAddress matches certain condition. Returns empty list if
// no records exist
func GetAllZyUserAddress(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZyUserAddress))
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

	var l []ZyUserAddress
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

// UpdateZyUserAddress updates ZyUserAddress by Id and returns error if
// the record to be updated doesn't exist
func UpdateZyUserAddressById(m *ZyUserAddress) (err error) {
	o := orm.NewOrm()
	v := ZyUserAddress{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZyUserAddress deletes ZyUserAddress by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZyUserAddress(id int) (err error) {
	o := orm.NewOrm()
	v := ZyUserAddress{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZyUserAddress{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
