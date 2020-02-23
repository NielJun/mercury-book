package generateid

import (
	"errors"
	"github.com/sony/sonyflake"
)

var (
	// 全局的sony id 生成器的对象 一切生成都是这个负责
	sonyFlake     *sonyflake.Sonyflake
	// 外部传入的machineid 机器标识 但是sonyflake 需要通过回调方式赋值 所以存起来
	sonyMachineId uint16
)

// 外部调用的该初始化函数，这个模块是用来进行唯一Id的生成的
func Init(machineId uint16) (err error) {
	sonyMachineId = machineId
	setting := sonyflake.Settings{}
	setting.MachineID = getMachineId // 回调函数 在初始化以后进行回调
	sonyFlake = sonyflake.NewSonyflake(setting)
	return
}

func getMachineId() (uint16, error) {
	return sonyMachineId, nil
}

func GetId() (id uint64, err error) {
	if sonyFlake == nil {
		err = errors.New("sony flake not inited ")
		return
	}

	id, err = sonyFlake.NextID()
	return

}
