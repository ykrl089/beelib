package content

import (
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id          int64
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"  json:"createdAt"`
	ContentUuid string    `orm:"size(36);" json:"contentUuid"`
	Article     string    `orm:"type(text)"`
	Md5         string    `orm:"size(36);" `
}

func (this *Article) TableName() string {
	return "article"
}
