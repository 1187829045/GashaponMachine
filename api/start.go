package api

import (
	"GaMachine/form"
	"GaMachine/internal"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func start_lottery(c *gin.Context) {
	superPrize := "超级大奖"
	regularPrizes := []string{"衣服", "鞋子", "帽子"}
	thankYou := "谢谢惠顾"
	// 生成奖品列表
	var prizes []string
	cnt := form.StartCnt{}

	// 绑定请求中的数据到 user 结构体
	err := c.ShouldBind(&cnt)
	if err != nil {
		// 创建一个自定义错误，并返回客户端
		c.JSON(400, gin.H{
			"error": "绑定请求数据失败: " + err.Error(),
		})
		return // 直接返回，防止继续执行后续代码
	}
	fmt.Println(cnt.Username, cnt.Start_cnt)

	if isAdd, _ := internal.CheckNameInFile(internal.NameFile, cnt.Username); !isAdd {
		c.JSON(400, gin.H{
			"error": "用户不存在",
		})
	}

	diamond_count, err := internal.ReadDiamondCount(internal.BrickworkFile, cnt.Username)
	fmt.Println("姓名：%s,砖石数量:%d", cnt.Username, diamond_count)
	if diamond_count < cnt.Start_cnt*5 {
		c.JSON(400, gin.H{
			"error": "砖石数量不够",
		})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{
			"error": "读取砖石数量失败: " + err.Error(),
		})
	}

	for i := 1; i <= cnt.Start_cnt; i++ {
		if i%60 == 0 { // 每60次必得一个超级大奖
			prizes = append(prizes, superPrize)
		} else if i%10 == 0 { // 每10次奖品包括一次衣服，一次鞋子，一次帽子
			prizes = append(prizes, regularPrizes...)
		} else { // 其他为谢谢惠顾
			prizes = append(prizes, thankYou)
		}
	}

	// 打乱奖品列表
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(prizes), func(i, j int) {
		prizes[i], prizes[j] = prizes[j], prizes[i]
	})

	// 输出抽奖结果
	for i, prize := range prizes {
		fmt.Printf("第 %d 次抽奖: %s\n", i+1, prize)
	}
}
