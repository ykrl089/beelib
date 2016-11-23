package account

import (
	"github.com/astaxie/beego/orm"
	"github.com/wayn3h0/go-uuid"
	_ "github.com/ykrl089/beelib/library/md5"
	_ "github.com/ykrl089/beelib/library/str"
	"time"
)

type User struct {
	Id            int64
	CreatedAt     time.Time `orm:"auto_now_add;type(datetime)"  json:"createdAt"`
	ModifiedAt    time.Time `orm:"null;type(datetime)"  json:"modifiedAt"`
	Username      string    `orm:"size(16);unique;index" form:"Username" json:"username"`
	PasswordPlain string    `orm:"-" form:"Password" json:"-"`
	PasswordCript string    `orm:"size(64);"  json:"-"`
	LoginToken    string    `orm:"null;size(32)" form:"LoginToken" json:"loginToken"`
	ExpiredAt     time.Time `orm:"null;type(datetime)"  json:"expiredAt"`
	Uuid          string    `orm:"size(36);unique;index" json:"uuid"`
	Salt          string    `orm:"size(18)" json:"-"`
	Nickname      string    `orm:"null;size(18)" form:"Nickname" json:"nickname"`
	HeadIcon      string    `orm:"null;size(256)" form:"HeadIcon" json:"headIcon"`
	Email         string    `orm:"null;size(256);index" form:"Email" json:"email"`
	Role          string    `orm:"null;size(64)" json:"role"`
	Profile       *Profile  `orm:"null;rel(one)"`
}

var (
	EnableLog   bool
	EnableCache bool
)

func init() {
	orm.RegisterModel(new(User))
	EnableLog = false
	EnableCache = true
}

func (this *User) TableName() string {
	return "users"
}

func (this *User) Login() error {
	if err = this.Get(); err != nil {
		if err == orm.ErrNoRows {
			return error.New("查询不到")
		} else if err == orm.ErrMissPK {
			return error.New("找不到对应的ID")
		}
	}
	if MD5(MD5(this.PasswordPlain)+this.Salt) == this.PasswordCript {
		this.LoginToken = RanStr(32)
		this.ExpiredAt = time.Now.Add(720 * time.Hour) //30天过期
		return this.Update("LoginToken", "ExpiredAt")
	} else {
		return error.New("密码错误")
	}

}

func (this *User) Get() error {
	o := orm.NewOrm()
	if this.Id > 0 {
		return o.Read(this)
	} else if this.Username != "" {
		return o.Read(this, "Username")
	} else if this.Uuid != "" {
		return o.Read(this, "Uuid")
	} else if this.Email != "" {
		return o.Read(this, "Email")
	} else {
		return Error.New("无效查询，请输入ID，Username、UUID或者Email")
	}

}
func (this *User) Update(fields ...string) error {
	o := orm.NewOrm()
	if this.Id <= 0 {
		return error.New("ID错误")
	}
	_, err := o.Update(this, fields...)
	return error.New("更新失败")
}
func (this *User) Create() error {
	o := orm.NewOrm()
	if this.Id > 0 {
		return error.New("无法创建已有主键的用户")
	}
	if len(this.PasswordPlain) < 6 {
		return error.New("密码长度过短")
	}
	if this.Username == "" && this.Email == "" {
		return error.New("未设置用户名和Email地址")
	}
	this.Salt = RandStr(18)
	this.PasswordCript = MD5(MD5(this.PasswordPlain) + this.Salt)
	this.Uuid = uuidv4.String()
	if _, err := o.Insert(this); err != nil {
		return error.New("创建失败")
	}
	return nil

}
func (this *User) Current() error {
	if len(this.LoginToken) != 32 {
		return error.New("Token 错误，请重新登录")
	}
	o := orm.NewOrm()
	if err := o.Read(this, "LoginToken"); err == orm.ErrNoRows {
		return error.New("Token 错误，请重新登录")
	} else {
		if this.ExpiredAt > time.Now() {
			this.Logout
			return error.New("用户登录已过期，请重新登录")
		} else {
			return nil
		}
	}
}
func (this *User) Logout() {
	this.LoginToken = ""
	this.ExpiredAt = time.Now()
	this.Update("LoginToken", "ExpiredAt")
}
