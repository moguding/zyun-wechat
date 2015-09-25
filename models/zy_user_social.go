package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZyUserSocial struct {
	Id           int       `orm:"column(id);auto"`
	UserId       int64     `orm:"column(user_id)"`
	UserName     string    `orm:"column(user_name);size(50)"`
	OpenId       string    `orm:"column(open_id);size(50);null"`
	SocialSource string    `orm:"column(social_source);size(20);null"`
	AuthToken    string    `orm:"column(auth_token);size(50);null"`
	Period       int64     `orm:"column(period);null"`
	TokenTime    time.Time `orm:"column(token_time);type(timestamp);null"`
	UpdatedAt    time.Time `orm:"column(updated_at);type(timestamp);null;auto_now"`
	CreatedAt    time.Time `orm:"column(created_at);type(datetime);null"`
}

func (t *ZyUserSocial) TableName() string {
	return "zy_user_social"
}

func init() {
	orm.RegisterModel(new(ZyUserSocial))
}

// AddZyUserSocial insert a new ZyUserSocial into database and returns
// last inserted Id on success.
func AddZyUserSocial(m *ZyUserSocial) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZyUserSocialById retrieves ZyUserSocial by Id. Returns error if
// Id doesn't exist
func GetZyUserSocialById(id int) (v *ZyUserSocial, err error) {
	o := orm.NewOrm()
	v = &ZyUserSocial{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZyUserSocial retrieves all ZyUserSocial matches certain condition. Returns empty list if
// no records exist
func GetAllZyUserSocial(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZyUserSocial))
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

	var l []ZyUserSocial
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

// UpdateZyUserSocial updates ZyUserSocial by Id and returns error if
// the record to be updated doesn't exist
func UpdateZyUserSocialById(m *ZyUserSocial) (err error) {
	o := orm.NewOrm()
	v := ZyUserSocial{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZyUserSocial deletes ZyUserSocial by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZyUserSocial(id int) (err error) {
	o := orm.NewOrm()
	v := ZyUserSocial{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZyUserSocial{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
