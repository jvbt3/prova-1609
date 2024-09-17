package request

import "time"

type CreateViagensRequest struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	DataSaida   time.Time `json:"data_saida"`
	DataChegada time.Time `json:"data_chegada"`
	Valor       float64   `json:"valor"`
}

type UpdateViagensRequest struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	DataSaida   time.Time `json:"data_saida"`
	DataChegada time.Time `json:"data_chegada"`
	Valor       float64   `json:"valor"`
}
