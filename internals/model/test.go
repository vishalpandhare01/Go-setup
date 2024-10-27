package model

type Programing struct {
	ID   string `gprm:"type:char(36)"`
	Name string `gorm:"type:varchar(255)"`
	Use  string `gorm:"type:varchar(10)"`
}
