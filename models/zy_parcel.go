package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZyParcel struct {
	Id           int       `orm:"column(id);auto"`
	UserId       int64     `orm:"column(user_id);null"`
	UserCode     string    `orm:"column(user_code);size(50);null"`
	TrackingNo   string    `orm:"column(tracking_no);size(50);null"`
	Express      int8      `orm:"column(express);null"`
	Tracking     string    `orm:"column(tracking);size(500);null"`
	FullName     string    `orm:"column(full_name);size(500);null"`
	Origin       string    `orm:"column(origin);size(1000);null"`
	Position     string    `orm:"column(position);size(50);null"`
	Items        string    `orm:"column(items);size(500);null"`
	Weight       int       `orm:"column(weight);null"`
	ReturnAmount float64   `orm:"column(return_amount);null;digits(8);decimals(2)"`
	InStoreTime  time.Time `orm:"column(in_store_time);type(datetime);null"`
	AbnormalType int8      `orm:"column(abnormal_type);null"`
	ParcelStatus int8      `orm:"column(parcel_status)"`
	TransportNo  string    `orm:"column(transport_no);size(50);null"`
	Delivery     string    `orm:"column(delivery);size(50);null"`
	DeliveryFee  float64   `orm:"column(delivery_fee);null;digits(10);decimals(2)"`
	Recipient    string    `orm:"column(recipient);size(500);null"`
	ParcelValue  float64   `orm:"column(parcel_value);null;digits(10);decimals(2)"`
	OutStoreTime time.Time `orm:"column(out_store_time);type(datetime);null"`
	Charges      string    `orm:"column(charges);size(200);null"`
	HandlingType int8      `orm:"column(handling_type);null"`
	Remark       string    `orm:"column(remark);size(200);null"`
	Version      int       `orm:"column(version)"`
	UpdatedAt    time.Time `orm:"column(updated_at);type(timestamp);null"`
	CreatedAt    time.Time `orm:"column(created_at);type(timestamp);auto_now_add"`
	CreatedBy    string    `orm:"column(created_by);size(50);null"`
	UpdatedBy    string    `orm:"column(updated_by);size(50);null"`
}

func (t *ZyParcel) TableName() string {
	return "zy_parcel"
}

func init() {
	orm.RegisterModel(new(ZyParcel))
}

// AddZyParcel insert a new ZyParcel into database and returns
// last inserted Id on success.
func AddZyParcel(m *ZyParcel) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZyParcelById retrieves ZyParcel by Id. Returns error if
// Id doesn't exist
func GetZyParcelById(id int) (v *ZyParcel, err error) {
	o := orm.NewOrm()
	v = &ZyParcel{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZyParcel retrieves all ZyParcel matches certain condition. Returns empty list if
// no records exist
func GetAllZyParcel(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZyParcel))
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

	var l []ZyParcel
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

// UpdateZyParcel updates ZyParcel by Id and returns error if
// the record to be updated doesn't exist
func UpdateZyParcelById(m *ZyParcel) (err error) {
	o := orm.NewOrm()
	v := ZyParcel{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZyParcel deletes ZyParcel by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZyParcel(id int) (err error) {
	o := orm.NewOrm()
	v := ZyParcel{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZyParcel{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
