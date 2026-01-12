package storage

import "fintech-ledger-demo/internal/models"

type Repository interface {
    SaveTransaction(tx *models.Transaction) error
    GetTransactionsByUser(userID int64) ([]*models.Transaction, error)
}
