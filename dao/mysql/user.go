package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

// 把每一步数据库操作封装成函数
// 待logic层根据业务需求调用

const secret = "jojo.lw123.top"

/**
 * @Author huchao
 * @Description //TODO 对密码进行加密
 * @Date 21:50 2022/2/10
 **/
func encryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(data))
}

/**
 * @Author huchao
 * @Description //TODO 检查指定用户名的用户是否存在
 * @Date 21:50 2022/2/10
 **/

func CheckUserExist(username string) (err error) {
	sqlstr := "select count(user_id) from user where username=?"
	var count int
	if err := db.Get(&count, sqlstr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExit
	}
	return
}

/**
 * @Author hucaho
 * @Description //TODO 注册业务-向数据库中插入一条新的用户
 * @Date 21:51 2022/2/10
 **/

func InsertUser(user models.User) (error error) {
	//对密码进行加密
	user.Password = encryptPassword([]byte(user.Password)) //将string类型转换为[]byte类型
	//执行sql语句入库
	sqlstr := "insert into user(user_id,username,password) values(?,?,?)"

	_, err := db.Exec(sqlstr, user.UserID, user.UserName, user.Password)
	return err
}

func Login(user *models.User) (err error) {
	orginPassword := user.Password //记录用户的登录密码
	sqlstr := "select user_id,username,password from user where username=?"
	err = db.Get(user, sqlstr, user.UserName)
	if err != nil && err != sql.ErrNoRows {
		//查询数据库出错
		return ErrorQueryFailed
	}
	if err == sql.ErrNoRows {
		//用户不存在
		return ErrorUserNotExit
	}
	password := encryptPassword([]byte(orginPassword))
	if password != user.Password {
		return ErrorPasswordWrong
	}
	return
}
