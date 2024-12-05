package repositories

import "sekolah/models"

type PesertaDidikRepository interface {
	GetPesertaDidik([]models.PesertaDidikAPI, error)
}
