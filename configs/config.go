package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Server struct {
		Port string
	} `yaml:"server"`
	DB struct {
		Host     string
		Port     string
		Username string
		Password string `env:"DB_PASSWORD"`
		DBName   string
		SSLMode  string
	} `yaml:"db"`
}

func LoadEnvironment() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	return nil
}

func ReadConfig() Config {
	var cfg Config
	if err := cleanenv.ReadConfig("configs/config.yml", &cfg); err != nil {
		logrus.Fatalf("error loading config file:%s", err.Error())
	}
	return cfg
}
