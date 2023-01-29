package redis

import (
	"github.com/go-redis/redis"
	"math"
	"time"
)

const (
	OneWeekInSeconds         = 7 * 24 * 3600
	VoteScore        float64 = 432 // 每一票的值432分
	PostPerAge               = 20
)

func VoteForPost(userID string, postID string, v float64) (err error) {
	//1.判断投票限制
	//去redis取帖子发布时间
	postTime := client.ZScore(KeyPostTimeZSet, postID).Val()
	if float64(time.Now().Unix())-postTime > OneWeekInSeconds {
		//不允许投票了
		return ErrorVoteTimeExpire
	}
	//2.更新帖子的分数
	// 2和3 需要放到一个pipeline事务中操作
	// 判断是否已经投过票 查当前用户给当前帖子的投票记录

	key := KeyPostVotedZSetPrefix + postID
	ov := client.ZScore(key, userID).Val()

	// 更新：如果这一次投票的值和之前保存的值一致，就提示不允许重复投票
	if v == ov {
		return ErrVoteRepested
	}
	var op float64
	if v > ov {
		op = 1
	} else {
		op = -1
	}
	diffAbs := math.Abs(ov - v)                                                        // 计算两次投票的差值
	pipeline := client.TxPipeline()                                                    // 事务操作
	_, err = pipeline.ZIncrBy(KeyPostScoreZSet, VoteScore*diffAbs*op, postID).Result() // 更新分数
	if ErrorVoteTimeExpire != nil {
		return err
	}
	// 3、记录用户为该帖子投票的数据
	if v == 0 {
		_, err = client.ZRem(key, postID).Result()
	} else {
		pipeline.ZAdd(key, redis.Z{ // 记录已投票
			Score:  v, // 赞成票还是反对票
			Member: userID,
		})
	}

	//switch math.Abs(ov) - math.Abs(v) {
	//case 1:
	//	// 取消投票 ov=1/-1 v=0
	//	// 投票数-1
	//	pipeline.HIncrBy(KeyPostInfoHashPrefix+postID, "votes", -1)
	//case 0:
	//	// 反转投票 ov=-1/1 v=1/-1
	//	// 投票数不用更新
	//case -1:
	//	// 新增投票 ov=0 v=1/-1
	//	// 投票数+1
	//	pipeline.HIncrBy(KeyPostInfoHashPrefix+postID, "votes", 1)
	//default:
	//	// 已经投过票了
	//	return ErrorVoted
	//}
	_, err = pipeline.Exec()
	return err
}
