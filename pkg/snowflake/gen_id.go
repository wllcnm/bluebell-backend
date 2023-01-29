package snowflake

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

var (
	sonyFlake     *sonyflake.Sonyflake //实例
	sonyMachineID uint16               //机器ID
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

// Init 需传入当前的机器ID
func Init(machineId uint16) (err error) {
	sonyMachineID = machineId
	t, _ := time.Parse("2006-01-02", "2023-01-27") //初始化一个开始时间
	settings := sonyflake.Settings{                //生成全局配置
		StartTime: t,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings) //用配置生成sonyflake节点
	return
}
func GetID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("snoy flake not inited")
		return
	}
	id, err = sonyFlake.NextID()
	return
}
