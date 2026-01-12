package storage

import (
    "errors"
    "fintech-ledger-demo/internal/models"
)

type MockRepository struct {
    Transactions []*models.Transaction
}

func NewMockRepository() *MockRepository {
    return &MockRepository{
        Transactions: []*models.Transaction{},
    }
}

func (m *MockRepository) SaveTransaction(tx *models.Transaction) error {
    if tx == nil {
        return errors.New("transaction is nil")
    }
    m.Transactions = append(m.Transactions, tx)
    return nil
}

func (m *MockRepository) GetTransactionsByUser(userID int64) ([]*models.Transaction, error) {
    var userTxs []*models.Transaction
    for _, tx := range m.Transactions {
        if tx.UserID == userID {
            userTxs = append(userTxs, tx)
        }
    }
    return userTxs, nil
}
