package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZyPost struct {
	Id            int       `orm:"column(post_id);auto"`
	PostType      int8      `orm:"column(post_type)"`
	ParentId      int64     `orm:"column(parent_id);null"`
	Title         string    `orm:"column(title);size(100)"`
	IdAlias       string    `orm:"column(id_alias);size(100)"`
	Summary       string    `orm:"column(summary);size(3000);null"`
	Author        string    `orm:"column(author);size(50);null"`
	SourceUrl     string    `orm:"column(source_url);size(250);null"`
	Publisher     string    `orm:"column(publisher);size(50);null"`
	PublishedAt   time.Time `orm:"column(published_at);type(datetime);null"`
	Status        int8      `orm:"column(status);null"`
	CommentAccess int8      `orm:"column(comment_access);null"`
	CommentCount  int       `orm:"column(comment_count);null"`
	ViewCount     int       `orm:"column(view_count);null"`
	AttachedCount int       `orm:"column(attached_count);null"`
	CoverId       int64     `orm:"column(cover_id);null"`
	ExpiredAt     time.Time `orm:"column(expired_at);type(datetime);null"`
	Ordering      int       `orm:"column(ordering);null"`
	Version       int       `orm:"column(version);null"`
	UpdatedBy     string    `orm:"column(updated_by);size(50);null"`
	UpdatedAt     time.Time `orm:"column(updated_at);type(datetime);null"`
	CreatedBy     string    `orm:"column(created_by);size(50);null"`
	CreatedAt     time.Time `orm:"column(created_at);type(datetime);null"`
}

func (t *ZyPost) TableName() string {
	return "zy_post"
}

func init() {
	orm.RegisterModel(new(ZyPost))
}

// AddZyPost insert a new ZyPost into database and returns
// last inserted Id on success.
func AddZyPost(m *ZyPost) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZyPostById retrieves ZyPost by Id. Returns error if
// Id doesn't exist
func GetZyPostById(id int) (v *ZyPost, err error) {
	o := orm.NewOrm()
	v = &ZyPost{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZyPost retrieves all ZyPost matches certain condition. Returns empty list if
// no records exist
func GetAllZyPost(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZyPost))
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

	var l []ZyPost
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

// UpdateZyPost updates ZyPost by Id and returns error if
// the record to be updated doesn't exist
func UpdateZyPostById(m *ZyPost) (err error) {
	o := orm.NewOrm()
	v := ZyPost{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZyPost deletes ZyPost by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZyPost(id int) (err error) {
	o := orm.NewOrm()
	v := ZyPost{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZyPost{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
