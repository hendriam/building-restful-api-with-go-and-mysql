package framework

import (
	"os"
	"strconv"
)

type Config struct {
	Server struct {
		Host string
		Port int
	}

	Database struct {
		MySQL struct {
			Dsn string
		}
	}

	Log struct {
		Level string
	}
}

func LoadConfig() Config {
	return Config{
		Server: struct {
			Host string
			Port int
		}{
			Host: os.Getenv("HOST"),
			Port: toInt(os.Getenv("PORT")),
		},
		Database: struct{ MySQL struct{ Dsn string } }{
			MySQL: struct{ Dsn string }{
				Dsn: os.Getenv("DATABASE_MYSQL_DSN"),
			},
		},
		Log: struct{ Level string }{
			Level: os.Getenv("LOG_LEVEL"),
		},
	}
}

func toInt(src string) int {
	number, _ := strconv.Atoi(src)
	return number
}
