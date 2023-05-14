package repositories

import (
	"database/sql"

	. "github.com/go-jet/jet/v2/postgres"
	. "github.com/marcioecom/clipbot-server/.gen/clipbot/public/table"
	"github.com/marcioecom/clipbot-server/infra/database/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(user *models.Users) error {
	insertStmt := Users.INSERT(Users.Email, Users.Password).MODEL(user)

	if _, err := insertStmt.Exec(u.db); err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetAll() ([]models.Users, error) {
	var users []models.Users

	selectStmt := Users.SELECT(Users.ID, Users.Email, Users.CreatedAt, Users.UpdatedAt)

	if err := selectStmt.Query(u.db, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) GetByEmail(email string) (*models.Users, error) {
	var user models.Users

	selectStmt := Users.SELECT(
		Users.ID,
		Users.Email,
		Users.Password,
		Users.CreatedAt,
		Users.UpdatedAt,
	).WHERE(
		Users.Email.EQ(String(email)),
	)

	if err := selectStmt.Query(u.db, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
