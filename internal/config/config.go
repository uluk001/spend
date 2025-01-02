package config

type Config struct {
    Database struct {
        Host     string
        Port     int
        User     string
        Password string
        Name     string
    }
    Telegram struct {
        Token string
    }
}