/*
* @Author: GuoDi
* @Date:   2016-11-23 22:28:55
* @Last Modified by:   GuoDi
* @Last Modified time: 2016-11-23 23:19:36
 */

package account

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ykrl089/beelib/library/str"
	"testing"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/account?charset=utf8")
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
	EnableLog = true
}

func Test_Create_User(t *testing.T) {
	fmt.Println("-------", "Test_Create_User", "------")
	usr := User{Username: str.RanStr(8), PasswordPlain: str.RanStr(12)}
	if err := usr.Create(); err != nil {
		fmt.Println(err)
	}
	usr = User{Username: "guodi", PasswordPlain: "aaaaaaa"}
	usr.Create()
	fmt.Println(usr)
	t.Log(usr)
}
func Test_Login_User(t *testing.T) {
	fmt.Println("-------", "Test_Login_User", "------")
	usr := User{Username: "guodi", PasswordPlain: "aaaaaaa"}
	if err := usr.Login("192.168.1.8"); err != nil {
		fmt.Println(err)
	}
	fmt.Println(usr)
	t.Log(usr)
}
func Test_Get_Login_User(t *testing.T) {
	fmt.Println("-------", "Test_Get_Login_User", "------")
	usr := User{Username: "guodi", PasswordPlain: "aaaaaaa"}
	if err := usr.Login("192.168.1.8"); err != nil {
		fmt.Println(err)
	}
	fmt.Println(usr)
	if usr.LoginToken != "" {
		usrn := User{LoginToken: usr.LoginToken}
		if err := usrn.Current(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("usrnnnn:", usrn)
		}
	} else {
		fmt.Println("login error")
	}
	t.Log(usr)
}
