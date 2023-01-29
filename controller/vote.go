package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VoteData struct {
	//从UserID string 从请求中获取当前的用户
	PostID    string `json:"post_id,string"`   //帖子id
	Direction int    `json:"direction,string"` //赞成票(1) 反对票(-1) 取消投票(0)
}

// UnmarshalJSON 为VoteData类型实现自定义的UnmarshalJSON方法
func (v *VoteData) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		PostID    string `json:"post_id"`
		Direction int    `json:"direction"`
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

/*
VoteHandler
@Description: 投票接口
@param c
*/

func VoteHandler(c *gin.Context) {
	//参数校验,给哪个文章投票
	vote := new(models.VoteDateForm)
	//进行参数校验
	err := c.ShouldBindJSON(&vote)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
			return
		}
		errdata := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParams, errdata)
		return
	}
	//获取当前请求用户的id
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNotLogin)
		return
	}
	//具体的投票业务
	err = logic.VoteForPost(userID, vote)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
