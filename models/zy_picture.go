package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZyPicture struct {
	Id        int       `orm:"column(id);auto"`
	BaseUri   string    `orm:"column(base_uri);size(100);null"`
	RelPath   string    `orm:"column(rel_path);size(100);null"`
	AbsPath   string    `orm:"column(abs_path);size(200);null"`
	Extension string    `orm:"column(extension);size(10);null"`
	MimeType  string    `orm:"column(mime_type);size(20);null"`
	Size      int64     `orm:"column(size);null"`
	Sha1      string    `orm:"column(sha1);size(40);null"`
	Location  int8      `orm:"column(location);null"`
	IsEnabled int8      `orm:"column(is_enabled);null"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp);auto_now_add"`
	CreatedAt time.Time `orm:"column(created_at);type(datetime);null"`
}

func (t *ZyPicture) TableName() string {
	return "zy_picture"
}

func init() {
	orm.RegisterModel(new(ZyPicture))
}

// AddZyPicture insert a new ZyPicture into database and returns
// last inserted Id on success.
func AddZyPicture(m *ZyPicture) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZyPictureById retrieves ZyPicture by Id. Returns error if
// Id doesn't exist
func GetZyPictureById(id int) (v *ZyPicture, err error) {
	o := orm.NewOrm()
	v = &ZyPicture{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZyPicture retrieves all ZyPicture matches certain condition. Returns empty list if
// no records exist
func GetAllZyPicture(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZyPicture))
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

	var l []ZyPicture
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

// UpdateZyPicture updates ZyPicture by Id and returns error if
// the record to be updated doesn't exist
func UpdateZyPictureById(m *ZyPicture) (err error) {
	o := orm.NewOrm()
	v := ZyPicture{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZyPicture deletes ZyPicture by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZyPicture(id int) (err error) {
	o := orm.NewOrm()
	v := ZyPicture{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZyPicture{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
