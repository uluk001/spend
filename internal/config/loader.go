package config

import (
    "github.com/spf13/viper"
)

func LoadConfig() (*Config, error) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("./config")
    
    // Загрузка .env переменных
    viper.AutomaticEnv()
    
    // Замена переменных в конфиге
    viper.SetEnvPrefix("APP")
    
    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }
    
    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, err
    }
    
    return &config, nil
}