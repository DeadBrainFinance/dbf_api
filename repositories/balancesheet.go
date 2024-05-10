package repositories

import (
	"context"
	"database/sql"

	"dbf_api/models"
	"dbf_api/schemas"
)

type BalanceSheetRepository struct {
	db *sql.DB
}

func NewBalanceSheetRepository(db *sql.DB) *BalanceSheetRepository {
	return &BalanceSheetRepository{
		db: db,
	}
}

const getBalanceSheet = `-- name: GetBalanceSheet :one
select id, month, year, allocation, paid, remaining, category_id
from balancesheet
where id = $1
limit 1
`
func (repo *BalanceSheetRepository) GetByID(ctx context.Context, id int64) (models.BalanceSheet, error) {
	row := repo.db.QueryRowContext(ctx, getBalanceSheet, id)
	var i models.BalanceSheet
	err := row.Scan(
		&i.ID,
		&i.Month,
		&i.Year,
		&i.Allocation,
		&i.Paid,
		&i.Remaining,
		&i.CategoryID,
	)
	return i, err
}

const createBalanceSheet = `-- name: CreateBalanceSheet :one
insert into balancesheet (month, year, allocation, paid, remaining, category_id)
values($1, $2, $3, $4, $5, $6)
returning id, month, year, allocation, paid, remaining, category_id
`

func (repo *BalanceSheetRepository) CreateBalanceSheet(ctx context.Context, arg schemas.CreateBalanceSheetParams) error {
	row := repo.db.QueryRowContext(ctx, createBalanceSheet, arg.Month, arg.Year, arg.Allocation, arg.Paid, arg.Remaining, arg.CategoryID)
	var i models.BalanceSheet
	err := row.Scan(
        &i.ID,
		&i.Month,
		&i.Year,
		&i.Allocation,
		&i.Paid,
		&i.Remaining,
		&i.CategoryID,
	)
	return err
}

const deleteBalanceSheet = `-- name: DeleteBalanceSheet :exec
delete
from balancesheet
where id = $1
`

func (repo *BalanceSheetRepository) DeleteBalanceSheet(ctx context.Context, id int64) error {
	_, err := repo.db.ExecContext(ctx, deleteBalanceSheet, id)
	return err
}

const partialUpdateBalanceSheet = `-- name: PartialUpdateBalanceSheet :one
update balancesheet
set month = case when $1::boolean then $2::int else month end,
    year = case when $3::boolean then $4::int else year end,
    allocation = case when $5::boolean then $6::real else allocation end,
    paid = case when $7::boolean then $8::real else paid end,
    remaining = case when $9::boolean then $10::real else remaining end,
    category_id = case when $11::boolean then $12::int else category_id end
where id = $13
returning id, month, year, allocation, paid, remaining, category_id
`

func (repo *BalanceSheetRepository) PartialUpdateBalanceSheet(ctx context.Context, arg schemas.PartialUpdateBalanceSheetParams) error {
	row := repo.db.QueryRowContext(ctx, partialUpdateBalanceSheet,
		arg.UpdateMonth,
		arg.Month,
		arg.UpdateYear,
		arg.Year,
		arg.UpdateAllocation,
		arg.Allocation,
		arg.UpdatePaid,
		arg.Paid,
		arg.UpdateRemaining,
		arg.Remaining,
		arg.UpdateCategories,
		arg.CategoryID,
		arg.ID,
	)
	var i models.BalanceSheet
	err := row.Scan(
		&i.ID,
		&i.Month,
		&i.Year,
		&i.Allocation,
		&i.Paid,
		&i.Remaining,
		&i.CategoryID,
	)
	return err
}

const listBalanceSheets = `-- name: ListBalanceSheets :many
select id, month, year, allocation, paid, remaining, category_id
from balancesheet
order by year desc, month desc
`

func (repo *BalanceSheetRepository) ListBalanceSheets(ctx context.Context) ([]models.BalanceSheet, error) {
	rows, err := repo.db.QueryContext(ctx, listBalanceSheets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.BalanceSheet
	for rows.Next() {
		var i models.BalanceSheet
		if err := rows.Scan(
			&i.ID,
			&i.Month,
			&i.Year,
			&i.Allocation,
			&i.Paid,
			&i.Remaining,
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
