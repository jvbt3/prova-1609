package response

type ViagensResponse struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	DataSaida   string  `json:"data_saida"`
	DataChegada string  `json:"data_chegada"`
	Valor       float64 `json:"valor"`
}

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
