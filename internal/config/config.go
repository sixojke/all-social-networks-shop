package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	yamlPath = "configs"
	yamlFile = "config"
	envFile  = ".env"
)

type Config struct {
	Postgres    Postgres
	Redis       Redis
	Cache       Cache
	HTTPServer  HTTPServer
	Handler     Handler
	Service     Service
	EmailSender EmailSender
	Payok       Payok
}

type EmailSender struct {
	From     string `mapstructure:"from"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}

type Postgres struct {
	Username string
	Password string
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DBName   string `mapstructure:"db_name"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

type Redis struct {
	Password string
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DBName   int    `mapstructure:"db_name"`
}

type Cache struct {
	Expiration time.Duration `mapstructure:"expiration"`
}

type HTTPServer struct {
	Port           string        `mapstructure:"port"`
	ReadTimeout    time.Duration `mapstructure:"read_timeout"`
	WriteTimeout   time.Duration `mapstructure:"write_timeout"`
	MaxHeaderBytes int           `mapstructure:"max_header_bytes"`
}

type Handler struct {
}

type Service struct {
	Users UsersService `mapstructure:"users"`
}

type UsersService struct {
	Auth Auth `mapstructure:"auth"`
}
type Auth struct {
	JWT                    JWT `mapstructure:"jwt"`
	VerificationCodeLength int `mapstructure:"verification_code_length"`
	SigningKey             string
}

type JWT struct {
	AccessTokenTTL  time.Duration `mapstructure:"access_token_ttl"`
	RefreshTokenTTL time.Duration `mapstructure:"refresh_token_ttl"`
}

type Payok struct {
	ShopId     string `mapstructure:"shop_id"`
	SuccessUrl string `mapstructure:"success_url"`
	SecretKey  string
}

func InitConfig() (*Config, error) {
	if err := read(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("email_sender", &cfg.EmailSender); err != nil {
		return fmt.Errorf("unmarshal email_sender config: %v", err)
	}

	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
		return fmt.Errorf("unmarshal postgres config: %v", err)
	}

	if err := viper.UnmarshalKey("redis", &cfg.Redis); err != nil {
		return fmt.Errorf("unmarshal redis config: %v", err)
	}

	if err := viper.UnmarshalKey("cache", &cfg.Cache); err != nil {
		return fmt.Errorf("unmarshal cache config: %v", err)
	}

	if err := viper.UnmarshalKey("http_server", &cfg.HTTPServer); err != nil {
		return fmt.Errorf("unmarshal http server config: %v", err)
	}

	if err := viper.UnmarshalKey("service", &cfg.Service); err != nil {
		return fmt.Errorf("unmarshal http server config: %v", err)
	}

	if err := viper.UnmarshalKey("payok", &cfg.Payok); err != nil {
		return fmt.Errorf("unmarshal email_sender config: %v", err)
	}

	if err := viper.UnmarshalKey("handler", &cfg.Handler); err != nil {
		return fmt.Errorf("unmarshal handler config: %v", err)
	}

	cfg.Payok.SecretKey = os.Getenv("PAYOK_KEY")

	cfg.Service.Users.Auth.SigningKey = os.Getenv("SIGNING_KEY")

	cfg.Postgres.Username = os.Getenv("POSTGRES_USER")
	cfg.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")

	return nil
}

func read() error {
	if err := godotenv.Load(envFile); err != nil {
		return fmt.Errorf("read env file: %s: %s", envFile, err)
	}

	viper.AddConfigPath(yamlPath)
	viper.SetConfigName(yamlFile)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read yaml file: %v: %v", envFile, err)
	}

	return nil
}
