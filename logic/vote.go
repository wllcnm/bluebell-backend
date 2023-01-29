package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"go.uber.org/zap"
	"strconv"
)

/*
VoteForPost
@Description: 投票功能业务层
*/
func VoteForPost(userId uint64, p *models.VoteDateForm) error {
	zap.L().Debug("VoteForPost",
		zap.Uint64("userId", userId),
		zap.String("postId", p.PostID),
		zap.Int8("Direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userId)), p.PostID, float64(p.Direction))
}
