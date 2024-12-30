package service

import (
	"GaMachine/consts"
	"GaMachine/internal/app/gachaSystem/dto"
	"GaMachine/internal/common"
	"GaMachine/model"
	"GaMachine/prize"
	"errors"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

type IGachaSystem interface {
	Lottery(c *gin.Context, req dto.Lottry)
	GetPrize(c *gin.Context)
}

type GachaSystem struct{}

func NewGachaSystem() IGachaSystem {
	return &GachaSystem{}
}

func (ga *GachaSystem) Lottery(c *gin.Context, req dto.Lottry) {

	if _, err := model.GetUserById(req.UserId); err != nil {
		err = errors.New(consts.UserNotFound)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//从文件中读取砖石数量
	diamond_count := model.DiamondCount(req.UserId)

	if diamond_count < req.PlayCnt*5 {
		err := errors.New("砖石数量不足")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	p := StartLottery(req.PlayCnt)

	//添加到数据库

	err := model.AddPrize(req.UserId, p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": consts.ErrDB,
		})
		return
	}
	err = model.UpdateUser(req.UserId, model.User{
		DiamondCount: diamond_count - req.PlayCnt*5,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": consts.ErrDB,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"prize": p,
	})
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
func (ga *GachaSystem) GetPrize(c *gin.Context) {
	uId, exists := c.Get("userId")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": consts.ErrNoLogin,
		})
		return
	}
	prizes, err := model.GetPrize(uId.(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": consts.ErrDB,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"prize": prizes,
	})
}
