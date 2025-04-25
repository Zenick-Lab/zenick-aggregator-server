package main

import (
	_ "Zenick-Lab/zenick-aggregator-server/docs"
	"Zenick-Lab/zenick-aggregator-server/src/controller"
	"log"

	"github.com/spf13/viper"
)

// @title Zenick Aggregator API
// @version 1.0
// @description This is the API documentation for Zenick Aggregator.
// @host api.lovelyglam.life
// @BasePath /
func main() {
	// viper.SetConfigFile(".env")
	// viper.AutomaticEnv()
	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Printf("error while reading config file: %s", err.Error())
	// 	return
	// }
	// log.Println("Config file loaded successfully")

	// for _, env := range viper.AllKeys() {
	// 	if viper.GetString(env) != "" {
	// 		_ = os.Setenv(env, viper.GetString(env))
	// 		_ = os.Setenv(strings.ToUpper(env), viper.GetString(env))
	// 	}
	// }

	viper.AutomaticEnv()
	log.Println("Environment variables loaded")

	router := controller.Controller()

	if err := router.Run(":8000"); err != nil {
		log.Fatalf("could not run Gin server: %v", err)
	}
}

// package main

// import (
// 	"log"
// 	"os"
// 	"strings"

// 	"Zenick-Lab/zenick-aggregator-server/src/model"
// 	"Zenick-Lab/zenick-aggregator-server/src/pkg/postgresql"

// 	"github.com/spf13/viper"
// )

// func main() {
// 	viper.SetConfigFile(".env")
// 	viper.AutomaticEnv()
// 	if err := viper.ReadInConfig(); err != nil {
// 		log.Printf("error while reading config file: %s", err.Error())
// 		return
// 	}
// 	log.Println("Config file loaded successfully")

// 	for _, env := range viper.AllKeys() {
// 		if viper.GetString(env) != "" {
// 			_ = os.Setenv(env, viper.GetString(env))
// 			_ = os.Setenv(strings.ToUpper(env), viper.GetString(env))
// 		}
// 	}

// 	db, err := postgresql.NewGormDB()
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the database: %v", err)
// 	}

// 	err = db.AutoMigrate(
// 		&model.Provider{},
// 		&model.Token{},
// 		&model.Operation{},
// 		&model.History{},
// 		&model.LiquidityPoolHistory{},
// 		&model.Operation{},
// 		&model.HistoryLink{},
// 		&model.LiquidityPoolHistoryLink{},
// 	)
// 	if err != nil {
// 		log.Fatalf("Failed to migrate the database: %v", err)
// 	}

// 	log.Println("Database migration completed successfully!")

// 	initDataSQL := `
// 		INSERT INTO providers(name) VALUES
// 			('suilend'),
// 			('naviprotocol'),
// 			('cetus'),
// 			('haedal'),
// 			('scallop'),
// 			('bluefin'),
// 			('bucket'),
// 			('alpha_fi'),
// 			('aftermath_finance'),
// 			('kai_finance'),
// 			('kriya'),
// 			('volosui');

// 		INSERT INTO tokens(name) VALUES
// 			('Sui'),
// 			('USDC');

// 		INSERT INTO operations(name) VALUES
// 			('borrow'),
// 			('lend'),
// 			('stake');
// 	`
// 	_ = db.Exec(initDataSQL)
// 	log.Println("Data initialization completed successfully!")
// }
