package algo

import "math/rand"

// 后洗牌算法

func AfterShuffle(count, amount int64) []int64{
	inds := make([]int64,0)
	// 1.计算最大可调度金额
	//max := amount - min * count
	remain := amount
	// 2.随机生成初级红包序列
	for i:=int64(0); i<count; i++ {
		x:=SimpleRand(count-i,remain)
		remain -= x
		inds=append(inds,x)
	}
	// 3.洗牌
	rand.Shuffle(len(inds), func(i, j int) {
		inds[i], inds[j] = inds[j], inds[i]
	})
	return inds
}