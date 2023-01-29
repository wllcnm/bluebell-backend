package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	//1.获取参数及校验参数
	var u *models.Post
	if err := c.BindJSON(&u); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}

	//2.获取用户id
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	u.AuthorId = userID
	//2.执行业务逻辑
	err = logic.CreatePost(u)
	if err != nil {
		zap.L().Error("logic.CreatePost(u) failed ", zap.Error(err))
		ResponseErrorWithMsg(c, CodeCreatePostFailed, err.Error())
		return
	}
	ResponseSuccess(c, u)
}
