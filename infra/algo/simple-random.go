package algo

import (
	"math/rand"
	"time"
)

const min = int64(1)

// 简单随机红包算法
/**
参数为 红包数量，红包金额
设定最小单位为分，故1元 = 100分
返回值转换为元
 */
func SimpleRand(count, amount int64)  int64{
	// 当红包数量剩余1个时，不用计算，直接返回剩余金额
	if count==1{
		return amount
	}
	// 最大可调度金额
	max := amount - min * count
	rand.Seed(time.Now().UnixNano())
	x := rand.Int63n(max) + min
	return x
}
