package content

import (
	"github.com/astaxie/beego/orm"
	"github.com/wayn3h0/go-uuid"
	"github.com/ykrl089/beelib/account"
)

type Content struct {
	Id         int64
	CreatedAt  time.Time     `orm:"auto_now_add;type(datetime)"  json:"createdAt"`
	ModifiedAt time.Time     `orm:"null;type(datetime)"  json:"modifiedAt"`
	CreatedBy  *account.User `orm:"rel(fk)"`
	ModifiedBy *account.User `orm:"rel(fk)"`
	Uuid       string        `orm:"size(36);unique;index" json:"uuid"`
	Title      string        `orm:"size(256)" form:"Title" json:"title"`
	Desc       string        `orm:"size(512)" form:"Desc" json:"desc"`
	Image      string        `orm:"null;size(256)" json:"Image"`
	Article    string        `orm:"-" form:"Article" json:"article"`
	Md5        string        `orm:"null;size(64)"`
}

func (this *Content) TableName() string {
	return "content"
}
