package account

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Profile struct {
	Id         int64
	CreatedAt  time.Time `orm:"auto_now_add;type(datetime)"  json:"createdAt"`
	ModifiedAt time.Time `orm:"null;type(datetime)"  json:"modifiedAt"`
	Nickname   string    `orm:"null;size(18)" form:"Nickname" json:"nickname"`
	HeadIcon   string    `orm:"null;size(256)" form:"HeadIcon" json:"headIcon"`
	Gender     string    `orm:"null;size(16)" form:"Gender" json:"gender"`
	Address    string    `orm:"null;size(256)" form:"Address" json:"address"`
	City       string    `orm:"null;size(64)" form:"City" json:"city"`
	Dob        time.Time `orm:"null;type(date)" form:"Dob" json:"dob"`
	User       *User     `orm:"reverse(one)"`
}

func init() {
	orm.RegisterModel(new(Profile))
}

func (this *Profile) TableName() string {
	return "user_profiles"
}
func (this *Profile) Create() error {
	_, err := orm.NewOrm().Insert(this)
	return err
}
func (this *Profile) Update(fields ...string) error {
	if this.Id <= 0 {
		return error.New("ID错误，无法更新")
	}
	this.ModifiedAt = time.Now()
	_, err := orm.NewOrm().Update(this, fields...)
	return err
}
