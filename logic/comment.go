package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"errors"
)

func CreateComment(comment *models.Comment) error {
	return mysql.CreateComment(comment)
}
func GetCommentListByID(id int64) (data []*models.CommentDetail, err error) {
	comments, err := mysql.GetCommentListByID(id)
	if err != nil {
		err = errors.New("查询评论出错")
		return
	}
	data = make([]*models.CommentDetail, 0, len(comments))

	for _, comment := range comments {
		user, err := mysql.GetUserByID(comment.AuthorID)
		if err != nil {
			continue
		}

		childCommentLists, err := mysql.GetChildComment(int64(comment.CommentID))
		var x []*models.Children
		for _, list := range childCommentLists {
			m := &models.Children{
				AuthorName: user.UserName,
				Comment:    list,
			}
			x = append(x, m)
		}
		m1 := &models.CommentDetail{
			Comment:    comment,
			AuthorName: user.UserName,
			Children:   x,
		}
		if comment.ParentID == 0 {
			data = append(data, m1)
		}
	}
	return
}
