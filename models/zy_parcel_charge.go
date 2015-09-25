package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZyParcelCharge struct {
	Id        int       `orm:"column(id);pk"`
	CType     int8      `orm:"column(c_type);null"`
	CKey      string    `orm:"column(c_key);size(50);null"`
	CValue    string    `orm:"column(c_value);size(50);null"`
	CVip      string    `orm:"column(c_vip);size(50);null"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp);null"`
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);auto_now_add"`
}

func (t *ZyParcelCharge) TableName() string {
	return "zy_parcel_charge"
}

func init() {
	orm.RegisterModel(new(ZyParcelCharge))
}

// AddZyParcelCharge insert a new ZyParcelCharge into database and returns
// last inserted Id on success.
func AddZyParcelCharge(m *ZyParcelCharge) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZyParcelChargeById retrieves ZyParcelCharge by Id. Returns error if
// Id doesn't exist
func GetZyParcelChargeById(id int) (v *ZyParcelCharge, err error) {
	o := orm.NewOrm()
	v = &ZyParcelCharge{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZyParcelCharge retrieves all ZyParcelCharge matches certain condition. Returns empty list if
// no records exist
func GetAllZyParcelCharge(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZyParcelCharge))
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

	var l []ZyParcelCharge
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

// UpdateZyParcelCharge updates ZyParcelCharge by Id and returns error if
// the record to be updated doesn't exist
func UpdateZyParcelChargeById(m *ZyParcelCharge) (err error) {
	o := orm.NewOrm()
	v := ZyParcelCharge{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZyParcelCharge deletes ZyParcelCharge by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZyParcelCharge(id int) (err error) {
	o := orm.NewOrm()
	v := ZyParcelCharge{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZyParcelCharge{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
