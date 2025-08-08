package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type DBConfig struct {
	Host     string `toml:"db_host"`
	Port     string `toml:"db_port"`
	User     string `toml:"db_user"`
	Password string `toml:"db_password"`
	Name     string `toml:"db_name"`
}

func GetDBConfig() string {
	isProd := IsProduction()

	if isProd {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		name := os.Getenv("DB_NAME")

		return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
			host, user, password, name, port)
	}

	// Development or other env - load from dev.toml
	devFile := "dev.toml"
	content, err := os.ReadFile(devFile)
	if err != nil {
		panic("Failed to read dev file: " + err.Error())
	}

	var conf DBConfig
	if _, err := toml.Decode(string(content), &conf); err != nil {
		panic("Failed to parse dev file: " + err.Error())
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		conf.Host, conf.User, conf.Password, conf.Name, conf.Port)
}
