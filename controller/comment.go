package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

/*
CommentHandler
@Description: 创建评论
@param c
*/
func CommentHandler(c *gin.Context) {
	var comment models.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	//生成评论id
	commentID, err := snowflake.GetID()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	//创建时间
	comment.CreateTime = time.Now()
	//获取作者id
	userId, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNotLogin)
		return
	}
	comment.CommentID = commentID
	comment.AuthorID = userId

	//创建评论
	err = logic.CreateComment(&comment)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

/*
CommentListHandler
@Description: 获取评论列表
@param c
*/
func CommentListHandler(c *gin.Context) {
	id, ok := c.GetQuery("id")
	i, _ := strconv.ParseInt(id, 10, 64)
	if !ok {
		ResponseError(c, CodeInvalidParams)
		return
	}
	list, err := logic.GetCommentListByID(i)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, list)
	return
}
