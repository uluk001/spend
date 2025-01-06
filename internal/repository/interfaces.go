package repository

import "github.com/uluk001/spend/internal/model"

type UserRepository interface {
    Create(user *model.User) error
    GetByID(id int) (*model.User, error)
    Update(id int, user *model.User) (*model.User, error)
    Delete(id int) error
}

type TransactionRepository interface {
    Create(transaction *model.Transaction) error
    GetByUserID(userID int) ([]model.Transaction, error)
}