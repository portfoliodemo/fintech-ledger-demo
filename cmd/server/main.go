package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"

    "github.com/portfoliodemo/fintech-ledger-demo/internal/service"
    "github.com/portfoliodemo/fintech-ledger-demo/internal/storage"
)

func main() {
    // Initialize in-memory repository
    repo := storage.NewMockRepository()
    ledger := service.NewLedgerService(repo)

    // AddCredit endpoint
    http.HandleFunc("/credit", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            w.WriteHeader(http.StatusMethodNotAllowed)
            return
        }

        var req struct {
            UserID int64   `json:"user_id"`
            Amount float64 `json:"amount"`
        }

        // Decode JSON body
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("invalid request body"))
            return
        }

        tx, err := ledger.AddCredit(req.UserID, req.Amount)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(err.Error()))
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(tx)
    })

	// AddDebit endpoint
	http.HandleFunc("/debit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			UserID int64   `json:"user_id"`
			Amount float64 `json:"amount"`
		}

		// Decode JSON body
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid request body"))
			return
		}

		tx, err := ledger.AddDebit(req.UserID, req.Amount)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tx)
	})


    // GetBalance endpoint
    http.HandleFunc("/balance", func(w http.ResponseWriter, r *http.Request) {
        userIDStr := r.URL.Query().Get("user_id")
        userID, err := strconv.ParseInt(userIDStr, 10, 64)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("invalid user_id"))
            return
        }

        balance, err := ledger.GetUserBalance(userID)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(err.Error()))
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "user_id": userID,
            "balance": balance,
        })
    })

    log.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
