package service

import (
	"github.com/go-playground/validator/v10"
	"prova-esoft8s/data/request"
	"prova-esoft8s/data/response"
	"prova-esoft8s/helper"
	"prova-esoft8s/model"
	"prova-esoft8s/repository"
)

type ViagensServiceImpl struct {
	ViagemRepository repository.ViagensRepository
	Validate         *validator.Validate
}

func NewViagemServiceImpl(viagemRepository repository.ViagensRepository, validate *validator.Validate) ViagensService {
	return &ViagensServiceImpl{
		ViagemRepository: viagemRepository,
		Validate:         validate,
	}
}

func (t ViagensServiceImpl) Create(viagem request.CreateViagensRequest) {
	err := t.Validate.Struct(viagem)
	helper.ErrorPanic(err)
	viagemModel := model.Viagens{
		Id:          viagem.Id,
		Name:        viagem.Name,
		DataSaida:   viagem.DataSaida,
		DataChegada: viagem.DataChegada,
		Valor:       viagem.Valor,
	}
	t.ViagemRepository.Save(viagemModel)
}

func (t ViagensServiceImpl) Update(viagem request.UpdateViagensRequest) {
	viagemData, err := t.ViagemRepository.FindById(viagem.Id)
	helper.ErrorPanic(err)
	viagemData.Name = viagem.Name
	viagemData.DataSaida = viagem.DataSaida
	viagemData.DataChegada = viagem.DataChegada
	viagemData.Valor = viagem.Valor

	t.ViagemRepository.Update(viagemData)
}

func (t ViagensServiceImpl) Delete(viagemId int) {
	t.ViagemRepository.Delete(viagemId)
}

func (t ViagensServiceImpl) FindById(viagemId int) response.ViagensResponse {
	viagemData, err := t.ViagemRepository.FindById(viagemId)
	helper.ErrorPanic(err)

	viagemResponse := response.ViagensResponse{
		Id:          viagemData.Id,
		Name:        viagemData.Name,
		DataSaida:   viagemData.DataSaida,
		DataChegada: viagemData.DataChegada,
		Valor:       viagemData.Valor,
	}
	return viagemResponse
}

func (t ViagensServiceImpl) FindAll() []response.ViagensResponse {
	result := t.ViagemRepository.FindAll()

	var viagems []response.ViagensResponse
	for _, value := range result {
		viagem := response.ViagensResponse{
			Id:          value.Id,
			Name:        value.Name,
			DataSaida:   value.DataSaida,
			DataChegada: value.DataChegada,
			Valor:       value.Valor,
		}
		viagems = append(viagems, viagem)
	}
	return viagems
}
