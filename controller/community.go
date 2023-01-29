package controller

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

/*
CommunityHandler
@Description: 查询所有社区,并以(community_id,community_name) 形式返回
@param c
@author awei
*/
func CommunityHandler(c *gin.Context) {
	communityList, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) //返回服务繁忙是避免把服务端错误暴露给外部
		return
	}
	ResponseSuccess(c, communityList)
	return
}

/*
CommunityDetailHandler
@Description: 获取社区详细信息
@param c
*/
func CommunityDetailHandler(c *gin.Context) {
	//1.从url中获取参数
	communityIdStr := c.Param("id")
	communityId, err := strconv.ParseUint(communityIdStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}

	//2.根据id获取社区详细信息
	communityDetail, err := logic.GetCommunityListById(communityId)
	if err != nil {
		zap.L().Error("logic.GetCommunityListById(communityId) failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}

	ResponseSuccess(c, communityDetail)
	return
}
