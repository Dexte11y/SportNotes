package configs

type Config struct {
	Server struct {
		Port string
	} `yaml:"server"`
	Db struct {
		Host     string
		Port     string
		Username string
		Password string
		DBName   string
		SSLMode  string
	} `yaml:"db"`
}
