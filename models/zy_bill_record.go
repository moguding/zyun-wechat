package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZyBillRecord struct {
	Id           int       `orm:"column(id);auto"`
	UserId       int64     `orm:"column(user_id)"`
	OrderNo      string    `orm:"column(order_no);size(50)"`
	OrderAmount  float64   `orm:"column(order_amount);digits(10);decimals(2)"`
	OrderType    int8      `orm:"column(order_type)"`
	ActualAmount float64   `orm:"column(actual_amount);digits(10);decimals(2)"`
	ChannelId    int       `orm:"column(channel_id);null"`
	ChannelName  string    `orm:"column(channel_name);size(50);null"`
	TradeNo      string    `orm:"column(trade_no);size(50);null"`
	TradeResult  string    `orm:"column(trade_result);size(50)"`
	TradeNote    string    `orm:"column(trade_note);size(200);null"`
	Sign         string    `orm:"column(sign);size(100)"`
	SignType     string    `orm:"column(sign_type);size(20)"`
	RecordStatus int8      `orm:"column(record_status)"`
	UserIp       string    `orm:"column(user_ip);size(50);null"`
	IsOffset     int8      `orm:"column(is_offset);null"`
	Remark       string    `orm:"column(remark);size(250);null"`
	UpdatedAt    time.Time `orm:"column(updated_at);type(datetime);null"`
	CreatedAt    time.Time `orm:"column(created_at);type(datetime);auto_now_add"`
}

func (t *ZyBillRecord) TableName() string {
	return "zy_bill_record"
}

func init() {
	orm.RegisterModel(new(ZyBillRecord))
}

// AddZyBillRecord insert a new ZyBillRecord into database and returns
// last inserted Id on success.
func AddZyBillRecord(m *ZyBillRecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZyBillRecordById retrieves ZyBillRecord by Id. Returns error if
// Id doesn't exist
func GetZyBillRecordById(id int) (v *ZyBillRecord, err error) {
	o := orm.NewOrm()
	v = &ZyBillRecord{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZyBillRecord retrieves all ZyBillRecord matches certain condition. Returns empty list if
// no records exist
func GetAllZyBillRecord(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZyBillRecord))
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

	var l []ZyBillRecord
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

// UpdateZyBillRecord updates ZyBillRecord by Id and returns error if
// the record to be updated doesn't exist
func UpdateZyBillRecordById(m *ZyBillRecord) (err error) {
	o := orm.NewOrm()
	v := ZyBillRecord{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZyBillRecord deletes ZyBillRecord by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZyBillRecord(id int) (err error) {
	o := orm.NewOrm()
	v := ZyBillRecord{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZyBillRecord{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
