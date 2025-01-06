package main

import (
	// "fmt"
	"log"

	"github.com/uluk001/spend/internal/config"
	"github.com/uluk001/spend/internal/model"
	"github.com/uluk001/spend/internal/repository/postgres"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
    // userRepo := postgres.NewUserRepository(db)
    // transactionRepo := postgres.NewTransactionRepository(db)

	// Bot
	bot, err := tgbotapi.NewBotAPI("7658794638:AAFPofqEBRbZqZuIpmalQRaTfjOWJjpGIY4")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
