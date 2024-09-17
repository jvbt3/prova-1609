package service

import (
	"prova-esoft8s/data/request"
	"prova-esoft8s/data/response"
)

type ViagensService interface {
	Create(viagem request.CreateViagensRequest)
	Update(viagem request.UpdateViagensRequest)
	Delete(viagemId int)
	FindById(viagemId int) response.ViagensResponse
	FindAll() []response.ViagensResponse
}
