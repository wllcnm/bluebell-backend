package mysql

import (
	"bluebell/models"
	"go.uber.org/zap"
)

/**
 * @Author huchao
 * @Description //TODO 创建帖子
 * @Date 19:53 2022/2/12
 **/

func CreatePost(post *models.Post) (err error) {
	sqlstr := "insert into post(post_id,title,content,author_id,community_id)" +
		"values(?,?,?,?,?)"
	_, err = db.Exec(sqlstr, post.PostID, post.Title, post.Content, post.AuthorId, post.CommunityID)
	if err != nil {
		zap.L().Error("insert post failed ", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}
