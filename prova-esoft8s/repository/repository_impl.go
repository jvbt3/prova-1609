package repository

import (
	"errors"
	"gorm.io/gorm"
	"prova-esoft8s/data/request"
	"prova-esoft8s/helper"
	"prova-esoft8s/model"
)

type ViagensRepositoryImpl struct {
	Db *gorm.DB
}

func NewViagensRepositoryImpl(Db *gorm.DB) ViagensRepository {
	return &ViagensRepositoryImpl{Db: Db}
}

func (t ViagensRepositoryImpl) Save(viagens model.Viagens) {
	result := t.Db.Save(&viagens)
	helper.ErrorPanic(result.Error)

}

func (t ViagensRepositoryImpl) Update(viagens model.Viagens) {
	var updateTag = request.UpdateViagensRequest{
		Id:          viagens.Id,
		Name:        viagens.Name,
		DataSaida:   viagens.DataSaida,
		DataChegada: viagens.DataChegada,
		Valor:       viagens.Valor,
	}
	result := t.Db.Model(&viagens).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}

func (t ViagensRepositoryImpl) Delete(viagensId int) {
	var viagens model.Viagens
	result := t.Db.Where("id = ?", viagensId).Delete(&viagens)
	helper.ErrorPanic(result.Error)
}

func (t ViagensRepositoryImpl) FindById(viagensId int) (model.Viagens, error) {
	var tag model.Viagens
	result := t.Db.Find(&tag, viagensId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

func (t ViagensRepositoryImpl) FindAll() []model.Viagens {
	var viagens []model.Viagens
	results := t.Db.Find(&viagens)
	helper.ErrorPanic(results.Error)
	return viagens
}
