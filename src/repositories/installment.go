package repositories

import (
	"context"
	"database/sql"

	"dbf_api/models"
	"dbf_api/schemas"
	"dbf_api/utils"
)

type InstallmentRepository struct {
	db *sql.DB
}

func NewInstallmentRepository(db *sql.DB) *InstallmentRepository {
	return &InstallmentRepository{
		db: db,
	}
}

func (repo *InstallmentRepository) GetByID(ctx context.Context, id int64) (models.Installment, error) {
	row := repo.db.QueryRowContext(ctx, utils.GetInstallment, id)
	var i models.Installment
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.TotalCost,
		&i.InterestRate,
		&i.PeriodNum,
		&i.PaidCost,
		&i.CurrentPeriod,
		&i.PeriodCost,
		&i.AccountID,
	)
	return i, err
}

func (repo *InstallmentRepository) CreateInstallment(ctx context.Context, arg schemas.CreateInstallmentParams) error {
	row := repo.db.QueryRowContext(ctx, utils.CreateInstallment,
		arg.Name,
		arg.TotalCost,
		arg.InterestRate,
		arg.PeriodNum,
		arg.PaidCost,
		arg.CurrentPeriod,
		arg.PeriodCost,
		arg.AccountID,
	)
	var i models.Installment
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.TotalCost,
		&i.InterestRate,
		&i.PeriodNum,
		&i.PaidCost,
		&i.CurrentPeriod,
		&i.PeriodCost,
		&i.AccountID,
	)
    return err
}

func (repo *InstallmentRepository) DeleteInstallment(ctx context.Context, id int64) error {
	_, err := repo.db.ExecContext(ctx, utils.DeleteInstallment, id)
	return err
}

func (repo *InstallmentRepository) PartialUpdateInstallment(ctx context.Context, arg schemas.PartialUpdateInstallmentParams) error {
	row := repo.db.QueryRowContext(ctx, utils.PartialUpdateInstallment,
		arg.UpdateName,
		arg.Name,
		arg.UpdateTotalCost,
		arg.TotalCost,
		arg.UpdateInterestRate,
		arg.InterestRate,
		arg.UpdatePeriodNum,
		arg.PeriodNum,
		arg.UpdatePaidCost,
		arg.PaidCost,
		arg.UpdateCurrentPeriod,
		arg.CurrentPeriod,
		arg.UpdatePeriodCost,
		arg.PeriodCost,
		arg.UpdateAccount,
		arg.AccountID,
		arg.ID,
	)
	var i models.Installment
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.TotalCost,
		&i.InterestRate,
		&i.PeriodNum,
		&i.PaidCost,
		&i.CurrentPeriod,
		&i.PeriodCost,
		&i.AccountID,
	)
	return err
}

func (repo *InstallmentRepository) ListInstallments(ctx context.Context) ([]models.Installment, error) {
	rows, err := repo.db.QueryContext(ctx, utils.ListInstallments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Installment
	for rows.Next() {
		var i models.Installment
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.TotalCost,
			&i.InterestRate,
			&i.PeriodNum,
			&i.PaidCost,
			&i.CurrentPeriod,
			&i.PeriodCost,
			&i.AccountID,
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

