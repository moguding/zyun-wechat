package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZyAddressTemplate struct {
	Id           int       `orm:"column(store_id);pk"`
	StoreName    string    `orm:"column(store_name);size(100)"`
	StoreIcon    string    `orm:"column(store_icon);size(100);null"`
	FirstAddress string    `orm:"column(first_address);size(100)"`
	LastAddress  string    `orm:"column(last_address);size(100);null"`
	Recipient    string    `orm:"column(recipient);size(100);null"`
	City         string    `orm:"column(city);size(50)"`
	Region       string    `orm:"column(region);size(50)"`
	Country      string    `orm:"column(country);size(50)"`
	Zip          string    `orm:"column(zip);size(20)"`
	PhoneNum     string    `orm:"column(phone_num);size(20)"`
	IsEnabled    int8      `orm:"column(is_enabled)"`
	UpdatedAt    time.Time `orm:"column(updated_at);type(datetime);null"`
	CreatedAt    time.Time `orm:"column(created_at);type(datetime);auto_now_add"`
}

func (t *ZyAddressTemplate) TableName() string {
	return "zy_address_template"
}

func init() {
	orm.RegisterModel(new(ZyAddressTemplate))
}

// AddZyAddressTemplate insert a new ZyAddressTemplate into database and returns
// last inserted Id on success.
func AddZyAddressTemplate(m *ZyAddressTemplate) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZyAddressTemplateById retrieves ZyAddressTemplate by Id. Returns error if
// Id doesn't exist
func GetZyAddressTemplateById(id int) (v *ZyAddressTemplate, err error) {
	o := orm.NewOrm()
	v = &ZyAddressTemplate{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZyAddressTemplate retrieves all ZyAddressTemplate matches certain condition. Returns empty list if
// no records exist
func GetAllZyAddressTemplate(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZyAddressTemplate))
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

	var l []ZyAddressTemplate
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

// UpdateZyAddressTemplate updates ZyAddressTemplate by Id and returns error if
// the record to be updated doesn't exist
func UpdateZyAddressTemplateById(m *ZyAddressTemplate) (err error) {
	o := orm.NewOrm()
	v := ZyAddressTemplate{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZyAddressTemplate deletes ZyAddressTemplate by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZyAddressTemplate(id int) (err error) {
	o := orm.NewOrm()
	v := ZyAddressTemplate{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZyAddressTemplate{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
