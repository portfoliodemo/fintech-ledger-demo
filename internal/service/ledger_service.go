package service

import (
    "fmt"
    "time"

    "github.com/portfoliodemo/fintech-ledger-demo/internal/models"
    "github.com/portfoliodemo/fintech-ledger-demo/internal/storage"
)

type LedgerService struct {
    repo storage.Repository
}

// Constructor
func NewLedgerService(repo storage.Repository) *LedgerService {
    return &LedgerService{
        repo: repo,
    }
}

// AddCredit adds a positive credit transaction for a user
func (s *LedgerService) AddCredit(userID int64, amount float64) (*models.Transaction, error) {
    if amount <= 0 {
        return nil, fmt.Errorf("amount must be positive")
    }

    tx := &models.Transaction{
        UserID:    userID,
        Amount:    amount,
        CreatedAt: time.Now(),
    }

    err := s.repo.SaveTransaction(tx)
    if err != nil {
        return nil, err
    }

    return tx, nil
}

// GetUserBalance calculates balance for a user
func (s *LedgerService) GetUserBalance(userID int64) (float64, error) {
    txs, err := s.repo.GetTransactionsByUser(userID)
    if err != nil {
        return 0, err
    }

    balance := 0.0
    for _, tx := range txs {
        balance += tx.Amount
    }

    return balance, nil
}


// AddDebit subtracts a positive amount from the user's balance if sufficient funds exist
func (s *LedgerService) AddDebit(userID int64, amount float64) (*models.Transaction, error) {
    if amount <= 0 {
        return nil, fmt.Errorf("amount must be positive")
    }

    // Check current balance
    balance, err := s.GetUserBalance(userID)
    if err != nil {
        return nil, err
    }

    if amount > balance {
        return nil, fmt.Errorf("insufficient funds: current balance %.2f", balance)
    }

    tx := &models.Transaction{
        UserID:    userID,
        Amount:    -amount, // negative because it's a debit
        CreatedAt: time.Now(),
    }

    err = s.repo.SaveTransaction(tx)
    if err != nil {
        return nil, err
    }

    return tx, nil
}
