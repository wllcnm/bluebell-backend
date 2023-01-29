package mysql

import (
	"bluebell/models"
	"database/sql"
	"go.uber.org/zap"
)

/**
 * @Author huchao
 * @Description //TODO 根据ID查询分类社区详情
 * @Date 17:08 2022/2/12
 **/

func GetCommunityByID(id uint64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlstr := "select community_id, community_name, introduction, create_time from community where community_id=?"
	err = db.Get(community, sqlstr, id)
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query community failed", zap.String("sql", sqlstr), zap.Error(err))
		err = ErrorQueryFailed
	}
	return
}

/*
GetCommunityList
@Description: dao层 获取所有社区列表
*/
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlstr := "select community_id,community_name from community"
	err = db.Select(&communityList, sqlstr) //切片类型需传入指针地址
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no community in db", zap.String("sql", sqlstr))
		return
	}
	return
}
func GetCommunityById(id uint64) (communityList *models.CommunityDetail, err error) {
	communityList = new(models.CommunityDetail)
	sqlstr := "select community_id,community_name,introduction,create_time from community where community_id=?"
	err = db.Get(communityList, sqlstr, id)
	//查询不到社区
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
		return
	}
	//查询社区出现错误
	if err != nil {
		zap.L().Error("query community by id failed ", zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	return
}
