package repository

import (
	"prova-esoft8s/model"
)

type ViagensRepository interface {
	Save(viagens model.Viagens)
	Update(viagens model.Viagens)
	Delete(viagensId int)
	FindById(viagensId int) (viagens model.Viagens, err error)
	FindAll() []model.Viagens
}
