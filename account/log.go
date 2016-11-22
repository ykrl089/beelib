/*
* @Author: GuoDi
* @Date:   2016-11-23 00:57:49
* @Last Modified by:   GuoDi
* @Last Modified time: 2016-11-23 01:29:26
 */

package account

import (
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	PasswordError = "密码错误"
	ForBiddenUser = "禁用用户"
	LoginSuccess  = "用户登录成功"
)

type Log struct {
	Id        int64
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"  json:"createdAt"`
	Ip        string    `orm:"size(16)"`
	Message   string    `orm:"null;size(1024)"`
	User      *User     `orm:"rel(fk)"`
	Status    bool      `orm:"default(1)"`
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
func (this *Log) List(startAt time.Time, uid int64) (logs []*Log) {
	orm.NewOrm().QueryTable(this).Filter("User__Id", uid).Filter("CreatedAt__gte", startAt).OrderBy("-CreatedAt").Limit(100).All(logs)
	return
}
