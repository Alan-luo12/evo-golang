//  注意这里实现的是黑盒，后面要自己掌握

package snowid

import (
	"time"

	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

// Init 初始化雪花ID生成器，项目启动时调用一次
func Init(machineID uint16) {
	st, err := time.Parse("2006-01-02", "2020-01-01")
	if err != nil {
		panic(err)
	}

	sf = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: st,
		MachineID: func() (uint16, error) {
			return machineID, nil // 单机部署直接写死1，分布式环境传入不同的机器ID
		},
	})
}

// NextID 生成一个雪花ID
func NextID() (uint64, error) {
	if sf == nil {
		panic("snowflake not initialized")
	}
	return sf.NextID()
}
