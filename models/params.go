package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

/**
 * @Author huchao
 * @Description //TODO 获取帖子列表query string参数
 * @Date 21:51 2023/1/31
 **/

type ParamPostList struct {
	CommunityID uint64 `json:"community_id" form:"community_id"`   // 可以为空
	Page        int64  `json:"page" form:"page"`                   // 页码
	Size        int64  `json:"size" form:"size"`                   // 每页数量
	Order       string `json:"order" form:"order" example:"score"` // 排序依据
}
