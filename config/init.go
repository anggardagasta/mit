package config

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Route *chi.Mux
	DB    *sql.DB
}

func Init() *Config {
	var cfg Config

	err := cfg.InitMysql()
	if err != nil {
		log.Panic(err)
	}

	err = cfg.InitService()
	if err != nil {
		log.Panic(err)
	}

	return &cfg
}

func (ox Config) Start() {
	fmt.Println("Listening port 8082")
	http.ListenAndServe(":8082", ox.Route)
}

func Env(key string, def string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return def
}

func U2Replace(s string, v map[string]string) string {
	for k, v := range v {
		s = strings.ReplaceAll(s, "__"+k+"__", v)
	}
	return s
}
