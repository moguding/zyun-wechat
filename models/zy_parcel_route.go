package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZyParcelRoute struct {
	Id           int       `orm:"column(id);auto"`
	ParcelId     int64     `orm:"column(parcel_id)"`
	RouteDesc    string    `orm:"column(route_desc);size(1000)"`
	RouteStatus  int8      `orm:"column(route_status);null"`
	DeliveryDesc string    `orm:"column(delivery_desc);size(2000);null"`
	CreatedAt    time.Time `orm:"column(created_at);type(datetime);auto_now_add"`
}

func (t *ZyParcelRoute) TableName() string {
	return "zy_parcel_route"
}

func init() {
	orm.RegisterModel(new(ZyParcelRoute))
}

// AddZyParcelRoute insert a new ZyParcelRoute into database and returns
// last inserted Id on success.
func AddZyParcelRoute(m *ZyParcelRoute) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZyParcelRouteById retrieves ZyParcelRoute by Id. Returns error if
// Id doesn't exist
func GetZyParcelRouteById(id int) (v *ZyParcelRoute, err error) {
	o := orm.NewOrm()
	v = &ZyParcelRoute{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZyParcelRoute retrieves all ZyParcelRoute matches certain condition. Returns empty list if
// no records exist
func GetAllZyParcelRoute(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZyParcelRoute))
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

	var l []ZyParcelRoute
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

// UpdateZyParcelRoute updates ZyParcelRoute by Id and returns error if
// the record to be updated doesn't exist
func UpdateZyParcelRouteById(m *ZyParcelRoute) (err error) {
	o := orm.NewOrm()
	v := ZyParcelRoute{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZyParcelRoute deletes ZyParcelRoute by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZyParcelRoute(id int) (err error) {
	o := orm.NewOrm()
	v := ZyParcelRoute{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZyParcelRoute{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
