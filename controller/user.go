package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func LoginHandler(c *gin.Context) {
	//1.获取请求参数及参数校验
	var u *models.LoginForm

	if err := c.ShouldBindJSON(&u); err != nil {
		//请求参数有误,直接返回响应
		zap.L().Error("SiginUp with invalid param", zap.Error(err))
		//判断err是不是validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//非validator.ValidationErrors类型错误直接返回
			ResponseError(c, CodeInvalidParams) //请求参数错误
			return
		}
		//validator.ValidationErrors错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		return
	}

	//2.业务逻辑处理
	user, err := logic.Login(u)
	if err != nil { //如果出错
		zap.L().Error("logic.Login failed", zap.String("username", u.UserName), zap.Error(err))
		if errors.Is(err, mysql.ErrorQueryFailed) {
			ResponseError(c, CodeServerBusy)
			return
		}
		if errors.Is(err, mysql.ErrorUserNotExit) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		if errors.Is(err, mysql.ErrorPasswordWrong) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseError(c, CodeInvalidParams)
		return
	}
	//3.如果没出错,则返回响应
	ResponseSuccess(c, gin.H{
		"user_id":       user.UserID,
		"user_name":     user.UserName,
		"access_token":  user.AccessToken,
		"refresh_token": user.RefreshToken,
	})

}
func SignUpHandler(c *gin.Context) {
	//1.获取参数
	var u *models.RegisterForm

	//如果参数出错
	if err := c.ShouldBindJSON(&u); err != nil {
		//请求参数有误,直接返回响应
		zap.L().Error("SiginUp with invalid param", zap.Error(err))
		//判断err是不是validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//非validator.ValidationErrors类型错误直接返回
			ResponseErrorWithMsg(c, CodeInvalidParams, err.Error()) //请求参数错误
			return
		}
		//validator.ValidationErrors错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		return
	}

	//2.注册业务逻辑处理
	if err := logic.SignUp(u); err != nil {
		zap.L().Error("logic.signup failed ,err is", zap.Error(err))

		//如果该用户已经被注册
		if errors.Is(err, mysql.ErrorUserExit) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, u)
}

/**
 * @Author huchao
 * @Description //TODO 刷新accessToken
 * @Date 17:09 2022/2/17
 **/

func RefreshTokenHandler(c *gin.Context) {
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// 这里假设Token放在Header的Authorization中，并使用Bearer开头
	// 这里的具体实现方式要依据你的实际业务情况决定
	rt := c.Query("refresh_token") //从请求头中获取token
	AuthHeader := c.GetHeader("Authorization")
	if AuthHeader == "" {
		ResponseErrorWithMsg(c, CodeInvalidToken, "请求头缺少Auth Token")
		c.Abort() //直接阻止程序
		return
	}
	//按空格分割
	parts := strings.SplitN(AuthHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ResponseErrorWithMsg(c, CodeInvalidToken, "token格式不对")
		c.Abort()
		return
	}
	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
	fmt.Println(err)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}
