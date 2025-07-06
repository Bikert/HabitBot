package config

import (
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
	"sync"
)

type Config struct {
	TGToken    string
	WebBaseUrl string
	DBFilePath string
}

var (
	cfg  *Config
	once sync.Once
)

func Init() {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found")
		}
		baseUrl := getHostURL(getEnv("WEB_BASE_URL"))
		cfg = &Config{
			TGToken:    getEnv("TG_TOKEN"),
			WebBaseUrl: baseUrl,
			DBFilePath: getEnv("DB_FILE_PATH"),
		}
	})
}

func getEnv(key string) string {
	result, exist := os.LookupEnv(key)
	if !exist {
		log.Fatal("Variable " + key + " not found")
	}
	return result
}

func getHostURL(urlStr string) string {
	parsed, err := url.Parse(urlStr)
	if err != nil {
		log.Fatalf("Invalid URL: %v", err)
	}
	baseUrl := parsed.Scheme + "://" + parsed.Host + "/"
	return baseUrl
}

func Get() *Config {
	if cfg == nil {
		log.Fatal("Config is not initialized.")
	}
	return cfg
}
