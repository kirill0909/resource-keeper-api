package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kirill0909/resource-keeper-api/models"
)

type UserResourcePostgres struct {
	db *sqlx.DB
}

func NewUserResourcePostgres(db *sqlx.DB) *UserResourcePostgres {
	return &UserResourcePostgres{db: db}
}

func (r *UserResourcePostgres) CreateResource(resource models.UserResource) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO %s (user_id, resource_name, resource_login,
	resource_password_hash, date_creation, last_update) VALUES ($1, $2, $3, $4, now(), now())
	RETURNING id`, usersResourcesTable)

	row := r.db.QueryRow(query, resource.UID, resource.ResourceName, resource.ResourceLogin, resource.ResourcePassword)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil

}

func (r *UserResourcePostgres) GetAllResources(userId int) ([]models.UserResource, error) {
	var resources []models.UserResource

	query := fmt.Sprintf(`SELECT users_resources.* FROM %s JOIN %s ON users.id = users_resources.user_id 
	WHERE user_id = $1`, usersResourcesTable, usersTable)

	err := r.db.Select(&resources, query, userId)

	return resources, err
}
