package models

import "github.com/astaxie/beego/client/orm"

var (
	UserList map[string]*User
)

type User struct {
	Id int64
	DeviceId string
	Password string
	NikeName  string
	HeadImg string
	Summary string
	RegFrom string
	Unionid string
	RegTime string
	LoginTime int64
	UserContents []*UserContent `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "ob_user"
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
