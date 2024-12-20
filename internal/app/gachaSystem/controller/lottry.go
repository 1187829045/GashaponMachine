package controller

import (
	"GaMachine/form"
	"GaMachine/internal/common"
	"GaMachine/prize"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

type IGame interface {
	Lottery(c *gin.Context) error
}

type Game struct {
}

func NewGame() IGame {
	return &Game{}
}
func (*Game) Lottery(c *gin.Context) error {
	// 生成奖品列表
	var prizes []string
	cnt := form.StartCnt{} //req dto service 内层

	// 绑定请求中的数据到 user 结构体
	err := c.ShouldBindJSON(&cnt)
	if err != nil {
		// 创建一个自定义错误，并返回客户端
		c.JSON(400, gin.H{
			"error": "绑定请求数据失败: " + err.Error(),
		})
		return // 直接返回，防止继续执行后续代码
	}

	if isAdd, _ := common.CheckNameInFile(common.NameFile, cnt.Username); !isAdd {
		c.JSON(400, gin.H{
			"error": "用户不存在",
		})
		return
	}

	//从文件中读取砖石数量
	diamond_count, err := common.ReadDiamondCount(common.BrickworkFile, cnt.Username)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "读取砖石数量失败: " + err.Error(),
		})
	}

	fmt.Printf("姓名%s,想玩%d次，拥有砖石%d", cnt.Username, cnt.Start_cnt, diamond_count)

	if diamond_count < cnt.Start_cnt*5 {
		c.JSON(400, gin.H{
			"error": "砖石数量不够",
		})
		return
	}
	fmt.Println("开始抽奖......")

	prizes = Start_lottery(cnt.Start_cnt, prizes)
	diamond_count -= cnt.Start_cnt * 5

	err = common.ModifyDiamondCount(common.BrickworkFile, cnt.Username, diamond_count)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "修改砖石数量失败: " + err.Error(),
		})
	}

	err = common.WritePrizeToFile(common.GiftFile, cnt.Username, prizes)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "写入奖品失败: " + err.Error(),
		})
	}

	fmt.Println("获得的奖品:%d", prizes)

	c.JSON(200, gin.H{
		"prize": "抽奖完成",
	})
}
func fallback(err error) error {
	fmt.Println("进入降级逻辑，原因:", err)
	return fmt.Errorf("获取用户信息失败，已触发降级逻辑")
}

func Start_lottery(lottery_cnt int, prizes []string) []string {
	hystrix.Do("lottery", func() error {
		for i := 1; i <= lottery_cnt; i++ {
			fmt.Println(prizes)
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
		return nil
	}, func(err error) error {
		return fallback(err)
	})
	return prizes
}
