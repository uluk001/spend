package main

import (
	"github.com/uluk001/spend/internal/config"
	"github.com/uluk001/spend/internal/repository/postgres"
	"github.com/uluk001/spend/internal/model"
)


func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load config")
	}

    // Инициализация БД
    db, err := postgres.NewPostgresDB(cfg)
    if err != nil {
        panic("Failed to initialize database")
    }

    err = db.GetDB().AutoMigrate(&model.User{})
    if err != nil {
        panic("Failed to migrate database: " + err.Error())
    }

    // Создание репозиториев
    userRepo := postgres.NewUserRepository(db)
    // transactionRepo := postgres.NewTransactionRepository(db)

	arlen := model.User{
		TelegramID: 1010215859,
		Username:   "ismailovvv001",
	};

	userRepo.Create(&arlen)
}
