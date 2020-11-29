package models

import "github.com/astaxie/beego/client/orm"

var (
	UserList map[string]*User
)

type User struct {
	Id int64
	Uid int64
	Username string
	Password string
	NikeName  string
	RegFrom string
	RegTime string
	LoginTime int64
	HeadImg string
	Summary string
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