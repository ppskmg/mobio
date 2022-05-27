package apiserver

import (
	"log"
	"net/http"
	"os"
)

func Start(config *Config, tls bool) error {
	srv := newServer()
	LOG_FILE := "./mobio_log"
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	if tls {
		log.Println("ListenAndServeTLS:", config.PortTLS)
		return http.ListenAndServeTLS(config.PortTLS, config.SSLCrt, config.SSLKey, srv)
	}
	log.Println("ListenAndServe:", config.Port)
	return http.ListenAndServe(config.Port, srv)
}
