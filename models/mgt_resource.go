package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type MgtResource struct {
	Id          int       `orm:"column(id);auto"`
	ParentId    int       `orm:"column(parent_id);null"`
	RootId      int       `orm:"column(root_id);null"`
	ResType     string    `orm:"column(res_type);size(16);null"`
	ResTitle    string    `orm:"column(res_title);size(50)"`
	ResCode     string    `orm:"column(res_code);size(50)"`
	ResUrl      string    `orm:"column(res_url);size(250)"`
	Description string    `orm:"column(description);size(250);null"`
	ResLevel    int       `orm:"column(res_level);null"`
	IsMenu      int8      `orm:"column(is_menu);null"`
	IsInternal  int8      `orm:"column(is_internal);null"`
	IsEnabled   int8      `orm:"column(is_enabled);null"`
	Ordering    int       `orm:"column(ordering);null"`
	UpdatedBy   string    `orm:"column(updated_by);size(50);null"`
	UpdatedAt   time.Time `orm:"column(updated_at);type(datetime)"`
	CreatedBy   string    `orm:"column(created_by);size(50);null"`
	CreatedAt   time.Time `orm:"column(created_at);type(datetime)"`
}

func (t *MgtResource) TableName() string {
	return "mgt_resource"
}

func init() {
	orm.RegisterModel(new(MgtResource))
}

// AddMgtResource insert a new MgtResource into database and returns
// last inserted Id on success.
func AddMgtResource(m *MgtResource) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMgtResourceById retrieves MgtResource by Id. Returns error if
// Id doesn't exist
func GetMgtResourceById(id int) (v *MgtResource, err error) {
	o := orm.NewOrm()
	v = &MgtResource{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMgtResource retrieves all MgtResource matches certain condition. Returns empty list if
// no records exist
func GetAllMgtResource(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MgtResource))
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

	var l []MgtResource
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

// UpdateMgtResource updates MgtResource by Id and returns error if
// the record to be updated doesn't exist
func UpdateMgtResourceById(m *MgtResource) (err error) {
	o := orm.NewOrm()
	v := MgtResource{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMgtResource deletes MgtResource by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMgtResource(id int) (err error) {
	o := orm.NewOrm()
	v := MgtResource{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MgtResource{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
