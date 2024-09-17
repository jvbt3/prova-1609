package model

import "time"

type Viagens struct {
	Id          int       `gorm:"type:int;primary_key"`
	Name        string    `gorm:"type:varchar(255)"`
	DataSaida   time.Time `gorm:"type:date"`
	DataChegada time.Time `gorm:"type:date"`
	Valor       float64   `gorm:"type:float"`
}
