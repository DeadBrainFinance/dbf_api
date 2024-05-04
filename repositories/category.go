package repositories

import (
	"context"
	"database/sql"
	"dbf_api/models"
	"dbf_api/schemas"
)

type CategoryRepository struct {
    db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
    return &CategoryRepository{
        db: db,
    }
}

const createCategory = `-- name: CreateCategory :one
insert into category (name)
values($1)
returning id, name
`
func (repo *CategoryRepository) CreateCategory(ctx context.Context, name string) error {
	row := repo.db.QueryRowContext(ctx, createCategory, name)
	var i models.Category
	err := row.Scan(&i.ID, &i.Name)
	return err
}

const getCategory = `-- name: GetCategory :one
select id, name
from category
where id = $1
limit 1
`
func (repo *CategoryRepository) GetByID(ctx context.Context, id int64) (models.Category, error) {
	row := repo.db.QueryRowContext(ctx, getCategory, id)
	var i models.Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const partialUpdateCategory = `-- name: PartialUpdateCategory :one
update category
set name = $2
where id = $1
returning id, name
`
func (repo *CategoryRepository) PartialUpdateCategory(ctx context.Context, arg schemas.PartialUpdateCategoryParams) error {
	row := repo.db.QueryRowContext(ctx, partialUpdateCategory, arg.ID, arg.Name)
	var i models.Category
	err := row.Scan(&i.ID, &i.Name)
	return err
}

const listCategories = `-- name: ListCategories :many
select id, name
from category
order by name
`
func (repo *CategoryRepository) ListCategories(ctx context.Context) ([]models.Category, error) {
	rows, err := repo.db.QueryContext(ctx, listCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Category
	for rows.Next() {
		var i models.Category
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}


const deleteCategory = `-- name: DeleteCategory :exec
delete
from category
where id = $1
`
func (repo *CategoryRepository) DeleteCategory(ctx context.Context, id int64) error {
	_, err := repo.db.ExecContext(ctx, deleteCategory, id)
	return err
}
