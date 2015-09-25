package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZyDomesticParcel struct {
	Id             int       `orm:"column(id);auto"`
	UserId         int64     `orm:"column(user_id);null"`
	TrackingNo     string    `orm:"column(tracking_no);size(50);null"`
	TransportNo    string    `orm:"column(transport_no);size(50);null"`
	Position       string    `orm:"column(position);size(20);null"`
	Delivery       string    `orm:"column(delivery);size(50);null"`
	ItemsInfo      string    `orm:"column(items_info);size(1000);null"`
	ParcelValue    float64   `orm:"column(parcel_value);null;digits(10);decimals(2)"`
	Weight         int       `orm:"column(weight);null"`
	TotalFee       float64   `orm:"column(total_fee);null;digits(10);decimals(2)"`
	ChargeInfo     string    `orm:"column(charge_info);size(500);null"`
	Recipient      string    `orm:"column(recipient);size(500);null"`
	Status         int8      `orm:"column(status);null"`
	Remark         string    `orm:"column(remark);size(500);null"`
	OutStoreTime   time.Time `orm:"column(out_store_time);type(datetime);null"`
	HandlingType   int8      `orm:"column(handling_type);null"`
	HandlingSource string    `orm:"column(handling_source);size(200);null"`
	Version        int       `orm:"column(version);null"`
	UpdatedAt      time.Time `orm:"column(updated_at);type(timestamp);null"`
	CreatedAt      time.Time `orm:"column(created_at);type(timestamp);null"`
	CreatedBy      string    `orm:"column(created_by);size(50);null"`
	UpdatedBy      string    `orm:"column(updated_by);size(50);null"`
}

func (t *ZyDomesticParcel) TableName() string {
	return "zy_domestic_parcel"
}

func init() {
	orm.RegisterModel(new(ZyDomesticParcel))
}

// AddZyDomesticParcel insert a new ZyDomesticParcel into database and returns
// last inserted Id on success.
func AddZyDomesticParcel(m *ZyDomesticParcel) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZyDomesticParcelById retrieves ZyDomesticParcel by Id. Returns error if
// Id doesn't exist
func GetZyDomesticParcelById(id int) (v *ZyDomesticParcel, err error) {
	o := orm.NewOrm()
	v = &ZyDomesticParcel{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZyDomesticParcel retrieves all ZyDomesticParcel matches certain condition. Returns empty list if
// no records exist
func GetAllZyDomesticParcel(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZyDomesticParcel))
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

	var l []ZyDomesticParcel
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

// UpdateZyDomesticParcel updates ZyDomesticParcel by Id and returns error if
// the record to be updated doesn't exist
func UpdateZyDomesticParcelById(m *ZyDomesticParcel) (err error) {
	o := orm.NewOrm()
	v := ZyDomesticParcel{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZyDomesticParcel deletes ZyDomesticParcel by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZyDomesticParcel(id int) (err error) {
	o := orm.NewOrm()
	v := ZyDomesticParcel{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZyDomesticParcel{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
