package repositories

import (
	"context"
	"database/sql"
	"dbf_api/models"
	"dbf_api/schemas"
)

type TransactionRepository struct {
    db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
    return &TransactionRepository{
        db: db,
    }
}

const getTransaction = `-- name: GetTransaction :one
select id, name, cost, time, category_id
from transaction
where id = $1
limit 1
`
func (repo *TransactionRepository) GetByID(ctx context.Context, id int64) (models.Transaction, error) {
    row := repo.db.QueryRowContext(ctx, getTransaction, id)
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

const createTransaction = `-- name: CreateTransaction :one
insert into transaction (name, cost, time, category_id)
values($1, $2, $3, $4)
returning id, name, cost, time, category_id
`
func (repo *TransactionRepository) CreateTransaction(ctx context.Context, arg schemas.CreateTransactionParams) error {
	row := repo.db.QueryRowContext(ctx, createTransaction, arg.Name, arg.Cost, arg.Time, arg.CategoryID)
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

const deleteTransaction = `-- name: DeleteTransaction :exec
delete
from transaction
where id = $1
`
func (repo *TransactionRepository) DeleteTransaction(ctx context.Context, id int64) error {
	_, err := repo.db.ExecContext(ctx, deleteTransaction, id)
	return err
}

const partialUpdateTransaction = `-- name: PartialUpdateTransaction :one
update transaction
set name = case when $1::boolean then $2::varchar(255) else name end,
    cost = case when $3::boolean then $4::real else cost end,
    time = case when $5::boolean then $6::timestamp else time end,
    category_id = case when $7::boolean then $8::int else category_id end
where id = $9
returning id, name, cost, time, category_id
`
func (repo *TransactionRepository) PartialUpdateTransaction(ctx context.Context, arg schemas.PartialUpdateTransactionParams) error {
	row := repo.db.QueryRowContext(ctx, partialUpdateTransaction,
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

const listTransactions = `-- name: ListTransactions :many
select id, name, cost, time, category_id
from transaction
order by name
`

func (repo *TransactionRepository) ListTransactions(ctx context.Context) ([]models.Transaction, error) {
	rows, err := repo.db.QueryContext(ctx, listTransactions)
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

