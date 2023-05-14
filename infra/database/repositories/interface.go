package repositories

import "github.com/marcioecom/clipbot-server/infra/database/models"

type IUserRepository interface {
	Create(user *models.Users) error
	GetAll() ([]models.Users, error)
	GetByEmail(email string) (*models.Users, error)
}
