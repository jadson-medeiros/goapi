package configs

import (
	"log"
	"path/filepath"

	"github.com/go-chi/jwtauth"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JwtExpiresIn  int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig() (*conf, error) {
	envPath := filepath.Join("cmd", "server", ".env")

	// Adiciona log para depuração
	log.Println("Loading .env file from:", envPath)

	// Carrega o arquivo .env
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, err
	}

	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(envPath)
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, err
}
