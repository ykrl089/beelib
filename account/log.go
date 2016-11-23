/*
* @Author: GuoDi
* @Date:   2016-11-23 00:57:49
* @Last Modified by:   GuoDi
* @Last Modified time: 2016-11-23 23:19:47
 */

package account

import (
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	PasswordError = 400
	ForBiddenUser = 401
	StatusInit    = 300
	ResetSuccess  = 200
	LogoutSuccess = 102
	LoginSuccess  = 100
)

type Log struct {
	Id        int64
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"  json:"createdAt"`
	Ip        string    `null;orm:"size(16)"`
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
	if EnableLog {
		_, err := orm.NewOrm().Insert(this)
		return err
	} else {
		return nil
	}

}

func (this *Log) List(startAt time.Time, uid int64) (logs []*Log) {
	orm.NewOrm().QueryTable(this).Filter("User__Id", uid).Filter("CreatedAt__lte", startAt).OrderBy("-CreatedAt").Limit(10).All(&logs)
	return
}
func (this *Log) ErrorCountLimitedInHour(uid int64, countLimit int) bool {
	if !EnableLog {
		return true
	}
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
	return true
}
