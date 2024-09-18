package request

type CreateViagensRequest struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	DataSaida   string  `json:"data_saida"`
	DataChegada string  `json:"data_chegada"`
	Valor       float64 `json:"valor"`
}

type UpdateViagensRequest struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	DataSaida   string  `json:"data_saida"`
	DataChegada string  `json:"data_chegada"`
	Valor       float64 `json:"valor"`
}
