package config

import (
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
	"strings"
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
		baseUrl := normalizeURL(getEnv("WEB_BASE_URL"))
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

func normalizeURL(urlStr string) string {
	parsed, err := url.Parse(urlStr)
	if err != nil {
		log.Fatalf("Invalid URL: %v", err)
	}

	if !strings.HasSuffix(parsed.Path, "/") {
		parsed.Path += "/"
	}

	return parsed.String()
}

func Get() *Config {
	if cfg == nil {
		log.Fatal("Config is not initialized.")
	}
	return cfg
}
