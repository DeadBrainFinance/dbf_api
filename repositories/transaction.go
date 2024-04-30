package repositories

import "dbf_api/models"

type TransactionRepository interface {
    CreateTransaction(transaction *models.Transaction) error
    UpdateTransaction(params *models.PartialUpdateTransactionParams) error
    DeleteTransaction(id int64) error
    GetByID(id int64) (*models.Transaction, error)
    GettAll() ([]*models.Transaction, error)
}
