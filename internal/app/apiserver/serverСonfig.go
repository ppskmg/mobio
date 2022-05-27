package apiserver

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port     string
	PortTLS  string
	SSLCrt   string
	SSLKey   string
	LogLevel string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	port, _ := os.LookupEnv("PORT")
	portTls, _ := os.LookupEnv("PORT_TLS")
	sslCrt, _ := os.LookupEnv("SSL_CRT")
	sslKey, _ := os.LookupEnv("SSL_KEY")

	return &Config{
		Port:        port,
		PortTLS:     portTls,
		SSLCrt:      sslCrt,
		SSLKey:      sslKey,
		LogLevel:    "debug",
	}
}
