package repositories

import (
	"context"
	"database/sql"

	"dbf_api/models"
	"dbf_api/schemas"
	"dbf_api/utils"
)

type BalanceSheetRepository struct {
	db *sql.DB
}

func NewBalanceSheetRepository(db *sql.DB) *BalanceSheetRepository {
	return &BalanceSheetRepository{
		db: db,
	}
}

func (repo *BalanceSheetRepository) GetByID(ctx context.Context, id int64) (models.BalanceSheet, error) {
	row := repo.db.QueryRowContext(ctx, utils.GetBalanceSheet, id)
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


func (repo *BalanceSheetRepository) CreateBalanceSheet(ctx context.Context, arg schemas.CreateBalanceSheetParams) error {
	row := repo.db.QueryRowContext(ctx, utils.CreateBalanceSheet, arg.Month, arg.Year, arg.Allocation, arg.Paid, arg.Remaining, arg.CategoryID)
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

func (repo *BalanceSheetRepository) DeleteBalanceSheet(ctx context.Context, id int64) error {
	_, err := repo.db.ExecContext(ctx, utils.CreateBalanceSheet, id)
	return err
}

func (repo *BalanceSheetRepository) PartialUpdateBalanceSheet(ctx context.Context, arg schemas.PartialUpdateBalanceSheetParams) error {
	row := repo.db.QueryRowContext(ctx, utils.PartialUpdateBalanceSheet,
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


func (repo *BalanceSheetRepository) ListBalanceSheets(ctx context.Context) ([]models.BalanceSheet, error) {
	rows, err := repo.db.QueryContext(ctx, utils.ListBalanceSheets)
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
