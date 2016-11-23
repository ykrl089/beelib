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
	StatusInit    = 300
	PasswordError = 401
	ForBiddenUser = 400
	ErrorReset    = 200
	LoginSuccess  = 100
)

type Log struct {
	Id        int64
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"  json:"createdAt"`
	Ip        string    `orm:"size(16)"`
	User      *User     `orm:"rel(fk)"`
	Status    int       `orm:"default(300)"`
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
	orm.NewOrm().QueryTable(this).Filter("User__Id", uid).Filter("CreatedAt__gte", startAt).OrderBy("-CreatedAt").Limit(10).All(logs)
	return
}
func (this *Log) ErrorCountLimitedInHour(uid int64, countLimit int) bool {
	hour, _ := time.ParseDuration("-1h")
	hourBefore := time.Now().Add(hour)
	logs := this.List(hourBefore, uid)
	if count := len(logs); count <= 0 {
		return true
	}
	for i, log := range logs {
		if i == countLimit {
			return false
		}
		if log.Status <= 200 {
			return true
		}
	}
}
