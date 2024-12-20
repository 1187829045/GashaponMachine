package model

type Prize struct {
	BaseModel
	UserId   uint     `gorm:"type:int(11)"`
	Mobile   string   `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	UserName string   `gorm:"type:varchar(20)"`
	Prizes   []string `gorm:"type:varchar(20)"`
}
