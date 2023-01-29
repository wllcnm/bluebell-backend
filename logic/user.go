package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	"errors"
)

func SignUp(p *models.RegisterForm) (err error) {
	//1.判断用户存不存在
	err = mysql.CheckUserExist(p.UserName)
	if err != nil {
		//存在同名用户
		if errors.Is(err, mysql.ErrorUserExit) {
			return mysql.ErrorUserExit
		}
		//数据库查询出错
		return err
	}

	//2.生成UID
	userId, err := snowflake.GetID()
	if err != nil {
		return mysql.ErrorGenIDFailed
	}
	//3.构造一个User实例
	u := models.User{
		UserID:   userId,
		UserName: p.UserName,
		Password: p.Password,
	}
	//4.保存进数据库
	return mysql.InsertUser(u)
}
func Login(p *models.LoginForm) (user *models.User, err error) {
	user = &models.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	//如果登录出错
	if err = mysql.Login(user); err != nil {
		return nil, err
	}
	//生成jwt
	token, rToken, err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		return nil, err
	}
	user.AccessToken = token
	user.RefreshToken = rToken
	return
}
