package dto

type Lottry struct {
	UserId  uint `json:"username" binding:"required"`
	PlayCnt int  `json:"start_cnt"binding:"required"`
}
