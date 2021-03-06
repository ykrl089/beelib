package account

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/wayn3h0/go-uuid"
	"github.com/ykrl089/beelib/library/md5"
	"github.com/ykrl089/beelib/library/str"
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
	IP            string    `orm:"-"` // 用户登录IP
	Nickname      string    `orm:"null;size(18)" form:"Nickname" json:"nickname"`
	HeadIcon      string    `orm:"null;size(256)" form:"HeadIcon" json:"headIcon"`
	Email         string    `orm:"null;size(256);index" form:"Email" json:"email"`
	Role          string    `orm:"null;size(64)" json:"role"`
	Profile       *Profile  `orm:"null;rel(one);on_delete(set_null)"`
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
	password := this.PasswordPlain
	if password == "" {
		return errors.New("未输入密码")
	}
	if err := this.Get(); err != nil {
		if err == orm.ErrNoRows {
			return errors.New("查询不到")
		} else if err == orm.ErrMissPK {
			return errors.New("找不到对应的ID")
		}
	}

	loginLog := Log{
		Ip:   this.IP,
		User: this,
	}
	if ok := loginLog.ErrorCountLimitedInHour(this.Id, 5); !ok {
		return errors.New("登录错误次数超限，请联系管理员")
	}
	fmt.Println("userpassword", this.PasswordPlain)
	if md5.MD5(md5.MD5(password)+this.Salt) == this.PasswordCript {
		loginLog.Status = LoginSuccess
		this.LoginToken = str.RanStr(32)
		this.ExpiredAt = time.Now().Add(720 * time.Hour) //30天过期
		loginLog.Create()
		return this.Update("LoginToken", "ExpiredAt")
	} else {
		loginLog.Status = PasswordError
		loginLog.Create()
		return errors.New("密码错误")
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
		return errors.New("无效查询，请输入ID，Username、UUID或者Email")
	}

}
func (this *User) Update(fields ...string) error {
	o := orm.NewOrm()
	if this.Id <= 0 {
		return errors.New("ID错误")
	}
	if _, err := o.Update(this, fields...); err != nil {
		return errors.New("更新失败")
	} else {
		return nil
	}
}
func (this *User) Create() error {
	o := orm.NewOrm()
	if this.Id > 0 {
		return errors.New("无法创建已有主键的用户")
	}
	if len(this.PasswordPlain) < 6 {
		return errors.New("密码长度过短")
	}
	if this.Username == "" && this.Email == "" {
		return errors.New("未设置用户名和Email地址")
	}
	this.Salt = str.RanStr(18)
	this.PasswordCript = md5.MD5(md5.MD5(this.PasswordPlain) + this.Salt)
	uuidv4, _ := uuid.NewV4()
	this.Uuid = uuidv4.String()
	if _, err := o.Insert(this); err != nil {
		return errors.New("创建失败")
	}
	return nil

}
func (this *User) Current() error {
	if len(this.LoginToken) != 32 {
		return errors.New("Token 错误，请重新登录")
	}
	o := orm.NewOrm()
	if err := o.Read(this, "LoginToken"); err == orm.ErrNoRows {
		return errors.New("您已丢失登录信息或登出，请重新登录")
	} else {
		if !time.Now().Before(this.ExpiredAt) {
			this.Logout()
			return errors.New("用户登录已过期，请重新登录")
		} else {
			return nil
		}
	}
}
func (this *User) Logout() {
	if this.Id <= 0 {
		return
	}

	this.LoginToken = ""
	this.ExpiredAt = time.Now()
	this.Update("LoginToken", "ExpiredAt")
	loginLog := Log{
		Ip:     this.IP,
		User:   this,
		Status: LogoutSuccess,
	}
	loginLog.Create()
}
func (this *User) Delete() error {
	if this.Id <= 0 {
		return errors.New("ID错误,删除失败")
	}

	if _, err := orm.NewOrm().Delete(this); err == nil {
		return nil
	} else {
		return err
	}
}
