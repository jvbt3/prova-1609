package model

type Viagens struct {
	Id          int     `gorm:"type:int;primary_key"`
	Name        string  `gorm:"type:varchar(255)"`
	DataSaida   string  `gorm:"type:varchar(255)"`
	DataChegada string  `gorm:"type:varchar(255)"`
	Valor       float64 `gorm:"type:float"`
}
