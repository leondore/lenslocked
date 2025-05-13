package database

import "fmt"

type Config struct {
	host     string
	port     string
	user     string
	password string
	database string
	sslmode  string
}

func (cfg Config) String() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.host,
		cfg.port,
		cfg.user,
		cfg.password,
		cfg.database,
		cfg.sslmode,
	)
}

func NewConfig() Config {
	return Config{
		host:     "localhost",
		port:     "5432",
		user:     "ldore",
		password: "7d8jwl59",
		database: "lenslocked",
		sslmode:  "disable",
	}
}
