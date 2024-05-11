package repositories

import (
	"context"
	"database/sql"

	"dbf_api/models"
	"dbf_api/schemas"
	"dbf_api/utils"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (repo *AccountRepository) GetByID(ctx context.Context, id int64) (models.Account, error) {
	row := repo.db.QueryRowContext(ctx, utils.CreateAccount, id)
	var i models.Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AccBalance,
		&i.AccNum,
		&i.CardNum,
		&i.Pin,
		&i.SecurityCode,
		&i.CreditLimit,
		&i.MethodID,
	)
	return i, err
}

func (repo *AccountRepository) CreateAccount(ctx context.Context, arg schemas.CreateAccountParams) error {
	row := repo.db.QueryRowContext(ctx, utils.CreateAccount,
		arg.Name,
		arg.AccBalance,
		arg.AccNum,
		arg.CardNum,
		arg.Pin,
		arg.SecurityCode,
		arg.CreditLimit,
		arg.MethodID,
	)
	var i models.Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AccBalance,
		&i.AccNum,
		&i.CardNum,
		&i.Pin,
		&i.SecurityCode,
		&i.CreditLimit,
		&i.MethodID,
	)
    return err
}

func (repo *AccountRepository) DeleteAccount(ctx context.Context, id int64) error {
	_, err := repo.db.ExecContext(ctx, utils.DeleteAccount, id)
	return err
}

func (repo *AccountRepository) PartialUpdateAccount(ctx context.Context, arg schemas.PartialUpdateAccountParams) error {
	row := repo.db.QueryRowContext(ctx, utils.PartialUpdateAccount,
		arg.UpdateName,
		arg.Name,
		arg.UpdateAccBalance,
		arg.AccBalance,
		arg.UpdateAccNum,
		arg.AccNum,
		arg.UpdateCardNum,
		arg.CardNum,
		arg.UpdatePin,
		arg.Pin,
		arg.UpdateSecurityCode,
		arg.SecurityCode,
		arg.UpdateCreditLimit,
		arg.CreditLimit,
		arg.UpdateMethod,
		arg.MethodID,
		arg.ID,
	)
	var i models.Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AccBalance,
		&i.AccNum,
		&i.CardNum,
		&i.Pin,
		&i.SecurityCode,
		&i.CreditLimit,
		&i.MethodID,
	)
	return err
}

func (repo *AccountRepository) ListAccounts(ctx context.Context) ([]models.Account, error) {
	rows, err := repo.db.QueryContext(ctx, utils.ListAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Account
	for rows.Next() {
		var i models.Account
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.AccBalance,
			&i.AccNum,
			&i.CardNum,
			&i.Pin,
			&i.SecurityCode,
			&i.CreditLimit,
			&i.MethodID,
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
