package config

import "os"

type Config struct {
	App      App
	Database Database
	Jwt      Jwt
}

func GetAll() *Config {
	cfg := &Config{
		App: App{
			Port: os.Getenv("APPPORT"),
		},
		Database: Database{
			Host: os.Getenv("DBHOST"),
			Port: os.Getenv("DBPORT"),
			User: os.Getenv("DBUSER"),
			Pass: os.Getenv("DBPASS"),
			Name: os.Getenv("DBNAME"),
		},
		Jwt: Jwt{
			Secret: os.Getenv("JWT_SECRET"),
		},
	}

	return cfg
}
