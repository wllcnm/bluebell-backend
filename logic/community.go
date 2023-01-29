package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}

func GetCommunityListById(id uint64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityById(id)
}
