package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

const (
	defaultServerAddress = "localhost:8080"
)

// Config - this is a structure for storing app init params
type Config struct {
	ServerAddress string `json:"server_address"`
	DatabaseDSN   string `json:"database_dsn"`
}

// ConfigBuilder - структура, возвращающая подготовленный кофиг в ходе инициализации приложения
type ConfigBuilder struct {
	c Config
}

// SetServerAddress - метод установки значения хоста в конфиг инициализации приложения
func (cb *ConfigBuilder) SetServerAddress(host string) *ConfigBuilder {
	if cb.c.ServerAddress != "" && host == defaultServerAddress {
		return cb
	}

	cb.c.ServerAddress = host
	return cb
}

// SetDatabaseDSN - метод установки заначения строки конфига для инициализации БД
func (cb *ConfigBuilder) SetDatabaseDSN(databaseDSN string) *ConfigBuilder {
	if cb.c.DatabaseDSN != "" && databaseDSN == "" {
		return cb
	}

	cb.c.DatabaseDSN = databaseDSN
	return cb
}

// Build - метод для формирования результирующего конфига для инициализации приложения
func (cb *ConfigBuilder) Build() Config {
	return cb.c
}

func (cb *ConfigBuilder) loadFromFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&cb.c)
	if err != nil {
		log.Fatal(err)
	}
}

// NewConfigBuilder - метод вызываемый для определения значений конфига инициализации при старте приложения
func NewConfigBuilder() Config {
	var cb ConfigBuilder

	var configFile string
	flag.StringVar(&configFile, "c", "", "path to json config file")

	var serverAddress string
	flag.StringVar(&serverAddress, "a", defaultServerAddress, "server host")

	var databaseDSN string
	flag.StringVar(&databaseDSN, "d", "", "Database DSN config string")
	// flag.StringVar(&databaseDSN, "d", "host=localhost port=5432 user=postgres password=root dbname=gophkeeper sslmode=disable", "Database DSN config string")

	flag.Parse()

	if envConfigFile := os.Getenv("CONFIG"); envConfigFile != "" {
		configFile = envConfigFile
	}

	if envServerSddress := os.Getenv("SERVER_ADDRES"); envServerSddress != "" {
		serverAddress = envServerSddress
	}

	if envDatabaseDSN := os.Getenv("DATABASE_DSN"); envDatabaseDSN != "" {
		databaseDSN = envDatabaseDSN
	}

	if configFile != "" {
		cb.loadFromFile(configFile)
	}

	return cb.
		SetServerAddress(serverAddress).
		SetDatabaseDSN(databaseDSN).
		Build()
}
