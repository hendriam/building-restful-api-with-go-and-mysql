package framework

import (
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
			Host: "0.0.0.0",
			Port: toInt("8080"),
		},
		Database: struct{ MySQL struct{ Dsn string } }{
			MySQL: struct{ Dsn string }{
				// Dsn: os.Getenv("DATABASE_MYSQL_DSN"),
				Dsn: "root:@tcp(localhost:3306)/movie_db",
			},
		},
		Log: struct{ Level string }{
			// Level: os.Getenv("LOG_LEVEL"),
			Level: "info",
		},
	}
}

func toInt(src string) int {
	number, _ := strconv.Atoi(src)
	return number
}
