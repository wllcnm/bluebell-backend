package mysql

import (
	"bluebell/models"
	"database/sql"
	"go.uber.org/zap"
)

func CreateComment(comment *models.Comment) (err error) {
	sqlstr := "insert into comment(comment_id,content,post_id,author_id,parent_id) values (?,?,?,?,?)"
	_, err = db.Exec(sqlstr, comment.CommentID, comment.Content, comment.PostID, comment.AuthorID, comment.ParentID)
	if err != nil {
		err = ErrorInsertFailed
		return
	}
	return
}
func GetCommentListByID(id int64) (commentList []*models.Comment, err error) {
	sqlstr := "select post_id,comment_id,content,author_id,create_time,parent_id from comment where post_id=?"
	err = db.Select(&commentList, sqlstr, id)
	if err != nil {
		zap.L().Error("mysql.GetCommentListByID(id string) failed", zap.Error(err))
		return nil, err
	}
	return
}

func GetChildComment(postID int64) (commentList []*models.Comment, err error) {
	sqlstr := "select post_id,comment_id,content,author_id,create_time,parent_id from comment where parent_id=?"
	err = db.Select(&commentList, sqlstr, postID)
	if err == sql.ErrNoRows {
		return
	}
	if err != nil {
		zap.L().Error("mysql.GetChildComment(id string) failed", zap.Error(err))
		return
	}
	return
}
