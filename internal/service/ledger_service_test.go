package service

import (
    "testing"
    "github.com/portfoliodemo/fintech-ledger-demo/internal/storage"
)

func TestAddCredit(t *testing.T) {
    repo := storage.NewMockRepository()
    service := NewLedgerService(repo)

    // Test valid credit
    tx, err := service.AddCredit(1, 100.0)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if tx.Amount != 100.0 {
        t.Errorf("expected 100.0, got %v", tx.Amount)
    }
    if tx.UserID != 1 {
        t.Errorf("expected userID 1, got %v", tx.UserID)
    }

    // Test invalid credit (negative)
    _, err = service.AddCredit(1, -50.0)
    if err == nil {
        t.Errorf("expected error for negative amount, got nil")
    }

    // Test balance calculation
    balance, err := service.GetUserBalance(1)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    expected := 100.0
    if balance != expected {
        t.Errorf("expected balance %v, got %v", expected, balance)
    }
}

func TestAddDebit(t *testing.T) {
    repo := storage.NewMockRepository()
    ledger := NewLedgerService(repo)

    // Seed a user with 100 credits
    _, _ = ledger.AddCredit(1, 100.0)

    // Successful debit
    tx, err := ledger.AddDebit(1, 50.0)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if tx.Amount != -50.0 {
        t.Errorf("expected -50.0, got %v", tx.Amount)
    }

    // Check balance
    balance, _ := ledger.GetUserBalance(1)
    expected := 50.0
    if balance != expected {
        t.Errorf("expected balance %v, got %v", expected, balance)
    }

    // Debit exceeding balance
    _, err = ledger.AddDebit(1, 100.0)
    if err == nil {
        t.Errorf("expected error for insufficient funds, got nil")
    }
}

