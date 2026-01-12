package models

import "time"

type User struct {
    ID   int64
    Name string
}

type Transaction struct {
    ID        int64
    UserID    int64
    Amount    float64
    CreatedAt time.Time
}
