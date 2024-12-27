package main

import (
	"baseApi/config"
	"baseApi/database"
	"baseApi/routes"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Load environment variables
	config.Init()

	cfg := config.GetAll() // Commented out since it's not used

	// Open new database connection
	DBHOST := cfg.Database.Host
	DBPORT := cfg.Database.Port
	DBUSER := cfg.Database.User
	DBPASS := cfg.Database.Pass
	DBNAME := cfg.Database.Name

	connstring := DBUSER + ":" + DBPASS + "@(" + DBHOST + ":" + DBPORT + ")/" + DBNAME + "?charset=utf8mb4&parseTime=True"
	var err error
	database.DBConn, err = gorm.Open(mysql.Open(connstring), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info), // uncomment to debug query
	})
	if err != nil {
		log.Println(err)
		panic("Failed to connect to the database")
	}
	database.Migrate(database.DBConn) // Comment this line if you don't want to migrate
}

func main() {
	cfg := config.GetAll()

	router := routes.SetupRouter()
	router.Run(":" + cfg.App.Port)
}
