package configs

type Config struct {
	Server struct {
		Port string
	} `yaml:"server"`
	DB struct {
		Host     string
		Port     string
		Username string
		Password string
		DBName   string
		SSLMode  string
	} `yaml:"db"`
}
