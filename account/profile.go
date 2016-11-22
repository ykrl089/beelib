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
	Email      string    `orm:"null;size(256)" form:"Email" json:"email"`
	Dob        time.Time `orm:"null;type(date)" form:"Dob" json:"dob"`
	User       *User     `orm:"reverse(one)"`
}

func init() {
	orm.RegisterModel(new(Profile))
}

func (this *Profile) TableName() string {
	return "user_profiles"
}
