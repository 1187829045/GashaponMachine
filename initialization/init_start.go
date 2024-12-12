package initialization

import "GaMachine/prize"

func InitPrize() {
	prize.Prize_pool = []string{"谢谢惠顾", "谢谢惠顾", "谢谢惠顾", "谢谢惠顾", "谢谢惠顾", "A级奖品", "B级奖品", "C级奖品", "D级奖品", "E级奖品"}

	//prize.Prize_cnt = make(map[string]int)

	//for _, p := range prize.Prize_pool {
	//
	//	prize.Prize_cnt[p]++
	//}
}
