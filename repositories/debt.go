package repositories

import (
	"context"
	"database/sql"
	"dbf_api/models"
	"dbf_api/schemas"
	"dbf_api/utils"
)

type DebtRepository struct {
	db *sql.DB
}

func NewDebtRepository(db *sql.DB) *DebtRepository {
	return &DebtRepository{
		db: db,
	}
}

func (repo *DebtRepository) GetByID(ctx context.Context, id int64) (models.Debt, error) {
	row := repo.db.QueryRowContext(ctx, utils.GetDebt, id)
	var i models.Debt

	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Lender,
		&i.Borrower,
		&i.InterestRate,
		&i.BorrowedAmount,
		&i.PaidAmount,
		&i.LendDate,
	)
	return i, err
}

func (repo *DebtRepository) CreateDebt(ctx context.Context, arg schemas.CreateDebtParams) error {
	row := repo.db.QueryRowContext(ctx, utils.CreateDebt, arg.Name, arg.Lender, arg.Borrower, arg.InterestRate, arg.BorrowedAmount, arg.PaidAmount, arg.LendDate)
	var i models.Debt
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Lender,
		&i.Borrower,
		&i.InterestRate,
		&i.BorrowedAmount,
		&i.PaidAmount,
		&i.LendDate,
	)
	return err
}

func (repo *DebtRepository) DeleteDebt(ctx context.Context, id int64) error {
	_, err := repo.db.ExecContext(ctx, utils.DeleteDebt, id)
	return err
}

func (repo *DebtRepository) PartialUpdateDebt(ctx context.Context, arg schemas.PartialUpdateDebtParams) error {
	row := repo.db.QueryRowContext(ctx, utils.PartialUpdateDebt,
		arg.UpdateName,
		arg.Name,
		arg.UpdateLender,
		arg.Lender,
		arg.UpdateBorrower,
		arg.InterestRate,
		arg.UpdateInterestRate,
		arg.BorrowedAmount,
		arg.UpdateBorrowedAmount,
		arg.PaidAmount,
		arg.UpdatePaidAmount,
		arg.LendDate,
		arg.UpdateLendDate,
	)
	var i models.Debt
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Lender,
		&i.Borrower,
		&i.InterestRate,
		&i.BorrowedAmount,
		&i.PaidAmount,
		&i.LendDate,
	)
	return err
}

func (repo *DebtRepository) ListDebts(ctx context.Context) ([]models.Debt, error) {
	rows, err := repo.db.QueryContext(ctx, utils.ListDebts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Debt
	for rows.Next() {
		var i models.Debt
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Lender,
			&i.Borrower,
			&i.InterestRate,
			&i.BorrowedAmount,
			&i.PaidAmount,
			&i.LendDate,
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
