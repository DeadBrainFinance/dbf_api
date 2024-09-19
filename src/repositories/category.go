package repositories

import (
	"context"
	"database/sql"

	"dbf_api/models"
	"dbf_api/schemas"
	"dbf_api/utils"
)

type CategoryRepository struct {
    db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
    return &CategoryRepository{
        db: db,
    }
}

func (repo *CategoryRepository) CreateCategory(ctx context.Context, name string) error {
	row := repo.db.QueryRowContext(ctx, utils.CreateCategory, name)
	var i models.Category
	err := row.Scan(&i.ID, &i.Name)
	return err
}

func (repo *CategoryRepository) GetByID(ctx context.Context, id int64) (models.Category, error) {
	row := repo.db.QueryRowContext(ctx, utils.GetCategory, id)
	var i models.Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

func (repo *CategoryRepository) PartialUpdateCategory(ctx context.Context, arg schemas.PartialUpdateCategoryParams) error {
	row := repo.db.QueryRowContext(ctx, utils.PartialUpdateCategory, arg.ID, arg.Name)
	var i models.Category
	err := row.Scan(&i.ID, &i.Name)
	return err
}

func (repo *CategoryRepository) ListCategories(ctx context.Context) ([]models.Category, error) {
	rows, err := repo.db.QueryContext(ctx, utils.ListCategories)
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

func (repo *CategoryRepository) DeleteCategory(ctx context.Context, id int64) error {
	_, err := repo.db.ExecContext(ctx, utils.DeleteCategory, id)
	return err
}
