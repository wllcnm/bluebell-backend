package models

import (
	"time"
)

type Comment struct {
	PostID     uint64    `json:"question_id" db:"post_id"`
	ParentID   uint64    `json:"parent_id" db:"parent_id"`
	CommentID  uint64    `json:"comment_id" db:"comment_id"`
	AuthorID   uint64    `json:"author_id" db:"author_id"`
	Content    string    `json:"content" db:"content"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}
type CommentDetail struct {
	*Comment
	AuthorName string
	Children   []*Children
}
type Children struct {
	AuthorName string
	*Comment
}
