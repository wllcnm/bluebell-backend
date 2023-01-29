package models

import (
	"encoding/json"
	"errors"
)

type LoginForm struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	UserID       uint64 `json:"user_id,string" db:"user_id"`
	UserName     string `json:"username" db:"username"`
	Password     string `json:"password" db:"password"`
	AccessToken  string
	RefreshToken string
}

type RegisterForm struct {
	UserName        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"` //密码确认,两次密码需要相同
}
type VoteDateForm struct {
	//UserID int 从请求中获取当前用户
	PostID    string `json:"post_id" binding:"required"`              //贴子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` //1:赞成票 -1:反对票 0:取消投票
}

// UnmarshalJSON ()为User结构体的自定义UnmarshalJSON方法
func (u *User) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName string `json:"username" db:"username"`
		Password string `json:"password" db:"password"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填的字段username")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填的字段password")
	} else {
		u.UserName = required.UserName //修改结构体中的username
		u.Password = required.Password //修改结构体中的password
	}

	return
}

// UnmarshalJSON ()为RegisterForm结构体的自定义UnmarshalJSON方法
func (r *RegisterForm) UnmarshalJSON(data []byte) (err error) {

	required := struct {
		UserName        string `json:"username"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}{}

	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填的username")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填的password")
	} else if !(required.Password == required.ConfirmPassword) {
		err = errors.New("两次密码不一致")
	} else {
		r.UserName = required.UserName
		r.Password = required.Password
		r.ConfirmPassword = required.ConfirmPassword
	}
	return

}
func (v *VoteDateForm) Unmarshal(data []byte) (err error) {
	required := struct {
		PostID    string `json:"post_id"`
		Direction int8   `json:"direction"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.PostID) == 0 {
		err = errors.New("缺少必填字段post_id")
	} else if required.Direction == 0 {
		err = errors.New("缺少必填字段direction")
	} else {
		v.PostID = required.PostID
		v.Direction = required.Direction
	}
	return
}
