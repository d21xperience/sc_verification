package repositories

import "sekolah/models"

type RombelRepository interface {
	GetRombel([]models.RombelApi, error)
}
