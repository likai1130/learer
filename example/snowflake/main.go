/**请使用go语言帮我实现一个雪花算法，谢谢.*/

package main

import (
	"fmt"
	"time"
)

// Snowflake 算法实现
type Snowflake struct {
	startTime         int64 // 开始时间戳，毫秒级
	workerIdBits      uint  // 机器id所占的位数
	dataCenterIdBits  uint  // 数据中心id所占的位数
	sequenceBits      uint  // 序列号所占的位数
	maxWorkerId       int64 // 最大机器id
	maxDataCenterId   int64 // 最大数据中心id
	workerIdShift     uint  // 机器id左移位数
	dataCenterIdShift uint  // 数据中心id左移位数
	timestampShift    uint  // 时间戳左移位数
	sequenceMask      int64 // 序列号掩码
	workerId          int64 // 机器id
	dataCenterId      int64 // 数据中心id
	sequence          int64 // 序列号
	lastTimestamp     int64 // 上一次时间戳
}

// NewSnowflake 初始化Snowflake算法
func NewSnowflake(startTime int64, workerIdBits uint, dataCenterIdBits uint, sequenceBits uint) *Snowflake {
	// 计算最大机器id和数据中心id
	maxWorkerId := int64(-1) ^ (int64(-1) << workerIdBits)
	maxDataCenterId := int64(-1) ^ (int64(-1) << dataCenterIdBits)

	// 计算机器id和数据中心id左移位数
	workerIdShift := sequenceBits
	dataCenterIdShift := sequenceBits + workerIdBits

	// 计算时间戳左移位数
	timestampShift := sequenceBits + workerIdBits + dataCenterIdBits

	// 计算序列号掩码
	sequenceMask := int64(-1) ^ (int64(-1) << sequenceBits)

	return &Snowflake{
		startTime:         startTime,
		workerIdBits:      workerIdBits,
		dataCenterIdBits:  dataCenterIdBits,
		sequenceBits:      sequenceBits,
		maxWorkerId:       maxWorkerId,
		maxDataCenterId:   maxDataCenterId,
		workerIdShift:     workerIdShift,
		dataCenterIdShift: dataCenterIdShift,
		timestampShift:    timestampShift,
		sequenceMask:      sequenceMask,
	}
}

// SetWorkerId 设置机器id
func (sf *Snowflake) SetWorkerId(workerId int64) error {
	if workerId < 0 || workerId > sf.maxWorkerId {
		return fmt.Errorf("worker Id can't be greater than %d or less than 0", sf.maxWorkerId)
	}
	sf.workerId = workerId
	return nil
}

// SetDataCenterId 设置数据中心id
func (sf *Snowflake) SetDataCenterId(dataCenterId int64) error {
	if dataCenterId < 0 || dataCenterId > sf.maxDataCenterId {
		return fmt.Errorf("data center Id can't be greater than %d or less than 0", sf.maxDataCenterId)
	}
	sf.dataCenterId = dataCenterId
	return nil
}

// Generate 生成id
func (sf *Snowflake) Generate() (int64, error) {
	// 获取当前时间戳
	timestamp := time.Now().UnixNano() / 1e6

	// 如果当前时间戳小于上一次时间戳，则表示时间戳获取出现异常
	if timestamp < sf.lastTimestamp {
		return 0, fmt.Errorf("invalid system clock")
	}

	// 如果当前时间戳等于上一次时间戳，则需要生成序列号
	if timestamp == sf.lastTimestamp {
		sf.sequence = (sf.sequence + 1) & sf.sequenceMask
		// 如果序列号溢出，则需要重新获取时间戳
		if sf.sequence == 0 {
			for timestamp <= sf.lastTimestamp {
				timestamp = time.Now().UnixNano() / 1e6
			}
		}
	} else { // 如果当前时间戳大于上一次时间戳，则需要重置序列号
		sf.sequence = 0
	}

	// 更新上一次时间戳
	sf.lastTimestamp = timestamp

	// 生成id
	id := ((timestamp - sf.startTime) << sf.timestampShift) | (sf.dataCenterId << sf.dataCenterIdShift) | (sf.workerId << sf.workerIdShift) | sf.sequence

	return id, nil
}

func main() {
	// 初始化Snowflake算法
	sf := NewSnowflake(1546272000000, 10, 10, 12)

	// 设置机器id和数据中心id
	sf.SetWorkerId(1)
	sf.SetDataCenterId(1)

	// 生成id
	for i := 0; i < 10; i++ {
		id, err := sf.Generate()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(id)
		}
	}

}
