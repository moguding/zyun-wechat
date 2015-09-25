package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZyAbroadParcel struct {
	Id           int       `orm:"column(id);auto"`
	UserId       int64     `orm:"column(user_id);null"`
	UserMark     string    `orm:"column(user_mark);size(100);null"`
	FullName     string    `orm:"column(full_name);size(50);null"`
	TrackingNo   string    `orm:"column(tracking_no);size(50);null"`
	Express      int8      `orm:"column(express);null"`
	Status       int8      `orm:"column(status);null"`
	Position     string    `orm:"column(position);size(20);null"`
	ParcelDesc   string    `orm:"column(parcel_desc);size(1000);null"`
	ItemsInfo    string    `orm:"column(items_info);size(1000);null"`
	Origin       string    `orm:"column(origin);size(100);null"`
	Weight       int       `orm:"column(weight);null"`
	AbnormalType int8      `orm:"column(abnormal_type);null"`
	Unknown      int8      `orm:"column(unknown);null"`
	ReturnAmount float64   `orm:"column(return_amount);null;digits(10);decimals(2)"`
	InStoreTime  time.Time `orm:"column(in_store_time);type(datetime);null"`
	UpdatedAt    time.Time `orm:"column(updated_at);type(datetime);null"`
	CreatedAt    time.Time `orm:"column(created_at);type(datetime);null"`
	Version      int       `orm:"column(version);null"`
	CreatedBy    string    `orm:"column(created_by);size(50);null"`
	UpdatedBy    string    `orm:"column(updated_by);size(50);null"`
}

func (t *ZyAbroadParcel) TableName() string {
	return "zy_abroad_parcel"
}

func init() {
	orm.RegisterModel(new(ZyAbroadParcel))
}

// AddZyAbroadParcel insert a new ZyAbroadParcel into database and returns
// last inserted Id on success.
func AddZyAbroadParcel(m *ZyAbroadParcel) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZyAbroadParcelById retrieves ZyAbroadParcel by Id. Returns error if
// Id doesn't exist
func GetZyAbroadParcelById(id int) (v *ZyAbroadParcel, err error) {
	o := orm.NewOrm()
	v = &ZyAbroadParcel{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZyAbroadParcel retrieves all ZyAbroadParcel matches certain condition. Returns empty list if
// no records exist
func GetAllZyAbroadParcel(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZyAbroadParcel))
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

	var l []ZyAbroadParcel
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

// UpdateZyAbroadParcel updates ZyAbroadParcel by Id and returns error if
// the record to be updated doesn't exist
func UpdateZyAbroadParcelById(m *ZyAbroadParcel) (err error) {
	o := orm.NewOrm()
	v := ZyAbroadParcel{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZyAbroadParcel deletes ZyAbroadParcel by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZyAbroadParcel(id int) (err error) {
	o := orm.NewOrm()
	v := ZyAbroadParcel{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZyAbroadParcel{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
