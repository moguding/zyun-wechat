package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZyMenu struct {
	Id         int       `orm:"column(id);auto"`
	Type       string    `orm:"column(type);size(30)"`
	SrcId      int64     `orm:"column(src_id)"`
	RootId     int       `orm:"column(root_id);null"`
	ParentId   int       `orm:"column(parent_id);null"`
	Label      string    `orm:"column(label);size(50);null"`
	Url        string    `orm:"column(url);size(200);null"`
	Title      string    `orm:"column(title);size(50);null"`
	Level      int       `orm:"column(level);null"`
	Info       string    `orm:"column(info);size(250);null"`
	Attributes string    `orm:"column(attributes);size(500);null"`
	IsDisplay  int8      `orm:"column(is_display);null"`
	UpdatedAt  time.Time `orm:"column(updated_at);type(datetime);null"`
	CreatedAt  time.Time `orm:"column(created_at);type(datetime);null"`
}

func (t *ZyMenu) TableName() string {
	return "zy_menu"
}

func init() {
	orm.RegisterModel(new(ZyMenu))
}

// AddZyMenu insert a new ZyMenu into database and returns
// last inserted Id on success.
func AddZyMenu(m *ZyMenu) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZyMenuById retrieves ZyMenu by Id. Returns error if
// Id doesn't exist
func GetZyMenuById(id int) (v *ZyMenu, err error) {
	o := orm.NewOrm()
	v = &ZyMenu{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZyMenu retrieves all ZyMenu matches certain condition. Returns empty list if
// no records exist
func GetAllZyMenu(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZyMenu))
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

	var l []ZyMenu
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

// UpdateZyMenu updates ZyMenu by Id and returns error if
// the record to be updated doesn't exist
func UpdateZyMenuById(m *ZyMenu) (err error) {
	o := orm.NewOrm()
	v := ZyMenu{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZyMenu deletes ZyMenu by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZyMenu(id int) (err error) {
	o := orm.NewOrm()
	v := ZyMenu{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZyMenu{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
