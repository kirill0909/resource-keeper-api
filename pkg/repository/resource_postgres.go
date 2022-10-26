package repository

import (
	"errors"
	"fmt"
	"strings"

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

	query := fmt.Sprintf(`INSERT INTO %s (user_id, resource_name, resource_login_enc,
	resource_password_enc, date_creation, last_update) VALUES ($1, $2, $3, $4, now(), now())
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

func (r *UserResourcePostgres) GetById(userId, resourceId int) (models.UserResource, error) {
	var resource models.UserResource

	query := fmt.Sprintf(`SELECT users_resources.* FROM %s JOIN %s ON users.id = users_resources.user_id
	WHERE user_id = $1 AND users_resources.id = $2`, usersResourcesTable, usersTable)

	err := r.db.Get(&resource, query, userId, resourceId)

	return resource, err
}

func (r *UserResourcePostgres) DeleteResource(userId, resourceId int) error {
	var id int

	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND id = $2 RETURNING id",
		usersResourcesTable)
	row := r.db.QueryRow(query, userId, resourceId)
	if err := row.Scan(&id); err != nil {
		return errors.New("resource id not found")
	}

	return nil
}

func (r *UserResourcePostgres) UpdateResource(userId, resourceId int, input models.UserResourceUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.ResourceName != nil && len(strings.TrimSpace(*input.ResourceName)) != 0 {
		setValues = append(setValues, fmt.Sprintf("resource_name=$%d", argId))
		args = append(args, *input.ResourceName)
		argId++
	}

	if input.ResourceLogin != nil && len(strings.TrimSpace(*input.ResourceLogin)) != 0 {
		setValues = append(setValues, fmt.Sprintf("resource_login_enc=$%d", argId))
		args = append(args, *input.ResourceLogin)
		argId++
	}

	if input.ResourcePassword != nil && len(strings.TrimSpace(*input.ResourcePassword)) != 0 {
		setValues = append(setValues, fmt.Sprintf("resource_password_enc=$%d", argId))
		args = append(args, *input.ResourcePassword)
	}

	if len(setValues) == 0 {
		return errors.New("no new value for set")
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s, last_update = now() WHERE user_id=%d AND id=%d", usersResourcesTable,
		setQuery, userId, resourceId)

	_, err := r.db.Exec(query, args...)

	return err
}
