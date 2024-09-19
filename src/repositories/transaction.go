package repositories

import (
	"context"
	"database/sql"
	"dbf_api/models"
	"dbf_api/schemas"
	"dbf_api/utils"
)

type TransactionRepository struct {
    db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
    return &TransactionRepository{
        db: db,
    }
}

func (repo *TransactionRepository) GetByID(ctx context.Context, id int64) (models.Transaction, error) {
    row := repo.db.QueryRowContext(ctx, utils.GetTransaction, id)
	var i models.Transaction
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Cost,
		&i.Time,
        &i.CategoryID,
	)
	return i, err
}

func (repo *TransactionRepository) CreateTransaction(ctx context.Context, arg schemas.CreateTransactionParams) error {
	row := repo.db.QueryRowContext(ctx, utils.CreateTransaction, arg.Name, arg.Cost, arg.Time, arg.CategoryID)
	var i models.Transaction
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Cost,
		&i.Time,
        &i.CategoryID,
	)
    return err
}

func (repo *TransactionRepository) DeleteTransaction(ctx context.Context, id int64) error {
	_, err := repo.db.ExecContext(ctx, utils.DeleteTransaction, id)
	return err
}

func (repo *TransactionRepository) PartialUpdateTransaction(ctx context.Context, arg schemas.PartialUpdateTransactionParams) error {
	row := repo.db.QueryRowContext(ctx, utils.PartialUpdateTransaction,
		arg.UpdateName,
		arg.Name,
		arg.UpdateCost,
		arg.Cost,
		arg.UpdateTime,
		arg.Time,
        arg.UpdateCategoryID,
        arg.CategoryID,
		arg.ID,
	)
	var i models.Transaction
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Cost,
		&i.Time,
        &i.CategoryID,
	)
	return err
}

func (repo *TransactionRepository) ListTransactions(ctx context.Context) ([]models.Transaction, error) {
	rows, err := repo.db.QueryContext(ctx, utils.ListTransactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Transaction
	for rows.Next() {
		var i models.Transaction
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Cost,
			&i.Time,
            &i.CategoryID,
		); err != nil {
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

