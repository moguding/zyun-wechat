package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZyUserAccount struct {
	Id             int64     `orm:"column(user_id);auto"`
	UserEmail      string    `orm:"column(user_email);size(50);null"`
	UserName       string    `orm:"column(user_name);size(50)"`
	MobileNo       string    `orm:"column(mobile_no);size(20);null"`
	HashedPasswd   string    `orm:"column(hashed_passwd);size(100);null"`
	PasswdSalt     string    `orm:"column(passwd_salt);size(50);null"`
	PasswdStrength int8      `orm:"column(passwd_strength)"`
	PasswdQuestion string    `orm:"column(passwd_question);size(100);null"`
	PasswdAnswer   string    `orm:"column(passwd_answer);size(100);null"`
	AnswerSalt     string    `orm:"column(answer_salt);size(50);null"`
	Gender         string    `orm:"column(gender);size(1);null"`
	Avatar         string    `orm:"column(avatar);size(250);null"`
	SmallAvatar    string    `orm:"column(small_avatar);size(250);null"`
	MobileVerified int8      `orm:"column(mobile_verified);null"`
	EmailVerified  int8      `orm:"column(email_verified);null"`
	UserType       int8      `orm:"column(user_type);null"`
	ActiveTime     time.Time `orm:"column(active_time);type(datetime);null"`
	Status         int8      `orm:"column(status);null"`
	CreatedIp      string    `orm:"column(created_ip);size(30);null"`
	Version        int       `orm:"column(version);null"`
	UpdatedAt      time.Time `orm:"column(updated_at);type(timestamp);null;auto_now"`
	CreatedAt      time.Time `orm:"column(created_at);type(datetime);null"`
}

func (t *ZyUserAccount) TableName() string {
	return "zy_user_account"
}

func init() {
	orm.RegisterModel(new(ZyUserAccount))
}

// AddZyUserAccount insert a new ZyUserAccount into database and returns
// last inserted Id on success.
func AddZyUserAccount(m *ZyUserAccount) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZyUserAccountById retrieves ZyUserAccount by Id. Returns error if
// Id doesn't exist
func GetZyUserAccountById(id int64) (v *ZyUserAccount, err error) {
	o := orm.NewOrm()
	v = &ZyUserAccount{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetZyUserAccountByMail(mail string) (v *ZyUserAccount, err error) {
	o := orm.NewOrm()
	v = &ZyUserAccount{UserEmail: mail}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZyUserAccount retrieves all ZyUserAccount matches certain condition. Returns empty list if
// no records exist
func GetAllZyUserAccount(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZyUserAccount))
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

	var l []ZyUserAccount
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

// UpdateZyUserAccount updates ZyUserAccount by Id and returns error if
// the record to be updated doesn't exist
func UpdateZyUserAccountById(m *ZyUserAccount) (err error) {
	o := orm.NewOrm()
	v := ZyUserAccount{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZyUserAccount deletes ZyUserAccount by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZyUserAccount(id int64) (err error) {
	o := orm.NewOrm()
	v := ZyUserAccount{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZyUserAccount{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
