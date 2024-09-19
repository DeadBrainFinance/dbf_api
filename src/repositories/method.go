package repositories

import (
	"context"
	"database/sql"

	"dbf_api/models"
	"dbf_api/schemas"
	"dbf_api/utils"
)

type MethodRepository struct {
    db *sql.DB
}

func NewMethodRepository(db *sql.DB) *MethodRepository {
    return &MethodRepository{
        db: db,
    }
}

func (repo *MethodRepository) CreateMethod(ctx context.Context, name string) error {
	row := repo.db.QueryRowContext(ctx, utils.CreateMethod, name)
	var i models.Method
	err := row.Scan(&i.ID, &i.Name)
	return err
}

func (repo *MethodRepository) GetByID(ctx context.Context, id int64) (models.Method, error) {
	row := repo.db.QueryRowContext(ctx, utils.GetMethod, id)
	var i models.Method
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

func (repo *MethodRepository) PartialUpdateMethod(ctx context.Context, arg schemas.PartialUpdateMethodParams) error {
	row := repo.db.QueryRowContext(ctx, utils.PartialUpdateMethod, arg.ID, arg.Name)
	var i models.Method
	err := row.Scan(&i.ID, &i.Name)
	return err
}

func (repo *MethodRepository) ListMethods(ctx context.Context) ([]models.Method, error) {
	rows, err := repo.db.QueryContext(ctx, utils.ListMethods)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Method
	for rows.Next() {
		var i models.Method
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


func (repo *MethodRepository) DeleteMethod(ctx context.Context, id int64) error {
	_, err := repo.db.ExecContext(ctx, utils.DeleteMethod, id)
	return err
}

