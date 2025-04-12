package main

import (
	"log"
	"os"
	"strings"

	"Zenick-Lab/zenick-aggregator-server/src/model"
	"Zenick-Lab/zenick-aggregator-server/src/pkg/postgresql"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("error while reading config file: %s", err.Error())
		return
	}
	log.Println("Config file loaded successfully")

	for _, env := range viper.AllKeys() {
		if viper.GetString(env) != "" {
			_ = os.Setenv(env, viper.GetString(env))
			_ = os.Setenv(strings.ToUpper(env), viper.GetString(env))
		}
	}

	db, err := postgresql.NewGormDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = db.AutoMigrate(
		&model.Provider{},
		&model.Token{},
		&model.Operation{},
		&model.History{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}

	log.Println("Database migration completed successfully!")

	initDataSQL := `
	INSERT INTO providers(name) VALUES
		('suilend'),
		('naviprotocol');

	INSERT INTO tokens(name) VALUES
		('Sui'),
		('USDC');

	INSERT INTO operations(name) VALUES
		('borrow'),
		('lend'),
		('stake');
`
	_ = db.Exec(initDataSQL)
	log.Println("Data initialization completed successfully!")
}
