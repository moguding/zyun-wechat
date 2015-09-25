package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZyLink struct {
	Id        int       `orm:"column(id);auto"`
	LinkUrl   string    `orm:"column(link_url);size(250)"`
	LinkTitle string    `orm:"column(link_title);size(250);null"`
	LinkText  string    `orm:"column(link_text);size(250);null"`
	LinkType  string    `orm:"column(link_type);size(20);null"`
	LinkImage string    `orm:"column(link_image);size(250);null"`
	LinkInfo  string    `orm:"column(link_info);size(250);null"`
	LinkProps string    `orm:"column(link_props);size(1000);null"`
	IsVisible int8      `orm:"column(is_visible);null"`
	UpdatedBy string    `orm:"column(updated_by);size(50);null"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp);null;auto_now"`
	CreatedBy string    `orm:"column(created_by);size(50);null"`
	CreatedAt time.Time `orm:"column(created_at);type(datetime);null"`
}

func (t *ZyLink) TableName() string {
	return "zy_link"
}

func init() {
	orm.RegisterModel(new(ZyLink))
}

// AddZyLink insert a new ZyLink into database and returns
// last inserted Id on success.
func AddZyLink(m *ZyLink) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZyLinkById retrieves ZyLink by Id. Returns error if
// Id doesn't exist
func GetZyLinkById(id int) (v *ZyLink, err error) {
	o := orm.NewOrm()
	v = &ZyLink{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZyLink retrieves all ZyLink matches certain condition. Returns empty list if
// no records exist
func GetAllZyLink(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZyLink))
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

	var l []ZyLink
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

// UpdateZyLink updates ZyLink by Id and returns error if
// the record to be updated doesn't exist
func UpdateZyLinkById(m *ZyLink) (err error) {
	o := orm.NewOrm()
	v := ZyLink{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZyLink deletes ZyLink by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZyLink(id int) (err error) {
	o := orm.NewOrm()
	v := ZyLink{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZyLink{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
