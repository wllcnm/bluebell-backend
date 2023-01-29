package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(post *models.Post) (err error) {
	//1.生成post_id(生成帖子id)
	postID, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("snowflake.GetID() failed ", zap.Error(err))
		return
	}
	post.PostID = postID

	//2.创建贴子 保存到数据库
	err = mysql.CreatePost(post)
	if err != nil {
		zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
		return
	}
	return
}
