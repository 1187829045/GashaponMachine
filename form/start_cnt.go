package form

type StartCnt struct {
	Username  string `json:"username" binding:"required"`
	Start_cnt int    `json:"start_cnt"binding:"required"`
}
