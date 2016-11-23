/*
* @Author: GuoDi
* @Date:   2016-11-23 00:57:49
* @Last Modified by:   GuoDi
* @Last Modified time: 2016-11-23 02:38:44
 */

package account

import (
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	ErrorDefault  = 500
	PasswordError = 400
	ForBiddenUser = 401
	LoginSuccess  = 200
	ResetSuccess  = 201
)

type Log struct {
	Id        int64
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"  json:"createdAt"`
	Ip        string    `orm:"size(16)"`
	Message   string    `orm:"null;size(1024)"`
	User      *User     `orm:"rel(fk)"`
	Status    int       `orm:"default(100)"`
}

func init() {
	orm.RegisterModel(new(Log))
}

func (this *Log) TableName() string {
	return "user_login_log"
}

func (this *Log) Create() error {
	_, err := orm.NewOrm().Insert(this)
	return err
}
func (this *Log) ListByUser(startAt time.Time, uid int64) (logs []*Log) {
	orm.NewOrm().QueryTable(this).Filter("User__Id", uid).Filter("CreatedAt__gte", startAt).OrderBy("-CreatedAt").Limit(10).All(logs)
	return
}
