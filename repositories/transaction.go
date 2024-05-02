package repositories

import (
	"context"
	"database/sql"
	"dbf_api/models"
	"dbf_api/schemas"
)

type Repository interface {
    CreateTransaction(ctx context.Context, arg schemas.CreateTransactionParams) error
    UpdateTransaction(ctx context.Context, arg schemas.PartialUpdateTransactionParams) error
    DeleteTransaction(ctx context.Context, id int64) error
    GetByID(ctx context.Context, id int64) (*models.Transaction, error)
    ListTransactions(ctx context.Context) ([]models.Transaction, error)
}

type TransactionRepository struct {
    db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
    return &TransactionRepository{
        db: db,
    }
}

const getTransaction = `-- name: GetTransaction :one
select id, name, cost, time
from transaction
where id = $1
limit 1
`
func (repo *TransactionRepository) GetByID(ctx context.Context, id int64) (*models.Transaction, error) {
    row := repo.db.QueryRowContext(ctx, getTransaction, id)
	var i *models.Transaction
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Cost,
		&i.Time,
	)
	return i, err
}


const createTransaction = `-- name: CreateTransaction :one
insert into transaction (name, cost, time)
values($1, $2, $3)
returning id, name, cost, time
`
func (repo *TransactionRepository) CreateTransaction(ctx context.Context, arg schemas.CreateTransactionParams) error {
	row := repo.db.QueryRowContext(ctx, createTransaction, arg.Name, arg.Cost, arg.Time)
	var i models.Transaction
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Cost,
		&i.Time,
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
set name = case when $1::boolean then $2::VARCHAR(255) else name end,
    cost = case when $3::boolean then $4::real else cost end,
    time = case when $5::boolean then $6::timestamp else time end
where id = $7
returning id, name, cost, time
`
func (repo *TransactionRepository) PartialUpdateTransaction(ctx context.Context, arg schemas.PartialUpdateTransactionParams) error {
	row := repo.db.QueryRowContext(ctx, partialUpdateTransaction,
		arg.UpdateName,
		arg.Name,
		arg.UpdateCost,
		arg.Cost,
		arg.UpdateTime,
		arg.Time,
		arg.ID,
	)
	var i models.Transaction
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Cost,
		&i.Time,
	)
	return err
}

const listTransactions = `-- name: ListTransactions :many
select id, name, cost, time
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

