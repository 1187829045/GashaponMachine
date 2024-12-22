package service

import (
	"GaMachine/consts"
	"GaMachine/internal/app/gachaSystem/dto"
	"GaMachine/internal/common"
	"GaMachine/model"
	"GaMachine/prize"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

type IGachaSystem interface {
	lottery(c *gin.Context) error
}

type GachaSystem struct{}

func NewGachaSystem() IGachaSystem {
	return &GachaSystem{}
}

func (ga *GachaSystem) lottery(c *gin.Context) error {

	lottry := dto.Lottry{}

	// 绑定请求中的数据到 user 结构体
	err := c.ShouldBindJSON(&lottry) //req dto service 内层
	if err != nil {
		// 创建一个自定义错误，并返回客户端
		err = errors.New(consts.ErrInvalidParameter)
		return err // 直接返回，防止继续执行后续代码
	}

	if _, err := model.GetUserById(lottry.UserId); err != nil {

		err = errors.New(consts.UserNotFound)
		return err
	}

	//从文件中读取砖石数量
	diamond_count := model.DiamondCount(lottry.UserId)

	if diamond_count < lottry.PlayCnt*5 {
		err = errors.New("砖石数量不足")
		return err
	}
	fmt.Println("开始抽奖......")
	p := StartLottery(lottry.PlayCnt)

	//添加到数据库

	err = model.AddPrize(lottry.UserId, p)
	if err != nil {
		return err
	}
	return nil

}

func StartLottery(PlayCnt int) []string {
	prizes := []string{}
	for i := 1; i <= PlayCnt; i++ {
		if i%60 == 0 { // 每60次必得一个超级大奖
			prizes = append(prizes, prize.SuperPrize)
		} else {
			// 初始化随机数生成器的种子，一般使用当前时间的纳秒时间戳
			rand.Seed(time.Now().UnixNano())
			// 生成一个在 [0, 100) 区间内的整数随机数
			randomInt := rand.Intn(100)
			index := randomInt % len(prize.Prize_pool)
			prizes = append(prizes, prize.Prize_pool[index])
			common.RemoveIndexPrize(index)
			//fmt.Println("当前奖品池:", prize.Prize_pool)
		}

		if i%10 == 0 && i != 0 {
			common.Reinitialization()
		}
	}
	return prizes
}
