package models

import "github.com/astaxie/beego/client/orm"

var (
	UserList map[string]*User
)

type User struct {
	Id int64 `json:"uid"`
	DeviceId string `json:"device_id"`
	Password string `json:"password"`
	NikeName  string `json:"nike_name"`
	HeadImg string `json:"head_img"`
	Summary string `json:"summary"`
	RegFrom string `json:"reg_from"`
	OauthUid string `json:"oauth_uid"`
	RegTime int64 `json:"reg_time"`
	LoginTime int64 `json:"login_time"`
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
