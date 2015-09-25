package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZyCategory struct {
	Id            int       `orm:"column(id);auto"`
	Taxonomy      string    `orm:"column(taxonomy);size(30)"`
	CateName      string    `orm:"column(cate_name);size(50)"`
	CateAlias     string    `orm:"column(cate_alias);size(50)"`
	RootId        int64     `orm:"column(root_id);null"`
	ParentId      int64     `orm:"column(parent_id);null"`
	CateDetail    string    `orm:"column(cate_detail);size(200);null"`
	Templates     string    `orm:"column(templates);size(500);null"`
	StatsCount    int64     `orm:"column(stats_count);null"`
	Extensions    string    `orm:"column(extensions);size(500);null"`
	ChildrenCount int       `orm:"column(children_count);null"`
	Ordering      int       `orm:"column(ordering);null"`
	CreatedAt     time.Time `orm:"column(created_at);type(datetime)"`
	CreatedBy     string    `orm:"column(created_by);size(50)"`
	UpdatedAt     time.Time `orm:"column(updated_at);type(datetime)"`
	UpdatedBy     string    `orm:"column(updated_by);size(50)"`
}

func (t *ZyCategory) TableName() string {
	return "zy_category"
}

func init() {
	orm.RegisterModel(new(ZyCategory))
}

// AddZyCategory insert a new ZyCategory into database and returns
// last inserted Id on success.
func AddZyCategory(m *ZyCategory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZyCategoryById retrieves ZyCategory by Id. Returns error if
// Id doesn't exist
func GetZyCategoryById(id int) (v *ZyCategory, err error) {
	o := orm.NewOrm()
	v = &ZyCategory{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZyCategory retrieves all ZyCategory matches certain condition. Returns empty list if
// no records exist
func GetAllZyCategory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZyCategory))
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

	var l []ZyCategory
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

// UpdateZyCategory updates ZyCategory by Id and returns error if
// the record to be updated doesn't exist
func UpdateZyCategoryById(m *ZyCategory) (err error) {
	o := orm.NewOrm()
	v := ZyCategory{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZyCategory deletes ZyCategory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZyCategory(id int) (err error) {
	o := orm.NewOrm()
	v := ZyCategory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZyCategory{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
